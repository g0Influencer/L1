package main

import (
	"fmt"
	"strings"
)

func unique(s string) bool {
	b := true
	mp := make(map[string]int)
	s = strings.ToLower(s)
	//Проходимся по каждому символу в строке и пишем в мапу
	for _, val := range s {
		if _, ok := mp[string(val)]; ok {
			return false
		} else {
			mp[string(val)]++
		}
	}
	return b

}

func main() {
	s := "ccc"

	fmt.Println(unique(s))

}
