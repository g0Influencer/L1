package main

import "fmt"

/*
	14. Что выведет данная программа и почему?
	Функция append возвращает новый срез содержащий ранее содержавшиеся в срезе элементы,
	а также новые элементы, переданные в качестве аргумента функции

 */
func main() {
	slice := []string{"a", "a"} // создаем слайс

	func(slice []string) {
		slice = append(slice, "a") // append возвращает новый слайс,теперь он ссылается на новый массив

		slice[0] = "b" // изменяем новый слайс
		slice[1] = "b"
		fmt.Print(slice) // [b,b,a]
	}(slice)
	fmt.Print(slice)// [a,a] - старый слайс не изменился
}

