resource "kong-mesh_mesh_trace" "my_meshtrace" {
  labels = {
    key = "value"
  }
  mesh = "...my_mesh..."
  name = "...my_name..."
  spec = {
    default = {
      backends = [
        {
          datadog = {
            split_service = false
            url           = "...my_url..."
          }
          open_telemetry = {
            endpoint = "otel-collector:4317"
          }
          type = "OpenTelemetry"
          zipkin = {
            api_version         = "httpProto"
            shared_span_context = false
            trace_id128bit      = false
            url                 = "...my_url..."
          }
        }
      ]
      sampling = {
        client = {
          str = "...my_str..."
        }
        overall = {
          integer = 4
        }
        random = {
          integer = 4
        }
      }
      tags = [
        {
          header = {
            default = "...my_default..."
            name    = "...my_name..."
          }
          literal = "...my_literal..."
          name    = "...my_name..."
        }
      ]
    }
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
  }
  type = "MeshTrace"
}