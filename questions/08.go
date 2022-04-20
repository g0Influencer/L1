package main

import "fmt"

/*
	8. В чем разница make и new?
	- new выделает место в памяти без инициализации
	- new возвращает указатель на переданный тип
	- make возвращает сам тип
	- make выделяет память и инциализирует только объекты типов: slice, map, chan.
 */

func main(){
	id := new(int)
	fmt.Printf("%T \n",id) // *int
	ch:=make(chan int)
	fmt.Printf("%T",ch) // chan int

}
