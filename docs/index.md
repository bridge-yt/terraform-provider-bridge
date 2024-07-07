# Bridge Provider

The Bridge provider allows you to manage resources with Bridge.

## Example Usage

```hcl
provider "bridge" {}

resource "bridge_resource" "example" {
  name = "example-resource"
}

output "example_arn" {
  value = bridge_resource.example.arn
}
