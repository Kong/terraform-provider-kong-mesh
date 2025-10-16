resource "kong-mesh_mesh_traffic_permission" "my_meshtrafficpermission" {
  labels = {
    key = "value"
  }
  mesh = "...my_mesh..."
  name = "...my_name..."
  spec = {
    from = [
      {
        default = {
          action = "Deny"
        }
        target_ref = {
          kind = "MeshService"
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
    ]
    rules = [
      {
        default = {
          allow = [
            {
              spiffe_id = {
                type  = "Exact"
                value = "...my_value..."
              }
            }
          ]
          allow_with_shadow_deny = [
            {
              spiffe_id = {
                type  = "Prefix"
                value = "...my_value..."
              }
            }
          ]
          deny = [
            {
              spiffe_id = {
                type  = "Prefix"
                value = "...my_value..."
              }
            }
          ]
        }
      }
    ]
    target_ref = {
      kind = "MeshHTTPRoute"
      labels = {
        key = "value"
      }
      mesh      = "...my_mesh..."
      name      = "...my_name..."
      namespace = "...my_namespace..."
      proxy_types = [
        "Sidecar"
      ]
      section_name = "...my_section_name..."
      tags = {
        key = "value"
      }
    }
  }
  type = "MeshTrafficPermission"
}