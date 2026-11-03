
package main

import (
	"context"
	"fmt"
	"time"
)

func worker(ctx context.Context) {
	select {
	case <-time.After(5 * time.Second):
		fmt.Println("Work completed")

	case <-ctx.Done():
		fmt.Println("Worker stopped:", ctx.Err())
	}
}

func main() {
	ctx, cancel := context.WithCancel(context.Background())

	go worker(ctx)

	time.Sleep(1 * time.Second)
	cancel()

	time.Sleep(100 * time.Millisecond)
}
