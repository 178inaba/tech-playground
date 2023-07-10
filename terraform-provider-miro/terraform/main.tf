terraform {
  required_providers {
    miro = {
      source = "178inaba/miro"
    }
  }
}

variable "access_token" {
  type      = string
  sensitive = true
}

provider "miro" {
  access_token = var.access_token
}

resource "miro_board" "test" {
  name        = "Test Board"
  description = "My test board for Software Design"
}
