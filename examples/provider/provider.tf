terraform {
  required_providers {
    kong-mesh = {
      source  = "kong/kong-mesh"
      version = "0.2.0"
    }
  }
}

provider "kong-mesh" {
  # Configuration options
}