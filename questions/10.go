package main

import "fmt"

/*
	10. Что выведет данная программа и почему?
	Функция update меняет значение указателя локально.Поэтому вывод в main не изменится
 */

func update(p *int) {
	b := 2
	p = &b

}

/* Для того, чтобы получить новое значение указателя, его следует вернуть из функции

func update(p *int) *int{
	b := 2
	p = &b
	return p
}*/

func main() {
	var (
		a = 1
		p = &a // p - указатель на a
	)
	fmt.Println(*p) //1
	update(p)
	fmt.Println(*p) // 1
}

