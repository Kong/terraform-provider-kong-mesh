resource "kong-mesh_mesh_service" "my_meshservice" {
  labels = {
    key = "value"
  }
  mesh = "...my_mesh..."
  name = "...my_name..."
  spec = {
    identities = [
      {
        type  = "SpiffeID"
        value = "...my_value..."
      }
    ]
    ports = [
      {
        app_protocol = "...my_app_protocol..."
        name         = "...my_name..."
        port         = 8
        target_port = {
          integer = 7
        }
      }
    ]
    selector = {
      dataplane_ref = {
        name = "...my_name..."
      }
      dataplane_tags = {
        key = "value"
      }
    }
    state = "Available"
  }
  type = "MeshService"
}