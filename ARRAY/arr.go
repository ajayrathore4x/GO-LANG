package main

import "fmt"

func main(){
	numbers := [5]int{10,20,30,40,50}
	
	fmt.Println(numbers[0]) // 10
    fmt.Println(numbers[2]) // 30

	numbers[1]=25

	for i:=0;i<len(numbers);i++{
		fmt.Println(numbers[i])
	}
}