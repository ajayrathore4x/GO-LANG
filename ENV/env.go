package main

import (
    "fmt"
    "os"
)

func main() {

    port := os.Getenv("PORT")

    if port == "" {
        port = "8080"
    }

    fmt.Println("Server running on port:", port)
}