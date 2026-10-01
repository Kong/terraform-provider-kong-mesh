resource "kong-mesh_mesh_open_telemetry_backend" "my_meshopentelemetrybackend" {
  labels = {
    key = "value"
  }
  mesh = "...my_mesh..."
  name = "...my_name..."
  spec = {
    endpoint = {
      address = "...my_address..."
      path    = "...my_path..."
      port    = 9
    }
    env = {
      allow_signal_overrides = true
      mode                   = "Optional"
      precedence             = "EnvFirst"
    }
    protocol = "grpc"
  }
  type = "MeshOpenTelemetryBackend"
}