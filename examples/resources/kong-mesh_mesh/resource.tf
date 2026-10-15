resource "kong-mesh_mesh" "my_mesh" {
  labels = {
    key = "value"
  }
  mesh_services = {
    mode = "Everywhere"
  }
  name = "...my_name..."
  type = "...my_type..."
}