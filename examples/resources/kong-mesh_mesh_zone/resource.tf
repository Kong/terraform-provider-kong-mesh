resource "kong-mesh_mesh_zone" "my_meshzone" {
  enabled = true
  labels = {
    key = "value"
  }
  name = "...my_name..."
  type = "...my_type..."
}