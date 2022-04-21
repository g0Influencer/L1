package main

import (
	"context"
	"fmt"
	"time"
)

func work(ctx context.Context) {
	for {
		select {
		case <-ctx.Done():
			fmt.Println("Exit")
			return
		default:
			fmt.Println("working...")
		}
		time.Sleep(time.Millisecond)
	}

}

func main() {
	ctx, cancel := context.WithCancel(context.Background())
	go work(ctx)
	time.Sleep(2 * time.Millisecond)
	cancel()
}
