package main

import "fmt"

func main() {
	// var age int

	// fmt.Print("enter your age :")
	// fmt.Scanln(&age)

	// if age >= 18 {
	// 	fmt.Println("you're eligible to vote")
	// } else {
	// 	fmt.Println("you're not eligible to vote ")
	// }

	// var marks int

	// fmt.Print("Enter your marks: ")
	// fmt.Scanln(&marks)

	// if marks >= 90 {
	// 	fmt.Println("A")
	// } else if marks >= 75 {
	// 	fmt.Println("B")
	// } else if marks >= 50 {
	// 	fmt.Println("C")
	// } else {
	// 	fmt.Println("Fail")
	// }

	// break
	// for i := 1; i <= 10; i++ {
    // if i == 5 {
    //     break
    // }
    // fmt.Println(i)
    // }

	//continue
	// for i := 1; i <= 5; i++ {
    // if i == 3 {
    //     continue
    // }
    // fmt.Println(i)
    // }

	day := 2

	switch day {
	case 1:
		fmt.Println("Monday")
	case 2:
		fmt.Println("Tuesday")
	case 3:
		fmt.Println("Wednesday")
	case 4:
		fmt.Println("thrusday")
	case 5:
		fmt.Println("friday")
	case 6:
		fmt.Println("saturday")
	case 7: 
		fmt.Println("sunday")			
	default:
		fmt.Println("Invalid day")
	}
}