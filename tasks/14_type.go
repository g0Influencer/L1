package main

import "fmt"

func defineType(i interface{}){
	switch s:=i.(type) {
	case int:
		fmt.Println("Это целое число", s)
	case string:
		fmt.Println("Это строка", s)
	case bool:
		fmt.Println("Это булева переменная", s)
	case chan string:
		fmt.Println("Это строковый канал", s)
	default:
	}

}

func main(){
	s:="hello"
	v:=5
	b:=true
	ch:=make(chan string,0)

	defineType(s)
	defineType(v)
	defineType(b)
	defineType(ch)
}
