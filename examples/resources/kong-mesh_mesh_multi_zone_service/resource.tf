resource "kong-mesh_mesh_multi_zone_service" "my_meshmultizoneservice" {
  labels = {
    key = "value"
  }
  mesh = "...my_mesh..."
  name = "...my_name..."
  spec = {
    ports = [
      {
        app_protocol = "...my_app_protocol..."
        name         = "...my_name..."
        port         = 5
      }
    ]
    selector = {
      mesh_service = {
        match_labels = {
          key = "value"
        }
      }
    }
  }
  type = "MeshMultiZoneService"
}