resource "kong-mesh_mesh_load_balancing_strategy" "my_meshloadbalancingstrategy" {
  labels = {
    key = "value"
  }
  mesh = "...my_mesh..."
  name = "...my_name..."
  spec = {
    target_ref = {
      kind = "Dataplane"
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
    to = [
      {
        default = {
          hash_policies = [
            {
              connection = {
                source_ip = false
              }
              cookie = {
                name = "...my_name..."
                path = "...my_path..."
                ttl  = "...my_ttl..."
              }
              filter_state = {
                key = "...my_key..."
              }
              header = {
                name = "...my_name..."
              }
              query_parameter = {
                name = "...my_name..."
              }
              terminal = true
              type     = "FilterState"
            }
          ]
          load_balancer = {
            least_request = {
              active_request_bias = {
                integer = 10
              }
              choice_count = 4
            }
            maglev = {
              hash_policies = [
                {
                  connection = {
                    source_ip = false
                  }
                  cookie = {
                    name = "...my_name..."
                    path = "...my_path..."
                    ttl  = "...my_ttl..."
                  }
                  filter_state = {
                    key = "...my_key..."
                  }
                  header = {
                    name = "...my_name..."
                  }
                  query_parameter = {
                    name = "...my_name..."
                  }
                  terminal = false
                  type     = "Connection"
                }
              ]
              table_size = 26413
            }
            random = {
              # ...
            }
            ring_hash = {
              hash_function = "XXHash"
              hash_policies = [
                {
                  connection = {
                    source_ip = false
                  }
                  cookie = {
                    name = "...my_name..."
                    path = "...my_path..."
                    ttl  = "...my_ttl..."
                  }
                  filter_state = {
                    key = "...my_key..."
                  }
                  header = {
                    name = "...my_name..."
                  }
                  query_parameter = {
                    name = "...my_name..."
                  }
                  terminal = false
                  type     = "QueryParameter"
                }
              ]
              max_ring_size = 5614666
              min_ring_size = 623920
            }
            round_robin = {
              # ...
            }
            type = "Maglev"
          }
          locality_awareness = {
            cross_zone = {
              failover = [
                {
                  from = {
                    zones = [
                      "..."
                    ]
                  }
                  to = {
                    type = "Any"
                    zones = [
                      "..."
                    ]
                  }
                }
              ]
              failover_threshold = {
                percentage = {
                  str = "...my_str..."
                }
              }
            }
            disabled = true
            local_zone = {
              affinity_tags = [
                {
                  key    = "...my_key..."
                  weight = 7
                }
              ]
            }
          }
        }
        target_ref = {
          kind = "MeshMultiZoneService"
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
  type = "MeshLoadBalancingStrategy"
}