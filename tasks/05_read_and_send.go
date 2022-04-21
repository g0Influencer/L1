package main

import (
	"context"
	"fmt"
	"time"
)

func send(ch chan int) {
	val := 0
	for {
		val++
		ch <- val
		time.Sleep(1 * time.Second)

	}
}

func main() {

	ch := make(chan int)

	ctx, cancel := context.WithTimeout(context.Background(), 10*time.Second) // контекст тайм-аута

	defer cancel() // отмена контекста после завершения программы
	go send(ch)

	for {
		select {
		case val := <-ch:
			fmt.Println(val)
		case <-ctx.Done():
			fmt.Println("The end")
			return
		}
	}

}
