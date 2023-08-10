create table "users" (
    "id" bigserial primary key,
    "username" varchar(20) not null,
    "email" varchar not null,
    "password" varchar(60) not null
);