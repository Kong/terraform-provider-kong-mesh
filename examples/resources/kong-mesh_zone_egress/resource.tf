resource "kong-mesh_zone_egress" "my_zoneegress" {
  labels = {
    key = "value"
  }
  name = "...my_name..."
  networking = {
    address = "...my_address..."
    admin = {
      port = 8
    }
    port = 9
  }
  type = "...my_type..."
  zone = "...my_zone..."
}