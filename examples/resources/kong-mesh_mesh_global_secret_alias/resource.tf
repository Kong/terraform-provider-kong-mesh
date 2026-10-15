resource "kong-mesh_mesh_global_secret_alias" "my_meshglobalsecretalias" {
  data = "...my_data..."
  labels = {
    key = "value"
  }
  name = "...my_name..."
  type = "...my_type..."
}