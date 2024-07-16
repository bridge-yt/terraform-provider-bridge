
```markdown
# Terraform Provider Bridge

## Overview

The Bridge Terraform provider allows you to interact with the Bridge API to manage and retrieve resource information. This provider is particularly useful for referencing and sharing resources seamlessly across different tools like Terraform, Crossplane, and ArgoCD.

## Installation

To install this provider, include the following in your Terraform configuration:

```hcl
terraform {
  required_providers {
    bridge = {
      source  = "bridge-yt/bridge"
      version = "1.0.0"
    }
  }
}

provider "bridge" {
  api_url = "http://your-api-url"
}
```

Replace `http://your-api-url` with the actual URL of your Bridge API.

## Usage

### Resource Example

This example demonstrates how to create a resource using the Bridge provider:

```hcl
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
  value = bridge_output.vpc_id.value
}
```

### Datasource Example

This example demonstrates how to use the Bridge provider to retrieve data:

```hcl
terraform {
  required_providers {
    bridge = {
      source  = "Helm-bridge-plugin-yt/Helm-bridge-plugin"
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
```

## Configuration

### Provider Configuration

The Bridge provider requires the following configuration parameters:

- `api_url`: The URL of the Bridge API.

Example configuration:

```hcl
provider "bridge" {
  api_url = "http://your-api-url"
}
```

### Resource: `bridge_output`

This resource allows you to create and manage outputs in the Bridge API.

#### Arguments

- `bridge_name` (Required): The name of the bridge resource.
- `value` (Required): The value of the resource.
- `arn` (Required): The Amazon Resource Name (ARN) of the resource.
- `resource_type` (Required): The type of the resource.
- `bridge_register` (Optional): Whether to register the resource in the Bridge API. Default is `false`.

Example usage:

```hcl
resource "bridge_output" "example" {
  bridge_name     = "example_name"
  value           = "example_value"
  arn             = "arn:aws:ec2:region:account-id:vpc/vpc-id"
  resource_type   = "vpc"
  bridge_register = true
}
```

### Data Source: `bridge_value`

This data source allows you to retrieve the value of a resource from the Bridge API.

#### Arguments

- `name` (Required): The name of the bridge resource.

Example usage:

```hcl
data "bridge_value" "example" {
  name = "example_name"
}

output "example_value" {
  value = data.bridge_value.example.value
}
```

## Examples

### Resource Example

This example demonstrates how to create a resource using the Bridge provider:

```hcl
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
  value = bridge_output.vpc_id.value
}
```

### Datasource Example

This example demonstrates how to use the Bridge provider to retrieve data:

```hcl
terraform {
  required_providers {
    bridge = {
      source  = "Helm-bridge-plugin-yt/Helm-bridge-plugin"
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
```

## Development

### Requirements

- Go 1.16 or later
- Terraform SDK v2

### Building the Provider

Clone the repository and navigate to the provider directory:

```bash
git clone https://github.com/bridge-yt/terraform-provider-bridge.git
cd terraform-provider-Helm-bridge-plugin
```

Build the provider:

```bash
go build -o terraform-provider-Helm-bridge-plugin
```

### Testing the Provider Locally

To test the provider locally, create a Terraform configuration that uses the local provider binary. For example:

```hcl
terraform {
  required_providers {
    bridge = {
      source = "local/Helm-bridge-plugin"
      version = "1.0.0"
    }
  }
}

provider "bridge" {
  api_url = "http://localhost:5000/api"
}

resource "bridge_output" "vpc_id" {
  bridge_name     = "prod_vpc_id"
  value           = "vpc-12345678"
  arn             = "arn:aws:ec2:region:account-id:vpc/vpc-id"
  resource_type   = "vpc"
  bridge_register = true
}

output "vpc_test" {
  value = bridge_output.vpc_id.value
}
```

Initialize and apply the configuration:

```bash
terraform init
terraform apply
```

## Contributing

We welcome contributions! Please open an issue or submit a pull request for any bugs or features.

## License

This project is licensed under the MIT License - see the [LICENSE](LICENSE) file for details.
```