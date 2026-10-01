resource "kong-mesh_mesh_http_route" "my_meshhttproute" {
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
        hostnames = [
          "..."
        ]
        rules = [
          {
            default = {
              backend_refs = [
                {
                  filters = [
                    {
                      one = {
                        request_header_modifier = {
                          add = [
                            {
                              name  = "...my_name..."
                              value = "...my_value..."
                            }
                          ]
                          remove = [
                            "..."
                          ]
                          set = [
                            {
                              name  = "...my_name..."
                              value = "...my_value..."
                            }
                          ]
                        }
                        request_mirror = {
                          backend_ref = {
                            kind = "MeshExternalService"
                            labels = {
                              key = "value"
                            }
                            port         = 4
                            section_name = "...my_section_name..."
                            weight       = 300045084
                          }
                          percentage = {
                            str = "...my_str..."
                          }
                        }
                        request_redirect = {
                          hostname = "...my_hostname..."
                          path = {
                            two = {
                              replace_full_path    = "...my_replace_full_path..."
                              replace_prefix_match = "...my_replace_prefix_match..."
                              type                 = "ReplaceFullPath"
                            }
                          }
                          port        = 39396
                          scheme      = "https"
                          status_code = 302
                        }
                        response_header_modifier = {
                          add = [
                            {
                              name  = "...my_name..."
                              value = "...my_value..."
                            }
                          ]
                          remove = [
                            "..."
                          ]
                          set = [
                            {
                              name  = "...my_name..."
                              value = "...my_value..."
                            }
                          ]
                        }
                        type = "ResponseHeaderModifier"
                        url_rewrite = {
                          host_to_backend_hostname = false
                          hostname                 = "...my_hostname..."
                          path = {
                            two = {
                              replace_full_path    = "...my_replace_full_path..."
                              replace_prefix_match = "...my_replace_prefix_match..."
                              type                 = "ReplacePrefixMatch"
                            }
                          }
                        }
                      }
                    }
                  ]
                  kind = "MeshMultiZoneService"
                  labels = {
                    key = "value"
                  }
                  port         = 10
                  section_name = "...my_section_name..."
                  weight       = 2238824958
                }
              ]
              filters = [
                {
                  five = {
                    request_header_modifier = {
                      add = [
                        {
                          name  = "...my_name..."
                          value = "...my_value..."
                        }
                      ]
                      remove = [
                        "..."
                      ]
                      set = [
                        {
                          name  = "...my_name..."
                          value = "...my_value..."
                        }
                      ]
                    }
                    request_mirror = {
                      backend_ref = {
                        kind = "MeshService"
                        labels = {
                          key = "value"
                        }
                        port         = 10
                        section_name = "...my_section_name..."
                        weight       = 3071501838
                      }
                      percentage = {
                        integer = 0
                      }
                    }
                    request_redirect = {
                      hostname = "...my_hostname..."
                      path = {
                        one = {
                          replace_full_path    = "...my_replace_full_path..."
                          replace_prefix_match = "...my_replace_prefix_match..."
                          type                 = "ReplaceFullPath"
                        }
                      }
                      port        = 31753
                      scheme      = "https"
                      status_code = 302
                    }
                    response_header_modifier = {
                      add = [
                        {
                          name  = "...my_name..."
                          value = "...my_value..."
                        }
                      ]
                      remove = [
                        "..."
                      ]
                      set = [
                        {
                          name  = "...my_name..."
                          value = "...my_value..."
                        }
                      ]
                    }
                    type = "RequestRedirect"
                    url_rewrite = {
                      host_to_backend_hostname = false
                      hostname                 = "...my_hostname..."
                      path = {
                        one = {
                          replace_full_path    = "...my_replace_full_path..."
                          replace_prefix_match = "...my_replace_prefix_match..."
                          type                 = "ReplaceFullPath"
                        }
                      }
                    }
                  }
                }
              ]
            }
            matches = [
              {
                headers = [
                  {
                    name  = "...my_name..."
                    type  = "Absent"
                    value = "...my_value..."
                  }
                ]
                method = "PATCH"
                path = {
                  type  = "Exact"
                  value = "...my_value..."
                }
                query_params = [
                  {
                    name  = "...my_name..."
                    type  = "Exact"
                    value = "...my_value..."
                  }
                ]
              }
            ]
          }
        ]
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
  type = "MeshHTTPRoute"
}