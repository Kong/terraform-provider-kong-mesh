resource "kong-mesh_mesh_fault_injection" "my_meshfaultinjection" {
  labels = {
    key = "value"
  }
  mesh = "...my_mesh..."
  name = "...my_name..."
  spec = {
    rules = [
      {
        default = {
          http = [
            {
              abort = {
                http_status = 4
                percentage = {
                  integer = 5
                }
              }
              delay = {
                percentage = {
                  integer = 7
                }
                value = "...my_value..."
              }
              response_bandwidth = {
                limit = "...my_limit..."
                percentage = {
                  str = "...my_str..."
                }
              }
            }
          ]
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
      kind = "Dataplane"
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
          http = [
            {
              abort = {
                http_status = 0
                percentage = {
                  integer = 10
                }
              }
              delay = {
                percentage = {
                  str = "...my_str..."
                }
                value = "...my_value..."
              }
              response_bandwidth = {
                limit = "...my_limit..."
                percentage = {
                  integer = 6
                }
              }
            }
          ]
        }
        target_ref = {
          kind = "MeshService"
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
  type = "MeshFaultInjection"
}