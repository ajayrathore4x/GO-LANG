package main

import "fmt"

func main(){
	numbers := []int{10,20,30}

	numbers=append(numbers,40,50)

	fmt.Println(numbers[1])

	fmt.Println(numbers[1:4])

	fmt.Println(len(numbers))

	for i:=0;i<len(numbers);i++{
		fmt.Println(numbers[i])
	}
}