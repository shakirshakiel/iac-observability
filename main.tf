terraform {
  required_version = ">= 1.11.1"
  required_providers {
    local = {
      source  = "hashicorp/local"
      version = "2.5.3"
    }
  }
}

provider "local" {}

data "local_file" "example" {
  filename = "${path.module}/example.txt"
}

output "file_content" {
  value = data.local_file.example.content
}
