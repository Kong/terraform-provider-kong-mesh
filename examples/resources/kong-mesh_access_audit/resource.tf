resource "kong-mesh_access_audit" "my_accessaudit" {
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
      access_all = false
      mesh       = "...my_mesh..."
      types = [
        "..."
      ]
    }
  ]
  type = "...my_type..."
}