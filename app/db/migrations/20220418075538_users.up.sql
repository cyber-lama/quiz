CREATE TABLE IF NOT EXISTS users(
    id bigserial primary key unique,
    username VARCHAR (50) unique null,
    password VARCHAR null,
    email VARCHAR (300) unique null,
    created_at      timestamp,
    updated_at      timestamp
);