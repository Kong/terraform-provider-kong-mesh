resource "kong-mesh_mesh_opa" "my_meshopa" {
  labels = {
    key = "value"
  }
  mesh = "...my_mesh..."
  name = "...my_name..."
  spec = {
    default = {
      agent_config = {
        env_var = {
          name = "...my_name..."
        }
        file = {
          path = "...my_path..."
        }
        insecure_inline = {
          value = "...my_value..."
        }
        secret_ref = {
          kind = "Secret"
          name = "...my_name..."
        }
        type = "File"
      }
      append_policies = [
        {
          ignore_decision = true
          rego = {
            env_var = {
              name = "...my_name..."
            }
            file = {
              path = "...my_path..."
            }
            insecure_inline = {
              value = "...my_value..."
            }
            secret_ref = {
              kind = "Secret"
              name = "...my_name..."
            }
            type = "Secret"
          }
        }
      ]
      auth_config = {
        on_agent_failure = "Deny"
        request_body = {
          max_size      = 6
          send_raw_body = true
        }
        status_on_error = 5
        timeout         = "...my_timeout..."
      }
    }
    target_ref = {
      kind = "Mesh"
      labels = {
        key = "value"
      }
      section_name = "...my_section_name..."
    }
  }
  type = "MeshOPA"
}