resource "kong-mesh_dataplane" "my_dataplane" {
  labels = {
    key = "value"
  }
  mesh = "...my_mesh..."
  metrics = {
    conf = {
      prometheus_metrics_backend_config = {
        aggregate = [
          {
            address = "...my_address..."
            enabled = true
            name    = "...my_name..."
            path    = "...my_path..."
            port    = 7
          }
        ]
        envoy = {
          filter_regex = "...my_filter_regex..."
          used_only    = true
        }
        path      = "...my_path..."
        port      = 9
        skip_mtls = false
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
  name = "...my_name..."
  networking = {
    address = "...my_address..."
    admin = {
      port = 9
    }
    advertised_address = "...my_advertised_address..."
    gateway = {
      tags = {
        key = "value"
      }
      type = {
        integer = 1
      }
    }
    inbound = [
      {
        address = "...my_address..."
        health = {
          ready = false
        }
        name            = "...my_name..."
        port            = 8
        service_address = "...my_service_address..."
        service_port    = 10
        service_probe = {
          healthy_threshold = 3
          interval = {
            nanos   = 2
            seconds = 1
          }
          tcp = {
            # ...
          }
          timeout = {
            nanos   = 9
            seconds = 5
          }
          unhealthy_threshold = 7
        }
        state = {
          integer = 6
        }
        tags = {
          key = "value"
        }
      }
    ]
    outbound = [
      {
        address = "...my_address..."
        backend_ref = {
          kind = "...my_kind..."
          labels = {
            key = "value"
          }
          name = "...my_name..."
          port = 2
        }
        port = 9
        tags = {
          key = "value"
        }
      }
    ]
    transparent_proxying = {
      direct_access_services = [
        "..."
      ]
      ip_family_mode = {
        integer = 4
      }
      reachable_backends = {
        refs = [
          {
            kind = "...my_kind..."
            labels = {
              key = "value"
            }
            name      = "...my_name..."
            namespace = "...my_namespace..."
            port      = 2
          }
        ]
      }
      reachable_services = [
        "..."
      ]
      redirect_port_inbound  = 0
      redirect_port_outbound = 6
    }
  }
  probes = {
    endpoints = [
      {
        inbound_path = "...my_inbound_path..."
        inbound_port = 0
        path         = "...my_path..."
      }
    ]
    port = 3
  }
  type = "...my_type..."
}