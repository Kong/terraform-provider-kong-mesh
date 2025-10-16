resource "kong-mesh_mesh_fault_injection" "my_meshfaultinjection" {
  labels = {
    key = "value"
  }
  mesh = "...my_mesh..."
  name = "...my_name..."
  spec = {
    from = [
      {
        default = {
          http = [
            {
              abort = {
                http_status = 1
                percentage = {
                  str = "...my_str..."
                }
              }
              delay = {
                percentage = {
                  integer = 8
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
        target_ref = {
          kind = "MeshServiceSubset"
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
    ]
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
          kind = "MeshGateway"
          labels = {
            key = "value"
          }
          mesh      = "...my_mesh..."
          name      = "...my_name..."
          namespace = "...my_namespace..."
          proxy_types = [
            "Sidecar"
          ]
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