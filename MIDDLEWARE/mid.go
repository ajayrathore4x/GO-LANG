
package main

import (
	"fmt"
	"net/http"
)

// Home handler
func homeHandler(w http.ResponseWriter, r *http.Request) {
	fmt.Fprintln(w, "Home endpoint")
}

// Users handler
func getUsers(w http.ResponseWriter, r *http.Request) {
	fmt.Fprintln(w, "Users endpoint")
}

// Payment handler
func getPayment(w http.ResponseWriter, r *http.Request) {
	fmt.Fprintln(w, "Payment endpoint")
}

// Logging middleware
func loggingMiddleware(next http.Handler) http.Handler {
	return http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {

		fmt.Println("Request received:", r.URL.Path)

		next.ServeHTTP(w, r)
	})
}

func main() {

	// Create router
	mux := http.NewServeMux()

	// Register routes
	mux.HandleFunc("/", homeHandler)
	mux.HandleFunc("/users", getUsers)
	mux.HandleFunc("/payment", getPayment)

	// Attach middleware
	handler := loggingMiddleware(mux)

	// Start server
	fmt.Println("Server running on :8080")

	err := http.ListenAndServe(":8080", handler)
	if err != nil {
		fmt.Println("Server error:", err)
	}
}
