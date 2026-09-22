package main

import "fmt"

func main() {

	fmt.Println("Start")

	defer fmt.Println("This runs at the end")
	defer fmt.Println("First")
	defer fmt.Println("Second")
	defer fmt.Println("Third")

	fmt.Println("Middle")
}