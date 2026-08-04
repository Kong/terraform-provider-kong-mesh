resource "kong-mesh_mesh_rate_limit" "my_meshratelimit" {
  labels = {
    key = "value"
  }
  mesh = "...my_mesh..."
  name = "...my_name..."
  spec = {
    rules = [
      {
        default = {
          local = {
            http = {
              disabled = false
              on_rate_limit = {
                headers = {
                  add = [
                    {
                      name  = "...my_name..."
                      value = "...my_value..."
                    }
                  ]
                  set = [
                    {
                      name  = "...my_name..."
                      value = "...my_value..."
                    }
                  ]
                }
                status = 5
              }
              request_rate = {
                interval = "...my_interval..."
                num      = 9
              }
            }
            tcp = {
              connection_rate = {
                interval = "...my_interval..."
                num      = 0
              }
              disabled = false
            }
          }
        }
        matches = [
          {
            sni = {
              type  = "Exact"
              value = "...my_value..."
            }
            spiffe_id = {
              type  = "Exact"
              value = "...my_value..."
            }
          }
        ]
      }
    ]
    target_ref = {
      kind = "MeshServiceSubset"
      labels = {
        key = "value"
      }
      mesh         = "...my_mesh..."
      name         = "...my_name..."
      namespace    = "...my_namespace..."
      section_name = "...my_section_name..."
      tags = {
        key = "value"
      }
    }
    to = [
      {
        default = {
          local = {
            http = {
              disabled = false
              on_rate_limit = {
                headers = {
                  add = [
                    {
                      name  = "...my_name..."
                      value = "...my_value..."
                    }
                  ]
                  set = [
                    {
                      name  = "...my_name..."
                      value = "...my_value..."
                    }
                  ]
                }
                status = 9
              }
              request_rate = {
                interval = "...my_interval..."
                num      = 10
              }
            }
            tcp = {
              connection_rate = {
                interval = "...my_interval..."
                num      = 5
              }
              disabled = false
            }
          }
        }
        target_ref = {
          kind = "MeshServiceSubset"
          labels = {
            key = "value"
          }
          mesh         = "...my_mesh..."
          name         = "...my_name..."
          namespace    = "...my_namespace..."
          section_name = "...my_section_name..."
          tags = {
            key = "value"
          }
        }
      }
    ]
  }
  type = "MeshRateLimit"
}