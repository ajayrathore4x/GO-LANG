package main

import "fmt"

// struct embedded

type Address struct{
	City  string
	State string
}

type User struct{
	Name string
	Age  int
	Address
}

func main() {
	user := User{
		Name: "Ajay",
		Age:  21,
		Address: Address{
			City:  "Jaipur",
			State: "Rajasthan",
		},
	}

	fmt.Println(user.Name)
	fmt.Println(user.City)
	fmt.Println(user.State)
}