resource "kong-mesh_mesh_zone_egress" "my_meshzoneegress" {
  labels = {
    key = "value"
  }
  name = "...my_name..."
  networking = {
    address = "...my_address..."
    admin = {
      port = 4
    }
    port = 2
  }
  type = "...my_type..."
  zone = "...my_zone..."
}