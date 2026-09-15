package main

import "fmt"

func main() {
	var age int

	fmt.Print("enter your age :")
	fmt.Scanln(&age)

	if age >= 18 {
		fmt.Println("you're eligible to vote")
	} else {
		fmt.Println("you're not eligible to vote ")
	}

}