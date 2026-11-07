
package main

import (
    "encoding/json"
    "net/http"
    "strconv"
)

type User struct {
    ID   int    `json:"id"`
    Name string `json:"name"`
}

var users = []User{
    {ID: 1, Name: "Ajay"},
    {ID: 2, Name: "Rahul"},
}

// GET /users
func getUsers(w http.ResponseWriter, r *http.Request) {
    w.Header().Set("Content-Type", "application/json")
    json.NewEncoder(w).Encode(users)
}

// POST /users
func createUser(w http.ResponseWriter, r *http.Request) {
    var user User

    err := json.NewDecoder(r.Body).Decode(&user)
    if err != nil {
        http.Error(w, "Invalid JSON", http.StatusBadRequest)
        return
    }

    user.ID = len(users) + 1
    users = append(users, user)

    w.Header().Set("Content-Type", "application/json")
    w.WriteHeader(http.StatusCreated)
    json.NewEncoder(w).Encode(user)
}

// DELETE /users/1
func deleteUser(w http.ResponseWriter, r *http.Request) {
    id, err := strconv.Atoi(r.PathValue("id"))
    if err != nil {
        http.Error(w, "Invalid user ID", http.StatusBadRequest)
        return
    }

    for i, user := range users {
        if user.ID == id {
            users = append(users[:i], users[i+1:]...)
            w.WriteHeader(http.StatusNoContent)
            return
        }
    }

    http.Error(w, "User not found", http.StatusNotFound)
}

func main() {
    mux := http.NewServeMux()

    mux.HandleFunc("GET /users", getUsers)
    mux.HandleFunc("POST /users", createUser)
    mux.HandleFunc("DELETE /users/{id}", deleteUser)

    http.ListenAndServe(":8080", mux)
}
