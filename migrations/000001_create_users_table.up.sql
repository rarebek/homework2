CREATE TABLE IF NOT EXISTS users (
    id uuid primary key,
    name text,
    email text,
    password text,
    refresh_token text
);