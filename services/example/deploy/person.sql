create table if not exists t_person
(
    id         bigint auto_increment
        primary key,
    name       varchar(64)                        null,
    age        int      default 0                 not null,
    email      varchar(64)                        null,
    password   varchar(64)                        null,
    created_at datetime default CURRENT_TIMESTAMP not null,
    updated_at datetime default CURRENT_TIMESTAMP not null on update CURRENT_TIMESTAMP,
    deleted_at datetime                           null
);

