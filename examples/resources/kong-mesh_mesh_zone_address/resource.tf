resource "kong-mesh_mesh_zone_address" "my_meshzoneaddress" {
  labels = {
    key = "value"
  }
  mesh = "...my_mesh..."
  name = "...my_name..."
  spec = {
    address = "...my_address..."
    port    = 62290
  }
  type = "MeshZoneAddress"
}