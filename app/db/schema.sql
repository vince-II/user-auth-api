CREATE TABLE IF NOT EXISTS users (
    id SERIAL PRIMARY KEY,
    first_name TEXT NOT NULL,
    last_name TEXT NOT NULL
);

CREATE TABLE IF NOT EXISTS user_logins (
    id SERIAL PRIMARY KEY,
    token TEXT NOT NULL,
    last_login TIMESTAMP NOT NULL,
    username varchar(255) NOT NULL,
    password varchar(255) NOT NULL
);

CREATE TABLE IF NOT EXISTS post (
    id SERIAL PRIMARY KEY,
    user_id INTEGER NOT NULL,
    content TEXT NOT NULL,
    created_at TIMESTAMP NOT NULL
);
