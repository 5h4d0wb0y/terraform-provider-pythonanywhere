provider "pythonanywhere" {
  token = "${var.api_token}"
  username = "${var.username}"
}

resource "webapps" "test" {
  domain_name = ""
  python_version = ""
}

resource "consoles" "test" {
}

resource "files" "test" {
}

resource "schedule" "test" {
}