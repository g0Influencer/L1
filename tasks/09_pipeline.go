package main

import (
	"fmt"
	"os"
	"sync"
)

func main() {
	arr := []int{2, 4, 6, 8, 10}
	firstChan := make(chan int)
	secondChan := make(chan int)
	wg:=new(sync.WaitGroup)

	wg.Add(3)
	go func(){
		defer wg.Done()
		for _,val:=range arr{
			firstChan<-val
		}
		close(firstChan)
	}()

	go func() {
		defer wg.Done()
		for val:=range firstChan{
			secondChan<-val*2
		}
		close(secondChan)
	}()

	go func(){
		defer wg.Done()
		for val:=range secondChan{
			fmt.Fprintln(os.Stdout,val)
		}
	}()



	wg.Wait()
}