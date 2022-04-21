package main

import (
	"fmt"
	"sync"
	"sync/atomic"
)

//Пакет atomic, судя по своему названию, позволяет выполнять атомарные операции с данными
//Атомарность обеспечивается функциями runtime_procPin / runtime_procUnpin.
//Данные функции обеспечивают то, что между ними планировщик GO не будет выполнять никакую другую горутину.
//Благодаря этому код между pin и unpin выполняется атомарно.

func main() {
	c := int32(0)
	n := 1000

	wg := sync.WaitGroup{}
	wg.Add(n)
	for i := 0; i < n; i++ {
		go func(i int) {

			atomic.AddInt32(&c, 1)
			wg.Done()
		}(i)
	}
	wg.Wait()

	fmt.Println(c) // 1000
}
