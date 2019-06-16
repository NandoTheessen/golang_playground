provider "aws" {
    region = "eu-central-1"
}

resource "random_string" "username" {
  length  = 16
  number  = false
  special = false
}

resource "random_string" "password" {
  length           = 32
  special          = true
  override_special = "!()+,-.:;<=>?[]{}"
}

resource "aws_security_group" "pricesdb_sec_group" {
  name = "pricesdb_sec_group"
  ingress {
      from_port = 5432
      to_port = 5432
      protocol = "tcp"
    }
}

resource "aws_db_instance" "prices_db" {
    name = "${var.db_name}"
    identifier = "${var.db_name}"

    engine            = "postgres"
    engine_version    = "9.6.9"
    instance_class    = "${var.instance_type}"
    allocated_storage = "${var.allocated_storage}"
    port              = "${var.PORT}"
    username          = "${random_string.username.result}"
    password          = "${random_string.password.result}"
    maintenance_window = "Mon:00:00-Mon:03:00"
}

output "data" {
  value = "${random_string.username.result}, ${random_string.password.result}"
}