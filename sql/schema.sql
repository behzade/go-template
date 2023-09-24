-- Only edit the schema.tmpl.sql template file and not the generated schema.sql as it will be overwritten
create database localdb;

create table localdb.user (
    `id` bigint unsigned not null auto_increment,
    `name` varchar(255) not null,
    `created_at` datetime NOT NULL DEFAULT CURRENT_TIMESTAMP,
    `updated_at` datetime NOT NULL DEFAULT CURRENT_TIMESTAMP ON UPDATE CURRENT_TIMESTAMP,
    PRIMARY KEY(`id`)
);
