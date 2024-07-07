terraform {
  required_providers {
    bridge = {
      source  = "local/bridge"
      version = "1.0.0"
    }
  }
}

provider "bridge" {
  api_url = "http://localhost:5000/api"
}

data "bridge_value" "prod_vpc_id" {
  name = "prod_vpc_id"
}

output "vpc_test" {
  value = data.bridge_value.prod_vpc_id.value
}
