package main

import "fmt"

type Student struct{
	name  string
	marks int
}

func (s Student) getDetails(){
	fmt.Println(s.name)
	fmt.Println(s.marks)
}

func (s *Student) incMarks(number int){
	s.marks+=number
}

func main(){
	student := Student{
		name : "Ajay",
		marks: 95,
	}

	student.getDetails()
	student.incMarks(4)
	student.getDetails()
}