
package main

import (
	"fmt"
	"net/http"
)

func homeHandler(w http.ResponseWriter, r *http.Request) {
	fmt.Fprintln(w, "Welcome to Ajay's backend")
}

func getUsers(w http.ResponseWriter, r *http.Request) {
	fmt.Fprintln(w, "User list: Ajay, Rahul, Aman")
}

func getPayment(w http.ResponseWriter, r *http.Request) {
	fmt.Fprintln(w, "Payment successful")
}

func main() {
	http.HandleFunc("/", homeHandler)
	http.HandleFunc("/users", getUsers)
	http.HandleFunc("/payment", getPayment)

	fmt.Println("Server running on :8080")

	err := http.ListenAndServe(":8080", nil)
	if err != nil {
		fmt.Println("Server error:", err)
	}
}
