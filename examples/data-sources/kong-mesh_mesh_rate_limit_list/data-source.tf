data "kong-mesh_mesh_rate_limit_list" "my_meshratelimitlist" {
  key    = "...my_key..."
  mesh   = "...my_mesh..."
  offset = 0
  size   = 25
  value  = "...my_value..."
}