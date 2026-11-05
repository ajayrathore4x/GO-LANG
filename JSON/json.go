package main

import (
    "encoding/json"
    "fmt"
)

type User struct {
    Name  string `json:"name"`
    Email string `json:"email"`
    Age   int    `json:"age"`
}

func main() {
    user := User{
        Name:  "Ajay",
        Email: "ajay@gmail.com",
        Age:   21,
    }

    data, err := json.Marshal(user)
    if err != nil {
        fmt.Println(err)
        return
    }

    fmt.Println(string(data))
}