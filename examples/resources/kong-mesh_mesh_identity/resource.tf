resource "kong-mesh_mesh_identity" "my_meshidentity" {
  labels = {
    key = "value"
  }
  mesh = "...my_mesh..."
  name = "...my_name..."
  spec = {
    provider = {
      three = {
        bundled = {
          autogenerate = {
            enabled = false
          }
          ca = {
            certificate = {
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
            private_key = {
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
              type = "InsecureInline"
            }
          }
          certificate_parameters = {
            expiry = "...my_expiry..."
          }
          insecure_allow_self_signed = false
          mesh_trust_creation        = "Disabled"
        }
        extension = {
          certmanager = {
            config = "{ \"see\": \"documentation\" }"
            name   = "...my_name..."
          }
        }
        spire = {
          agent = {
            timeout = "...my_timeout..."
          }
        }
        type = "Spire"
      }
    }
    selector = {
      dataplane = {
        match_labels = {
          key = "value"
        }
      }
    }
    spiffe_id = {
      path         = "...my_path..."
      trust_domain = "...my_trust_domain..."
    }
  }
  type = "MeshIdentity"
}