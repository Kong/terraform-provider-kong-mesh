resource "kong-mesh_access_role_binding" "my_accessrolebinding" {
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