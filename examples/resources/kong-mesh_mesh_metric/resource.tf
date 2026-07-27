resource "kong-mesh_mesh_metric" "my_meshmetric" {
  labels = {
    key = "value"
  }
  mesh = "...my_mesh..."
  name = "...my_name..."
  spec = {
    default = {
      applications = [
        {
          address = "...my_address..."
          name    = "...my_name..."
          path    = "/metrics"
          port    = 8
        }
      ]
      backends = [
        {
          open_telemetry = {
            endpoint         = "...my_endpoint..."
            refresh_interval = "...my_refresh_interval..."
          }
          prometheus = {
            client_id = "...my_client_id..."
            path      = "/metrics"
            port      = 5670
            tls = {
              mode = "Disabled"
            }
          }
          type = "Prometheus"
        }
      ]
      sidecar = {
        include_unused = false
        profiles = {
          append_profiles = [
            {
              name = "Basic"
            }
          ]
          exclude = [
            {
              match = "...my_match..."
              type  = "Contains"
            }
          ]
          include = [
            {
              match = "...my_match..."
              type  = "Regex"
            }
          ]
        }
      }
    }
    target_ref = {
      kind = "MeshSubset"
      labels = {
        key = "value"
      }
      mesh      = "...my_mesh..."
      name      = "...my_name..."
      namespace = "...my_namespace..."
      proxy_types = [
        "Gateway"
      ]
      section_name = "...my_section_name..."
      tags = {
        key = "value"
      }
    }
  }
  type = "MeshMetric"
}