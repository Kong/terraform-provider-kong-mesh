terraform {
  required_providers {
    kong-mesh = {
      source  = "kong/kong-mesh"
      version = "0.5.3"
    }
  }
}

provider "kong-mesh" {
  # Configuration options
}