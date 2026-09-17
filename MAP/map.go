package main

import "fmt"

func main(){
	student := map[string]string{
		"name":   "Ajay",
		"age":    "21",
		"course": "Go",
	}

	fmt.Println(student["name"])
    
student["city"]="jaipur"

student["age"]="20"

delete(student,"course")
for key , value := range student{
	fmt.Println(key,value)
}
}