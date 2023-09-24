variable "schema" {
    type = string
}

data "template_dir" "schema" {
    path = "sql/schema"
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
