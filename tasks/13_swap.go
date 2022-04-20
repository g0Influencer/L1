package main

import "fmt"

func main(){
	a,b:= 3,8
	fmt.Println(a, b)
	b,a = a,b
	fmt.Println(a, b)


}
