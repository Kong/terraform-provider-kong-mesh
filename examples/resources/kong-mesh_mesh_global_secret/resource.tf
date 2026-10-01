resource "kong-mesh_mesh_global_secret" "my_meshglobalsecret" {
  data = "...my_data..."
  labels = {
    key = "value"
  }
  name = "...my_name..."
  type = "...my_type..."
}