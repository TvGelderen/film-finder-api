package main

import (
    "fmt"
    "log"
    "os"
    "database/sql"
    "net/http"
    "github.com/joho/godotenv"
    "github.com/go-chi/chi/v5"
    "github.com/go-chi/cors"
    "github.com/TvGelderen/film-finder-api/internal/database"

    _ "github.com/lib/pq"
)

type apiConfig struct {
    DB *database.Queries
}

func main() {
    godotenv.Load(".env")

    port := os.Getenv("PORT")
    if port == "" {
        log.Fatal("PORT is missing")
    }

    dbConnectionString := os.Getenv("DB_CONNECTION_STRING")
    if dbConnectionString == "" {
        log.Fatal("Database connection string is missing")
    }

    connection, err := sql.Open("postgres", dbConnectionString)
    if err != nil {
        log.Fatal("Unable to establish connection with database: ", err)
    }

    apiCfg := apiConfig {
        DB: database.New(connection),
    }

    fmt.Println("Server starting on port: ", port)

    router := chi.NewRouter()

    router.Use(cors.Handler(cors.Options {
        AllowedOrigins: []string { "https://*", "http://*" },
        AllowedMethods: []string { "GET", "POST", "PUT", "DELETE", "OPTIONS" },
        AllowedHeaders: []string { "*" },
        ExposedHeaders: []string { "Link" },
        AllowCredentials: false,
        MaxAge: 300,
    }))

    router.Get("/health", handlerSuccess)

    router.Get("/users", apiCfg.middlewareAuth(apiCfg.handlerGetUser))

    // Auth
    router.Post("/auth/register", apiCfg.handlerRegister)
    router.Post("/auth/login", apiCfg.handlerLogin)

    // Save movies
    router.Post("/movies/save", apiCfg.middlewareAuth(apiCfg.handlerSaveMovie))
    router.Get("/movies", apiCfg.middlewareAuth(apiCfg.handlerGetSavedMovies))

    server := &http.Server {
        Handler: router,
        Addr: ":" + port,
    }

    err = server.ListenAndServe()
    if err != nil {
        log.Fatal(err)
    }
}
