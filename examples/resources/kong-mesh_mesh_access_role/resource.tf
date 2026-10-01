resource "kong-mesh_mesh_access_role" "my_meshaccessrole" {
  labels = {
    key = "value"
  }
  name = "...my_name..."
  rules = [
    {
      access = [
        "GENERATE_ZONE_TOKEN"
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
              labels = {
                key = "value"
              }
              name = "...my_name..."
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
            labels = {
              key = "value"
            }
            name = "...my_name..."
          }
          to = {
            target_ref = {
              kind = "...my_kind..."
              labels = {
                key = "value"
              }
              name = "...my_name..."
            }
          }
        }
      ]
    }
  ]
  type = "...my_type..."
}