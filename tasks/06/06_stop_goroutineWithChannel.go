package main

import "fmt"

func main() {
	quit := make(chan bool)
	go func() {
		fmt.Println("Start")
		for {
			select {
			case quit <- true:
				fmt.Println("Stop")
				return
			}
		}
	}()
	<-quit
}
