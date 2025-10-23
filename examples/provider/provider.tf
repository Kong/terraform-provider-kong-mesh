terraform {
  required_providers {
    kong-mesh = {
      source  = "kong/kong-mesh"
      version = "0.6.0"
    }
  }
}

provider "kong-mesh" {
  # Configuration options
}