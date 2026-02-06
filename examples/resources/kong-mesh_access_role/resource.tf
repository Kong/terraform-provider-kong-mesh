resource "kong-mesh_access_role" "my_accessrole" {
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
      mesh = "...my_mesh..."
      names = [
        "..."
      ]
      types = [
        "..."
      ]
      when = [
        {
          destinations = {
            match = {
              key = "value"
            }
          }
          dp_token = {
            tags = [
              {
                name  = "...my_name..."
                value = "...my_value..."
              }
            ]
          }
          from = {
            target_ref = {
              kind = "...my_kind..."
              mesh = "...my_mesh..."
              name = "...my_name..."
              tags = {
                key = "value"
              }
            }
          }
          selectors = {
            match = {
              key = "value"
            }
          }
          sources = {
            match = {
              key = "value"
            }
          }
          target_ref = {
            kind = "...my_kind..."
            mesh = "...my_mesh..."
            name = "...my_name..."
            tags = {
              key = "value"
            }
          }
          to = {
            target_ref = {
              kind = "...my_kind..."
              mesh = "...my_mesh..."
              name = "...my_name..."
              tags = {
                key = "value"
              }
            }
          }
        }
      ]
    }
  ]
  type = "...my_type..."
}