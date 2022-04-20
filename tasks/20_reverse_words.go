package main

import (
	"fmt"
	"strings"
)

func main(){
	s:= "snow dog sun"
	arr:=strings.Split(s," ")

	for i:=0; i<len(arr)/2; i++{

		arr[i],arr[len(arr)-1-i] = arr[len(arr)-1-i],arr[i]
	}


	fmt.Println(strings.Join(arr," "))

}
