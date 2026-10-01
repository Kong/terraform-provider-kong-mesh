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
      section_name = "...my_section_name..."
    }
    to = [
      {
        default = {
          hash_policies = [
            {
              one = {
                connection = {
                  source_ip = true
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
                type     = "FilterState"
              }
            }
          ]
          load_balancer = {
            three = {
              least_request = {
                active_request_bias = {
                  str = "...my_str..."
                }
                choice_count = 5
              }
              maglev = {
                table_size = 3499575
              }
              random = {
                # ...
              }
              ring_hash = {
                hash_function = "MurmurHash2"
                max_ring_size = 5981191
                min_ring_size = 554037
              }
              round_robin = {
                # ...
              }
              type = "Random"
            }
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
          section_name = "...my_section_name..."
        }
      }
    ]
  }
  type = "MeshLoadBalancingStrategy"
}