package main

import "fmt"

func main(){
	var (
	val int64
	i int
	)

	fmt.Scan(&val,&i)

	val |= 1<<i  // установить бит в 1
	fmt.Println(val)
	val &= ^(1<<i) // сбросить бит
	fmt.Println(val)


}
