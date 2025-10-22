resource "kong-mesh_zone_ingress" "my_zoneingress" {
  available_services = [
    {
      external_service = true
      instances        = 10
      mesh             = "...my_mesh..."
      tags = {
        key = "value"
      }
    }
  ]
  labels = {
    key = "value"
  }
  name = "...my_name..."
  networking = {
    address = "...my_address..."
    admin = {
      port = 0
    }
    advertised_address = "...my_advertised_address..."
    advertised_port    = 9
    port               = 4
  }
  type = "...my_type..."
  zone = "...my_zone..."
}