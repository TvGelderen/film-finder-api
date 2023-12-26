-- name: SaveMovie :one
INSERT INTO saved_movies(movie_id, user_id)
VALUES ($1, $2)
RETURNING *;

-- name: GetUserSavedMovies :many
SELECT movie_id FROM saved_movies WHERE user_id = $1;
