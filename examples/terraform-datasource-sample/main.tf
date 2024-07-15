provider "bridge" {
  api_url = "http://localhost:5000/api"
}

data "bridge_value" "prod_vpc_id" {
  namespace = "production"
  name      = "prod_vpc_id"
}

output "vpc_test" {
  value = data.bridge_value.prod_vpc_id.value
}
