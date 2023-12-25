package main

import (
    "time"
    "github.com/google/uuid"
    "github.com/TvGelderen/film-finder-api/internal/database"
)

type User struct {
    ID uuid.UUID `json:"id"`
    Email string `json:"email"`
    CreatedAt time.Time `json:"createdAt"`
    UpdatedAt time.Time `json:"updatedAt"`
    Name string `json:"name"`
}

func mapDbUserToUser(dbUser database.User) User {
    return User {
        ID: dbUser.ID,
        Email: dbUser.Email,
        CreatedAt: dbUser.CreatedAt,
        UpdatedAt: dbUser.UpdatedAt,
        Name: dbUser.Name,
    }
}
