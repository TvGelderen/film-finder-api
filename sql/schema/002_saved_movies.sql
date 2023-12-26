-- +goose Up
CREATE TABLE saved_movies (
    id SERIAL PRIMARY KEY,
    movie_id INT NOT NULL,
    user_id UUID NOT NULL REFERENCES users(id) ON DELETE CASCADE
);

-- +goose Down
DROP TABLE saved_movies;
