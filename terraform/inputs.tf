variable "PORT" {
  type = "string"
  description = "Port used by the rds db"
  default = 5432
}

variable "instance_type" {
  type = "string"
  description = "instance class for rds db"
  default = "db.t2.micro"
}

variable "CONSUMER" {
  type = "string"
  description = "app name"
  default = "mobile-prices"
}

variable "allocated_storage" {
  type = "string"
  description = "allocated size for db"
  default = 20
}

variable "ENVIRONMENT" {
  type = "string"
  description = "Environment description"
  default = "staging"
}

variable "db_name" {
  type = "string"
  description = "internal db name"
  default = "pricesdb"
}