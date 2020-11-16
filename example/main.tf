terraform {
  required_providers {
    grafanaauth = {
      source = "github.com/orendain/grafanaauth"
      versions = ["0.0.2"]
    }
  }
}

provider "grafanaauth" {
  url = "http://localhost:3000"
  username = "admin"
  password = "admin"
//  token = "can use token instead of username and password"
//  organization_id = 2
}

resource "grafanaauth_api_key" "foo" {
  name = "key_foo"
  role = "Viewer"
}

resource "grafanaauth_api_key" "bar" {
  name = "key_bar"
  role = "Admin"
  seconds_to_live = 30
}


output "api_key_foo_key_only" {
  value = grafanaauth_api_key.foo.key
  sensitive = true
}

output "api_key_bar" {
  value = grafanaauth_api_key.bar
}
