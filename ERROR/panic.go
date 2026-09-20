package main

import "fmt"

func main(){
	fmt.Println("starting")

	panic("something went wrong")

	fmt.Println("End")
}