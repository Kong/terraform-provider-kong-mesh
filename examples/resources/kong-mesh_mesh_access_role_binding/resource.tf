resource "kong-mesh_mesh_access_role_binding" "my_meshaccessrolebinding" {
  labels = {
    key = "value"
  }
  name = "...my_name..."
  roles = [
    "..."
  ]
  subjects = [
    {
      name = "...my_name..."
      type = "...my_type..."
    }
  ]
  type = "...my_type..."
}