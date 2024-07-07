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
