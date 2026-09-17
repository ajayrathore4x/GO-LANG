package main

import "fmt"

func greet(){
	fmt.Println("helloo! Ajay")
}

func add(a int, b int) int {
    return (a + b)
}

func calculate(a, b int) (int, int) {
    sum := a + b
    difference := a - b

    return sum, difference
}
func main(){
	greet()
	fmt.Println(add(7,7))
	sum, difference := calculate(10, 3)

	fmt.Println(sum)
	fmt.Println(difference)
}