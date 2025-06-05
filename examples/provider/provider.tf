terraform {
  required_providers {
    kong-mesh = {
      source  = "kong/kong-mesh"
      version = "0.5.2"
    }
  }
}

provider "kong-mesh" {
  # Configuration options
}