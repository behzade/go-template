create database {{ .schema }};

create table {{ .schema }}.user (
    `id` bigint unsigned not null auto_increment,
    `name` varchar(255) not null,
    `created_at` datetime NOT NULL DEFAULT CURRENT_TIMESTAMP,
    `updated_at` datetime NOT NULL DEFAULT CURRENT_TIMESTAMP ON UPDATE CURRENT_TIMESTAMP,
    PRIMARY KEY(`id`)
);
