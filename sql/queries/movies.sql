-- name: SaveMovie :one
INSERT INTO saved_movies(movie_id, user_id)
VALUES ($1, $2)
RETURNING *;

-- name: RemoveMovie :exec
DELETE FROM saved_movies WHERE movie_id = $1 AND user_id = $2;

-- name: GetUserSavedMovies :many
SELECT movie_id FROM saved_movies WHERE user_id = $1;
