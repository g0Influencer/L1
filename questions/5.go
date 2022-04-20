package main

import (
	"fmt"
	"unsafe"
)

/*
	5. Какой размер у структуры struct{}{}?
	0 байт
 */
type Empty struct {}

func main(){

	fmt.Println(unsafe.Sizeof(Empty{})) // 0
}
