-- +goose Up
CREATE TABLE users (
    id UUID PRIMARY KEY,
    name TEXT NOT NULL,
    email TEXT UNIQUE NOT NULL,
    password_hash BYTEA,
    created_at TIMESTAMP NOT NULL,
    updated_at TIMESTAMP NOT NULL
);

CREATE TABLE saved_movies (
    id SERIAL PRIMARY KEY,
    movie_id INT NOT NULL,
    user_id UUID NOT NULL REFERENCES users(id) ON DELETE CASCADE
);

ALTER TABLE saved_movies
ADD CONSTRAINT uq_saved_movies UNIQUE(movie_id, user_id);

-- +goose Down
DROP TABLE saved_movies;
DROP TABLE users;
