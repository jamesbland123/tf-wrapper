# Define global vars here

region = "us-west-2"
environment = "dev"

# This is the value of the output from terraform-s3
bucket = "mytest-bucket"

# Key for bastion, gameserver, & playserver
key_name = "DefaultKeyPair"

# Bastion
bastion_ami_name_filter = "bastion-*"
bastion_instance_type = "t2.micro"
