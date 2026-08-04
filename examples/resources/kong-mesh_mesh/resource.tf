resource "kong-mesh_mesh" "my_mesh" {
  labels = {
    key = "value"
  }
  mtls = {
    backends = [
      {
        conf = {
          vault_certificate_authority_config = {
            vault_certificate_authority_config_from_cp = {
              from_cp = {
                address       = "...my_address..."
                agent_address = "...my_agent_address..."
                auth = {
                  vault_certificate_authority_config_from_cp_auth_token = {
                    token = {
                      data_source_secret = {
                        secret = "...my_secret..."
                      }
                    }
                  }
                }
                common_name = "...my_common_name..."
                namespace   = "...my_namespace..."
                pki         = "...my_pki..."
                role        = "...my_role..."
                tls = {
                  ca_cert = {
                    data_source_secret = {
                      secret = "...my_secret..."
                    }
                  }
                  server_name = "...my_server_name..."
                  skip_verify = false
                }
              }
            }
          }
        }
        dp_cert = {
          request_timeout = {
            nanos   = 5
            seconds = 5
          }
          rotation = {
            expiration = "...my_expiration..."
          }
        }
        mode = {
          integer = 4
        }
        name = "...my_name..."
        root_chain = {
          request_timeout = {
            nanos   = 9
            seconds = 4
          }
        }
        type = "...my_type..."
      }
    ]
    enabled_backend = "...my_enabled_backend..."
    skip_validation = true
  }
  name = "...my_name..."
  routing = {
    # ...
  }
  skip_creating_initial_policies = [
    "..."
  ]
  type = "...my_type..."
}