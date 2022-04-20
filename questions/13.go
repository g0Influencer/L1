package main

import "fmt"
/*
	13. Что выведет данная программа и почему?
 */

func someAction(v []int8, b int8) {
	v[0] = 100
	v = append(v, b) // емкости не хватает,будет создан новый срез
}

func main() {
	var a = []int8{1, 2, 3, 4, 5} // емкость равна 5
	someAction(a, 6)
	fmt.Println(a) // [100,2,3,4,5] -новый элемент не добавился,т.к. append вернула новый массив
}

