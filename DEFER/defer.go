package main

import "fmt"

func main() {

	fmt.Println("Start")

	defer fmt.Println("This runs at the end")

	fmt.Println("Middle")
}