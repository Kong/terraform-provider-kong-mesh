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
          protocol = "tcp"
          type     = "IP"
          value    = "...my_value..."
        }
      ]
      passthrough_mode = "Matched"
    }
    target_ref = {
      kind = "Dataplane"
      labels = {
        key = "value"
      }
      section_name = "...my_section_name..."
    }
  }
  type = "MeshPassthrough"
}