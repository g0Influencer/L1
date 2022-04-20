package main

import (
	"fmt"
)


type Human struct {
	Name string
	Age int
}

func(h Human) Echo(){ // метод "класса" Human
	fmt.Println(h.Age)
}

type Action struct {
	 Human   // встраивание структуры
}


func main(){
	a:=Action{Human{Name: "Kirill",Age: 20}}
	fmt.Println(a.Name, a.Age)

}
