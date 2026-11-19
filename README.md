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

## Kong Mesh compatibility

| Provider version            | Kong Mesh | Released from    |
|-----------------------------|-----------|------------------|
| 0.10.x and later minors     | 3.x       | `main`           |
| 0.9.x                       | 2.14      | `kong-mesh-2.14` |

The 0.9.x line only receives patch releases. New features land on `main`.

Terraform's `~> 0.9` constraint also matches 0.10.x and later, so if you're on Kong Mesh 2.14, pin the patch line instead:

```hcl
kong-mesh = {
  source  = "kong/kong-mesh"
  version = "~> 0.9.0"
}
```

<!-- No SDK Installation -->
<!-- No SDK Example Usage -->
<!-- No SDK Available Operations -->
<!-- Placeholder for Future Speakeasy SDK Sections -->
