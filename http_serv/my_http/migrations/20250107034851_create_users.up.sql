create table users(
    id bigserial primary key not null,
    Email varchar not null unique,
    encrypted_password varchar not null
);