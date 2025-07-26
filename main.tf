terraform {
  required_version = ">= 1.11.1"
  required_providers {
    local = {
      source  = "hashicorp/local"
      version = "2.5.3"
    }
    observability = {
      source  = "shakiel.com/providers/observability"
      version = "1.0.0"
    }
  }
}

provider "local" {}

data "local_file" "example" {
  filename = "${path.module}/example.txt"
}

provider "observability" {
  endpoint = data.local_file.example.content
}

resource "observability_example" "example" {
}

output "example" {
  value = observability_example.example.id
}
