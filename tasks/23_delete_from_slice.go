package main

import "fmt"

func main(){
	arr:=[]int{1,3,5,6}
	i:=1
	copy(arr[i:],arr[i+1:])
	arr[len(arr)-1] = 0
	arr = arr[:len(arr)-1]
	fmt.Println(arr)
}
