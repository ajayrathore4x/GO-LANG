package main

import "fmt"
import "errors"

func divide(a, b float64) (float64, error) {

    if b == 0 {
        return 0, fmt.Errorf("cannot divide by zero")
    }

    return a / b, nil
}

func withdraw(balance, amount float64) (float64, error) {

	if amount > balance {
		return balance, errors.New("insufficient balance")
	}

	return balance - amount, nil
}

func main() {

    result, err := divide(10, 2)

    if err != nil {
        fmt.Println("Error:", err)
        return
    }

    fmt.Println("Result:", result)

	balance, err := withdraw(1000, 1200)

	if err != nil {
		fmt.Println("Error:", err)
		return
	}

	fmt.Println("Remaining balance:", balance)
}
