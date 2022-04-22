package main

import (
	"fmt"
	"time"
)

func Sleep(dur time.Duration) {
	timer := time.NewTimer(dur)
	<-timer.C // C - канал, который должен вернуть значение через заданное время
}

func main() {
	Sleep(5 * time.Second)
	fmt.Println("end")
}
