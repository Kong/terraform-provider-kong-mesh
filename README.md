# terraform-provider-kong-mesh

This repository contains a **BETA** Terraform provider for Kong Mesh.

## Usage

The provider can be installed from the Terraform registry.

```hcl
# Configure the provider to use your Kong Konnect account
terraform {
  required_providers {
    kong-mesh = {
      source  = "kong/kong-mesh"
    }
  }
}

provider "kong-mesh" {
}
```

<!-- No SDK Installation -->
<!-- No SDK Example Usage -->
<!-- No SDK Available Operations -->
<!-- Placeholder for Future Speakeasy SDK Sections -->
