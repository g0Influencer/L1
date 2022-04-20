package main

import (
	"fmt"
	"strings"
)

/* 1. Какой самый эффективный способ конкатенации строк?
 С помощью strings.Builder.
 Builder используется для эффективного создания строки с использованием методов Write.
 Это сводит к минимуму копирование памяти. */

func main(){
	var sb strings.Builder
	for i := 0; i < 1000; i++ {
		sb.WriteString("a")
	}

	fmt.Println(sb.String())
}

