resource "kong-mesh_mesh_gateway" "my_meshgateway" {
  conf = {
    listeners = [
      {
        cross_mesh = true
        hostname   = "...my_hostname..."
        port       = 5
        protocol = {
          str = "...my_str..."
        }
        resources = {
          connection_limit = 5
        }
        tags = {
          key = "value"
        }
        tls = {
          certificates = [
            {
              data_source_file = {
                file = "...my_file..."
              }
            }
          ]
          mode = {
            integer = 4
          }
          options = {
            # ...
          }
        }
      }
    ]
  }
  labels = {
    key = "value"
  }
  mesh = "...my_mesh..."
  name = "...my_name..."
  selectors = [
    {
      match = {
        key = "value"
      }
    }
  ]
  tags = {
    key = "value"
  }
  type = "...my_type..."
}