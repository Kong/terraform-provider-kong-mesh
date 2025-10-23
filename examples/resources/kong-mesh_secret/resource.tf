resource "kong-mesh_secret" "my_secret" {
  data = "...my_data..."
  labels = {
    key = "value"
  }
  mesh = "...my_mesh..."
  name = "...my_name..."
  type = "...my_type..."
}