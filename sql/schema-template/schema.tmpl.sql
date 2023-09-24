-- Only edit the schema.tmpl.sql template file and not the generated schema.sql as it will be overwritten
create database {{.schema}};

create table {{.schema}}.user (
    {{template "id"}},
    `name` varchar(255) not null,
    {{template "createdAt"}},
    {{template "updatedAt"}},
    PRIMARY KEY(`id`)
);
