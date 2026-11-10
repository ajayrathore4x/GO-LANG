package main

import (
    "context"
    "database/sql"
    "fmt"
    "log"
    "os"
    "time"

    _ "github.com/jackc/pgx/v5/stdlib"
)

func main() {

    dbURL := os.Getenv("DATABASE_URL")

    if dbURL == "" {
        log.Fatal("DATABASE_URL is not set")
    }

    db, err := sql.Open("pgx", dbURL)
    if err != nil {
        log.Fatal(err)
    }
    defer db.Close()

    ctx, cancel := context.WithTimeout(
        context.Background(),
        5*time.Second,
    )
    defer cancel()

    err = db.PingContext(ctx)
    if err != nil {
        log.Fatal("Database connection failed:", err)
    }

    fmt.Println("PostgreSQL connected successfully!")
}
