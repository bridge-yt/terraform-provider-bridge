provider "bridge" {
  api_url = "http://localhost:5000/api"
}

locals {
  mock_vpc_id = "vpc-12345678"  # Mock VPC ID
}

resource "null_resource" "mock_vpc" {
  provisioner "local-exec" {
    command = "echo Mock VPC created"
  }
}

resource "bridge_output" "vpc_id" {
  bridge_name     = "prod_vpc_id"
  value           = local.mock_vpc_id
  arn             = "arn:aws:ec2:region:account-id:vpc/vpc-id"
  resource_type   = "vpc"
  bridge_register = true
}


output "vpc_test" {
  value = bridge.prod_vpc_id.value
}