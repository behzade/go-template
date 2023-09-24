create database {{.schema}};

create table {{.schema}}.user (
    {{template "id"}},
    `name` varchar(255) not null,
    {{template "createdAt"}},
    {{template "updatedAt"}},
    PRIMARY KEY(`id`)
);
