terraform {
  required_providers {
    miro     = {
      source = "178inaba/miro"
    }
  }
}

provider "miro" {
  access_token = var.access_token
}
