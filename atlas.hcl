variable "schema" {
    type = string
    default = "default_db"
}

variable "path" {
    type = string
    description = "template dir path"
    default = "sql/schema"
}

data "template_dir" "schema" {
    path = var.path
    vars = {
        schema = var.schema
    }
}

env "dev" {
    src = data.template_dir.schema.url
    schemas = [var.schema]
    url = "mysql://root:root@mysql:3306"
    dev = "mysql://root:root@mysql_dev:3306"
}
