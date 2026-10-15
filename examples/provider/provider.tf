terraform {
  required_providers {
    kong-mesh = {
      source  = "kong/kong-mesh"
      version = "0.9.0"
    }
  }
}

provider "kong-mesh" {
  server_url = "..." # Optional - can use SERVER_URL environment variable
}