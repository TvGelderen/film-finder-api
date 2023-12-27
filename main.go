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
    "github.com/TvGelderen/film-finder-api/handlers"

    _ "github.com/lib/pq"
)

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

    apiCfg := handlers.ApiConfig {
        DB: database.New(connection),
    }

    fmt.Println("Server starting on port: ", port)

    router := chi.NewRouter()

    router.Use(cors.Handler(cors.Options {
        AllowedOrigins: []string { "https://*", "http://*" },
        AllowedMethods: []string { "GET", "POST", "PUT", "DELETE", "OPTIONS" },
        AllowedHeaders: []string { "*" },
        ExposedHeaders: []string { "Link" },
        AllowCredentials: true,
        MaxAge: 300,
    }))

    router.Get("/health", handlers.HandlerSuccess)

    router.Get("/users", apiCfg.MiddlewareAuth(apiCfg.HandlerGetUser))

    // Auth
    router.Post("/auth/register", apiCfg.HandlerRegister)
    router.Post("/auth/login", apiCfg.HandlerLogin)
    router.Post("/auth/logout", apiCfg.MiddlewareAuth(apiCfg.HandlerLogout))

    // Save movies
    router.Get("/movies", apiCfg.MiddlewareAuth(apiCfg.HandlerGetSavedMovies))
    router.Post("/movies", apiCfg.MiddlewareAuth(apiCfg.HandlerSaveMovie))
    router.Delete("/movies", apiCfg.MiddlewareAuth(apiCfg.HandlerRemoveMovie))

    server := &http.Server {
        Handler: router,
        Addr: ":" + port,
    }

    err = server.ListenAndServe()
    if err != nil {
        log.Fatal(err)
    }
}
