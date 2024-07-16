# Bridge Provider Documentation

## Overview
The Bridge provider allows you to manage resources seamlessly across various environments using the Bridge platform.

## Example Usage

### Provider Configuration
```hcl
provider "bridge" {
  api_url = "http://your-bridge-api-url"
}
```

### Resource Definition
```hcl
resource "bridge_resource" "example" {
  namespace       = "default"
  bridge_name     = "example-resource"
  value           = "example-value"
  arn             = "arn:aws:iam::123456789012:role/example-role"
  resource_type   = "example-type"
  bridge_register = true
}
```

### Data Source
```hcl
data "bridge_value" "example" {
  namespace = "default"
  name      = "example-resource"
}

output "example_value" {
  value = data.bridge_value.example.value
}
```

## Arguments

### Provider
- **api_url**: (Required) The API URL of the Bridge server.

### Resource `bridge_resource`
- **namespace**: (Required) The namespace to which the resource belongs.
- **bridge_name**: (Required) The unique name of the resource.
- **value**: (Optional) The value of the resource.
- **arn**: (Optional) The ARN of the resource.
- **resource_type**: (Optional) The type of the resource.
- **bridge_register**: (Optional) Whether to register the resource with Bridge.

### Data Source `bridge_value`
- **namespace**: (Required) The namespace to which the resource belongs.
- **name**: (Required) The unique name of the resource.

## Outputs
- **value**: The value of the resource.
- **arn**: The ARN of the resource.
- **resource_type**: The type of the resource.

## Notes
- Ensure the Bridge server is running and accessible via the provided `api_url`.
- Use namespaces to segregate resources in different environments or projects.

## Report an Issue
For any issues or feature requests, please [report here](https://github.com/bridge.yt/issues).

---
