CREATE TABLE IF NOT EXISTS users (
    id SERIAL PRIMARY KEY,
    username varchar(255) UNIQUE NOT NULL,
    password varchar(255) NOT NULL,
    first_name TEXT NOT NULL,
    last_name TEXT NOT NULL,
    last_login TIMESTAMP,
    last_logout TIMESTAMP,
    created_at TIMESTAMP NOT NULL DEFAULT now()
);


CREATE TABLE IF NOT EXISTS post (
    id SERIAL PRIMARY KEY,
    user_id INTEGER NOT NULL REFERENCES users(id) ON DELETE CASCADE,
    content TEXT NOT NULL,
    created_at TIMESTAMP NOT NULL DEFAULT now()
);
