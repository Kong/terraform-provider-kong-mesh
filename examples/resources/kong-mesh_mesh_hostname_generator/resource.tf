resource "kong-mesh_mesh_hostname_generator" "my_meshhostnamegenerator" {
  labels = {
    key = "value"
  }
  name = "...my_name..."
  spec = {
    extension = {
      config = "{ \"see\": \"documentation\" }"
      type   = "...my_type..."
    }
    selector = {
      mesh_external_service = {
        match_labels = {
          key = "value"
        }
      }
      mesh_multi_zone_service = {
        match_labels = {
          key = "value"
        }
      }
      mesh_service = {
        match_labels = {
          key = "value"
        }
      }
    }
    template = "...my_template..."
  }
  type = "HostnameGenerator"
}