## Remote State Definition
terraform {
  backend "s3" {
    workspace_key_prefix = "tf"
    key                  = "terraform.tfstate"
    dynamodb_table       = "terraform-state-lock"
    encrypt              = "true"
  }
}
