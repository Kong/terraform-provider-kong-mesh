terraform {
  required_providers {
    kong-mesh = {
      source  = "kong/kong-mesh"
      version = "0.1.5"
    }
  }
}

provider "kong-mesh" {
  # Configuration options
}