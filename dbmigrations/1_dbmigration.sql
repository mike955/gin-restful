create table `account` (
    `id` bigint unsigned auto_increment,
    `username` varchar(255) not null,
    `password` char(32) not null,
    `created_at` datetime,
    `updated_at` datetime,
    primary key(`id`)
)ENGINE=InnoDB  DEFAULT CHARSET=utf8mb4 COMMENT='账户信息';
