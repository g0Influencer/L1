package main

import (
	"fmt"
	"math/rand"
	"os"
	"os/signal"
	"sync"
	"syscall"
)


func printData(ch chan int, wg *sync.WaitGroup){
	for val:=range ch{
		_,err:= fmt.Fprintln(os.Stdout,val)
		if err!=nil{
			return
		}
	}
	wg.Done()
}

func main(){
	wg:=sync.WaitGroup{}
	var N int
	fmt.Scan(&N)
	sigCh:=make(chan os.Signal,1)
	ch:=make(chan int)
	signal.Notify(sigCh,syscall.SIGINT,syscall.SIGTERM)

	for i:=0;i<N;i++{
		wg.Add(1)
		go printData(ch,&wg)
	}

	func() {
		for {
			select{
			case <-sigCh:
				fmt.Println("Exit")
				return
			default:
				ch <- rand.Int()

			}
		}
	}()
	close(sigCh)
	close(ch)
	wg.Wait()

}
