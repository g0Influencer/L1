package main

import "fmt"

func main() {
	//Способы создания слайсов
	var slice1 []int
	var slice2 = []int{2, 3, 4}
	slice3 := []int{}
	slice4 := make([]int, 5)
	slice5 := slice2[1:]

	//Способы создания map
	var m1 map[int]int     //так делать опасно,т.к. map явялется ссылкой на хеш-таблицу,в данном случае она равна nil
	m1[1] = 1              // ошибка
	var m2 = map[int]int{} // так можно
	m2 := map[int]int{}
	m2 := make(map[string]int)

}
