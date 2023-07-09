terraform {
  required_providers {
    miro = {
      source = "178inaba/miro"
    }
  }
}

provider "miro" {
  access_token = var.access_token
}

resource "miro_board" "test" {
  name        = "Test Board"
  description = "My test board for Software Design"
}
