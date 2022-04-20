package main

import (
	"fmt"
	"math"
)

/*
	2. Что такое интерфейсы, как они применяются в Go?
	Интерфейс - это контракт,который определяет методы,которым должен удовлетворять тип,чтобы реализовывать этот интерфейс.
	Интерфейсы помогают уменьшить дублирование, то есть количество шаблонного кода.
	Пустой интерфейсный тип не описывает методы. У него нет правил. И поэтому любой объект удовлетворяет пустому интерфейсу.
 */

type Size interface { // разные типы реализуют один интерфейс
	Area() float64
}
type Circle struct {
	Radius float64
}
type Rectangle struct {
	Height float64
	Width float64
}
func (c Circle) Area() float64{ //тип Circle определяет метод Area(),следовательно он реализует интерфейс Size
	return math.Pi * math.Pow(c.Radius, 2)
}
func (r Rectangle) Area() float64{
	return r.Width * r.Height
}

func main(){
	person := make(map[string]interface{}, 0)//тип значения map-пустой интерфейс

	person["name"] = "Alice"
	person["age"] = 21
	person["height"] = 167.64

	fmt.Printf("%+v \n", person)


	c:=Circle{
		Radius: 10,
	}
	r:=Rectangle{
		Width: 3,
		Height: 4,
	}
	fmt.Println(c.Area())
	fmt.Println(r.Area())
}

