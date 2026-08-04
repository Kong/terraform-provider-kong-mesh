resource "kong-mesh_mesh_tls" "my_meshtls" {
  labels = {
    key = "value"
  }
  mesh = "...my_mesh..."
  name = "...my_name..."
  spec = {
    rules = [
      {
        default = {
          mode = "Strict"
          tls_ciphers = [
            "ECDHE-RSA-AES256-GCM-SHA384"
          ]
          tls_version = {
            max = "TLSAuto"
            min = "TLSAuto"
          }
        }
      }
    ]
    target_ref = {
      kind = "MeshExternalService"
      labels = {
        key = "value"
      }
      mesh         = "...my_mesh..."
      name         = "...my_name..."
      namespace    = "...my_namespace..."
      section_name = "...my_section_name..."
      tags = {
        key = "value"
      }
    }
  }
  type = "MeshTLS"
}