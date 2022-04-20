package main

import (
	"fmt"
	"time"
)

func send(ch chan int){
	val:=0
	for {
		val++
		ch <- val
		time.Sleep(1 * time.Second)

	}
}

func main(){

	ch:=make(chan int)
	defer close(ch)
	go send(ch)
	stop := time.NewTimer(5*time.Second)
	for {
		select {
		case val := <-ch:
			fmt.Println(val)
		case <-stop.C:
			fmt.Println("The end")
			return
		}
	}

}
