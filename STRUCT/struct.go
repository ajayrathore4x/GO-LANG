package main

import "fmt"

type student struct{
	name string
	age  int
	cgpa float64
}

func main(){
	student1:=student{
		name: "ajay",
		age:  21,
		cgpa: 7.2,
	}

	fmt.Println(student1.name)
	fmt.Println(student1.age)
	fmt.Println(student1.cgpa)

	student1.cgpa = 7.5

	fmt.Println(student1.cgpa)
}