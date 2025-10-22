terraform {
  required_providers {
    kong-mesh = {
      source  = "kong/kong-mesh"
      version = "0.5.5"
    }
  }
}

provider "kong-mesh" {
  # Configuration options
}