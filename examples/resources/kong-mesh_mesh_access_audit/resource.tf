resource "kong-mesh_mesh_access_audit" "my_meshaccessaudit" {
  labels = {
    key = "value"
  }
  name = "...my_name..."
  rules = [
    {
      access = [
        {
          str = "...my_str..."
        }
      ]
      access_all = true
      mesh       = "...my_mesh..."
      types = [
        "..."
      ]
    }
  ]
  type = "...my_type..."
}