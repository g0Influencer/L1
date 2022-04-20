package main

import "fmt"

//Проблема: Символы латиницы и некоторые специальные символы (ASCII) занимают 1 байт на 1 символ, а
//русские буквы уже занимают 2 байта на символ.
//Решение: конвертировать всю строку в rune

var justString string
func someFunc() {
	v := createHugeString(1 << 10)
	r := []rune(v)
	justString := string(r[:100])
	fmt.Println(justString)
}
func createHugeString(size int) string {
	var v string
	for i := 0; i < size; i++ {
		v += "К"
	}

	return v
}

func main() {
	someFunc()
}

