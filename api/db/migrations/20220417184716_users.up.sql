CREATE TABLE IF NOT EXISTS users(
    id bigserial primary key unique,
    username VARCHAR (50) unique,
    password VARCHAR,
    email VARCHAR (300) unique,
    created_at      timestamp,
    updated_at      timestamp
);