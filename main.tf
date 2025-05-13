resource "aws_instance" "example" {
  ami           = var.ami_id
  instance_type = var.instance_type
}

provider "aws" {
  region = "us-east-1"
}