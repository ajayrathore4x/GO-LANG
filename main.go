package main

import "fmt"

func main() {
	var name string 
    age := 21
    var cgpa float64 = 7.2
    var isStudent bool = true

	fmt.Print("Enter your name: ")
    fmt.Scanln(&name)
    fmt.Println(name)
    fmt.Println(age)
    fmt.Println(cgpa)
    fmt.Println(isStudent)

	fmt.Printf("My name is %s, I am %d years old, my CGPA is %.2f and it is %t that I am a student.\n", name, age, cgpa, isStudent)

}
