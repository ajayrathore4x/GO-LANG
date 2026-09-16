package main

import "fmt"

func greet(){
	fmt.Println("helloo! Ajay")
}

func add(a int, b int) int {
    return (a + b)
}

func main(){
	greet()
	fmt.Println(add(7,7))
}