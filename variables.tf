variable "ami_id" {
  description = "ami-0f88e80871fd81e91"
  type        = string
}

variable "instance_type" {
  description = "Instance type for the EC2 instance"
  type        = string
  default     = "t2.micro"
}
