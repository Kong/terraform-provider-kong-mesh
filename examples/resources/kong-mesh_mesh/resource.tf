resource "kong-mesh_mesh" "my_mesh" {
  constraints = {
    dataplane_proxy = {
      requirements = [
        {
          tags = {
            key = "value"
          }
        }
      ]
      restrictions = [
        {
          tags = {
            key = "value"
          }
        }
      ]
    }
  }
  labels = {
    key = "value"
  }
  logging = {
    backends = [
      {
        conf = {
          file_logging_backend_config = {
            path = "...my_path..."
          }
        }
        format = "...my_format..."
        name   = "...my_name..."
        type   = "...my_type..."
      }
    ]
    default_backend = "...my_default_backend..."
  }
  mesh_services = {
    mode = {
      str = "...my_str..."
    }
  }
  metrics = {
    backends = [
      {
        conf = {
          prometheus_metrics_backend_config = {
            aggregate = [
              {
                address = "...my_address..."
                enabled = true
                name    = "...my_name..."
                path    = "...my_path..."
                port    = 10
              }
            ]
            envoy = {
              filter_regex = "...my_filter_regex..."
              used_only    = true
            }
            path      = "...my_path..."
            port      = 5
            skip_mtls = true
            tags = {
              key = "value"
            }
            tls = {
              mode = {
                str = "...my_str..."
              }
            }
          }
        }
        name = "...my_name..."
        type = "...my_type..."
      }
    ]
    enabled_backend = "...my_enabled_backend..."
  }
  mtls = {
    backends = [
      {
        conf = {
          vault_certificate_authority_config = {
            mode = "{ \"see\": \"documentation\" }"
          }
        }
        dp_cert = {
          request_timeout = {
            nanos   = 5
            seconds = 5
          }
          rotation = {
            expiration = "...my_expiration..."
          }
        }
        mode = {
          integer = 4
        }
        name = "...my_name..."
        root_chain = {
          request_timeout = {
            nanos   = 9
            seconds = 4
          }
        }
        type = "...my_type..."
      }
    ]
    enabled_backend = "...my_enabled_backend..."
    skip_validation = true
  }
  name = "...my_name..."
  networking = {
    outbound = {
      passthrough = false
    }
  }
  routing = {
    default_forbid_mesh_external_service_access = false
    locality_aware_load_balancing               = false
    zone_egress                                 = false
  }
  skip_creating_initial_policies = [
    "..."
  ]
  tracing = {
    backends = [
      {
        conf = {
          zipkin_tracing_backend_config = {
            api_version         = "...my_api_version..."
            shared_span_context = false
            trace_id128bit      = true
            url                 = "...my_url..."
          }
        }
        name     = "...my_name..."
        sampling = 9.35
        type     = "...my_type..."
      }
    ]
    default_backend = "...my_default_backend..."
  }
  type = "...my_type..."
}