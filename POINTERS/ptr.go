package main

import "fmt"

func main(){
	number:=10
	
	ptr:=&number

	fmt.Println(ptr)

	*ptr=50

	fmt.Println(number)
}