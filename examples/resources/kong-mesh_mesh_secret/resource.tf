resource "kong-mesh_mesh_secret" "my_meshsecret" {
  data = "...my_data..."
  labels = {
    key = "value"
  }
  mesh = "...my_mesh..."
  name = "...my_name..."
  type = "...my_type..."
}