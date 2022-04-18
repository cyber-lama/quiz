CREATE TABLE IF NOT EXISTS user_token(
     id bigserial primary key unique,
     user_id bigserial not null,
     FOREIGN KEY (user_id)
         REFERENCES users(id)
         ON DELETE CASCADE,
     token       varchar not null unique,
     created_at  timestamp not null,
     updated_at  timestamp not null
);