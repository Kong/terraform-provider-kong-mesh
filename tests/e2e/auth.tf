terraform {
  required_providers {
    kong-mesh = {
      source  = "kong/kong-mesh"
      version = "999.99.9" # This ensures that we only test with the locally built provider
    }
  }
}

provider "kong-mesh" {
  server_url = "https://us.api.konghq.com"
}

