resource "kong-mesh_mesh_passthrough" "my_meshpassthrough" {
  labels = {
    key = "value"
  }
  mesh = "...my_mesh..."
  name = "...my_name..."
  spec = {
    default = {
      append_match = [
        {
          port     = 6
          protocol = "mysql"
          type     = "IP"
          value    = "...my_value..."
        }
      ]
      passthrough_mode = "Matched"
    }
    target_ref = {
      kind = "MeshMultiZoneService"
      labels = {
        key = "value"
      }
      mesh      = "...my_mesh..."
      name      = "...my_name..."
      namespace = "...my_namespace..."
      proxy_types = [
        "Gateway"
      ]
      section_name = "...my_section_name..."
      tags = {
        key = "value"
      }
    }
  }
  type = "MeshPassthrough"
}