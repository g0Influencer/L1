package main

import (
	"fmt"
	"sync"
)
/*
	11. Что выведет данная программа и почему?
	После выполенения данного кода мы получим deadlock,т.к. передали объект WaitGroup по значению.
	Метод Done() уменьшает количество исполняемых горутин в группе,
	но по определению он должен вызываться от указателя на WaitGroup.

 */

func main() {
	wg :=  sync.WaitGroup{}
	for i := 0; i < 5; i++ {
		wg.Add(1)
		go func(wg sync.WaitGroup, i int){
			fmt.Println(i)
			defer wg.Done()
		}(wg,i)
	}
	wg.Wait()
	fmt.Println("exit")
}
