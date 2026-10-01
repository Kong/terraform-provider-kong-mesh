resource "kong-mesh_mesh_access_log" "my_meshaccesslog" {
  labels = {
    key = "value"
  }
  mesh = "...my_mesh..."
  name = "...my_name..."
  spec = {
    rules = [
      {
        default = {
          backends = [
            {
              one = {
                file = {
                  format = {
                    two = {
                      json = [
                        {
                          key   = "...my_key..."
                          value = "...my_value..."
                        }
                      ]
                      omit_empty_values = false
                      plain             = "[%START_TIME%] %KUMA_MESH% %UPSTREAM_HOST%"
                      type              = "Plain"
                    }
                  }
                  path = "/tmp/access.log"
                }
                open_telemetry = {
                  attributes = [
                    {
                      key   = "...my_key..."
                      value = "...my_value..."
                    }
                  ]
                  backend_ref = {
                    kind = "MeshOpenTelemetryBackend"
                    labels = {
                      key = "value"
                    }
                  }
                  body = { "kvlistValue" : { "values" : [{ "key" : "mesh", "value" : { "stringValue" : "%KUMA_MESH%" } }] } }
                }
                tcp = {
                  address = "127.0.0.1:5000"
                  format = {
                    one = {
                      json = [
                        {
                          key   = "...my_key..."
                          value = "...my_value..."
                        }
                      ]
                      omit_empty_values = false
                      plain             = "[%START_TIME%] %KUMA_MESH% %UPSTREAM_HOST%"
                      type              = "Json"
                    }
                  }
                }
                type = "OpenTelemetry"
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
              type  = "Prefix"
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
      section_name = "...my_section_name..."
    }
    to = [
      {
        default = {
          backends = [
            {
              one = {
                file = {
                  format = {
                    two = {
                      json = [
                        {
                          key   = "...my_key..."
                          value = "...my_value..."
                        }
                      ]
                      omit_empty_values = false
                      plain             = "[%START_TIME%] %KUMA_MESH% %UPSTREAM_HOST%"
                      type              = "Json"
                    }
                  }
                  path = "/tmp/access.log"
                }
                open_telemetry = {
                  attributes = [
                    {
                      key   = "...my_key..."
                      value = "...my_value..."
                    }
                  ]
                  backend_ref = {
                    kind = "MeshOpenTelemetryBackend"
                    labels = {
                      key = "value"
                    }
                  }
                  body = { "kvlistValue" : { "values" : [{ "key" : "mesh", "value" : { "stringValue" : "%KUMA_MESH%" } }] } }
                }
                tcp = {
                  address = "127.0.0.1:5000"
                  format = {
                    two = {
                      json = [
                        {
                          key   = "...my_key..."
                          value = "...my_value..."
                        }
                      ]
                      omit_empty_values = false
                      plain             = "[%START_TIME%] %KUMA_MESH% %UPSTREAM_HOST%"
                      type              = "Json"
                    }
                  }
                }
                type = "Tcp"
              }
            }
          ]
        }
        target_ref = {
          kind = "MeshExternalService"
          labels = {
            key = "value"
          }
          section_name = "...my_section_name..."
        }
      }
    ]
  }
  type = "MeshAccessLog"
}