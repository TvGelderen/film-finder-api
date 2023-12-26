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

type ReturnUser struct {
    ID uuid.UUID `json:"id"`
    Name string `json:"name"`
    Email string `json:"email"`
    AccessToken string `json:"accessToken"`
}

func mapDbUserToReturnUser(dbUser database.User, accessToken string) ReturnUser {
    return ReturnUser {
        ID: dbUser.ID,
        Name: dbUser.Name,
        Email: dbUser.Email,
        AccessToken: accessToken,
    }
}
