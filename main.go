package main

import (
    "fmt"
    "log"
    "os"
    "net/http"
    "github.com/joho/godotenv"
    "github.com/go-chi/chi/v5"
    "github.com/go-chi/cors"
)

func main() {
    godotenv.Load(".env")

    port := os.Getenv("PORT")

    if port == "" {
        log.Fatal("PORT missing")
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
    router.Get("/error", handlerError)

    server := &http.Server {
        Handler: router,
        Addr: ":" + port,
    }

    err := server.ListenAndServe()

    if err != nil {
        log.Fatal(err)
    }
}
