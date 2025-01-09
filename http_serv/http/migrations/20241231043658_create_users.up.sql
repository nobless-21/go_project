create table users(
    id bigserial primary key not null,
    email varchar not null unique,
    encrypted_password varchar not null
);