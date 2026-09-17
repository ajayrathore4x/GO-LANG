package main

import "fmt"
import "strings"

func main(){
	name := "Ajay Rathore"

	fmt.Println(len(name))
	fmt.Println(strings.ToUpper(name))
	fmt.Println(strings.ToLower(name))

	fmt.Println(strings.Contains(name, "Ajay"))
	fmt.Println(strings.Replace(name,"Rathore","Developer",1))

}