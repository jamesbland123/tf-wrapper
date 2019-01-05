# The bastion host acts as the "jump point" for the rest of the infrastructure.
# Since most of our instances aren't exposed to the external internet, the
# bastion acts as the gatekeeper for any direct SSH access.
# The bastion is provisioned using the key name that you pass to the stack (and
# hopefully have stored somewhere).
# If you ever need to access an instance directly, you can do it by "jumping
# through" the bastion:
#
#    $ terraform output --module=network | grep bastion
#    $ ssh -i <path/to/key> ubuntu@<bastion-ip> ssh ubuntu@<internal-ip>
provider "aws" {
  region = "${var.region}"
}

locals {
  region_shortened = "${replace(var.region, "-" , "")}"
}

data "aws_caller_identity" "current" {}

data "terraform_remote_state" "vpc" {
  backend   = "s3"
  workspace = "network-${var.environment}-${local.region_shortened}"

  config {
    bucket               = "${var.bucket}"
    workspace_key_prefix = "tf"
    key                  = "terraform.tfstate"
    region               = "${var.region}"
  }
}

resource "aws_key_pair" "default" {
  key_name   = "DefaultKeyPair"
  public_key = "ssh-rsa REPLACE_ME"
}

resource "aws_security_group" "external_ssh" {
  name        = "${format("%s-external-ssh", var.environment)}"
  description = "Allows SSH from the world"
  vpc_id      = "${data.terraform_remote_state.vpc.vpc_id}"

  ingress {
    from_port   = 22
    to_port     = 22
    protocol    = "tcp"
    cidr_blocks = ["0.0.0.0/0"]
  }

  egress {
    from_port   = 0
    to_port     = 0
    protocol    = "-1"
    cidr_blocks = ["0.0.0.0/0"]
  }

  lifecycle {
    create_before_destroy = true
  }

  tags = "${merge(var.tags, map("Name", format("%s-external-ssh", var.environment)))}"
}

data "aws_ami" "base" {
  most_recent = true

  filter {
    name   = "name"
    values = ["${var.bastion_ami_name_filter}"]
  }

  filter {
    name   = "virtualization-type"
    values = ["hvm"]
  }

  owners = ["${data.aws_caller_identity.current.account_id}"]
}

resource "aws_instance" "bastion" {
  ami                    = "${data.aws_ami.base.id}"
  instance_type          = "${var.bastion_instance_type}"
  subnet_id              = "${element(data.terraform_remote_state.vpc.public_subnets, 0)}"
  key_name               = "${aws_key_pair.default.key_name}"
  vpc_security_group_ids = ["${aws_security_group.external_ssh.id}"]
  tags                   = "${merge(var.tags, map("Name", format("%s-bastion", var.environment)))}"
  monitoring             = false
}

resource "aws_eip" "bastion" {
  vpc      = true
  instance = "${aws_instance.bastion.id}"
}
