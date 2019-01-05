## Variable definitions

variable "region" {}

# Platform environment.
variable "environment" {}

variable "tags" {
  type = "map"

  default = {
    Owner = "Me"
  }
}

################
# Bastion Host #
################

variable "bastion_ami_name_filter" {}

variable "bastion_instance_type" {}

variable "key_name" {}

variable "bucket" {}
