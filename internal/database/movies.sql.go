// Code generated by sqlc. DO NOT EDIT.
// versions:
//   sqlc v1.24.0
// source: movies.sql

package database

import (
	"context"

	"github.com/google/uuid"
)

const getUserSavedMovies = `-- name: GetUserSavedMovies :many
SELECT movie_id FROM saved_movies WHERE user_id = $1
`

func (q *Queries) GetUserSavedMovies(ctx context.Context, userID uuid.UUID) ([]int32, error) {
	rows, err := q.db.QueryContext(ctx, getUserSavedMovies, userID)
	if err != nil {
		return nil, err
	}
	defer rows.Close()
	var items []int32
	for rows.Next() {
		var movie_id int32
		if err := rows.Scan(&movie_id); err != nil {
			return nil, err
		}
		items = append(items, movie_id)
	}
	if err := rows.Close(); err != nil {
		return nil, err
	}
	if err := rows.Err(); err != nil {
		return nil, err
	}
	return items, nil
}

const saveMovie = `-- name: SaveMovie :one
INSERT INTO saved_movies(movie_id, user_id)
VALUES ($1, $2)
RETURNING id, movie_id, user_id
`

type SaveMovieParams struct {
	MovieID int32
	UserID  uuid.UUID
}

func (q *Queries) SaveMovie(ctx context.Context, arg SaveMovieParams) (SavedMovie, error) {
	row := q.db.QueryRowContext(ctx, saveMovie, arg.MovieID, arg.UserID)
	var i SavedMovie
	err := row.Scan(&i.ID, &i.MovieID, &i.UserID)
	return i, err
}