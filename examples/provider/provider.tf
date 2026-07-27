terraform {
  required_providers {
    kong-mesh = {
      source  = "kong/kong-mesh"
      version = "0.8.2"
    }
  }
}

provider "kong-mesh" {
  server_url = "..." # Optional - can use SERVER_URL environment variable
}