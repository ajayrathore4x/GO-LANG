
package main

import (
	"context"
	"fmt"
	"time"
)

func main() {
	ctx, cancel := context.WithTimeout(
		context.Background(),
		2*time.Second,
	)
	defer cancel()

	select {
	case <-time.After(5 * time.Second):
		fmt.Println("Payment completed")

	case <-ctx.Done():
		fmt.Println("Payment request timed out:", ctx.Err())
	}
}
