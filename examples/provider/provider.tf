terraform {
  required_providers {
    kong-mesh = {
      source  = "kong/kong-mesh"
      version = "0.6.4"
    }
  }
}

provider "kong-mesh" {
  # Configuration options
}