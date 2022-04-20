package main

import (
	"fmt"
	"strings"
)

func unique(s string) bool{
	s = strings.ToLower(s)
	for i:=range s{
		tmp:=string(s[i])
		count:= strings.Count(s,tmp)
		if count %2 == 0{
			return false
		}
	}
	return true

}


func main(){
	s:= "cC"

	fmt.Println(unique(s))

}
