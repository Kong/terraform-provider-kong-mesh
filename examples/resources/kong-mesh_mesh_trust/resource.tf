resource "kong-mesh_mesh_trust" "my_meshtrust" {
  labels = {
    key = "value"
  }
  mesh = "...my_mesh..."
  name = "...my_name..."
  spec = {
    ca_bundles = [
      {
        pem = {
          value = "...my_value..."
        }
        type = "Pem"
      }
    ]
    origin = {
      kri = "...my_kri..."
    }
    trust_domain = "...my_trust_domain..."
  }
  type = "MeshTrust"
}