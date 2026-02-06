resource "kong-mesh_mesh_workload" "my_meshworkload" {
  labels = {
    key = "value"
  }
  mesh = "...my_mesh..."
  name = "...my_name..."
  spec = {
    # ...
  }
  type = "Workload"
}