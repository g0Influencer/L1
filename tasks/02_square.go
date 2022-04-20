package main

import (
	"fmt"
	"sync"
)

func main(){
	arr:=[]int{2,4,6,8,10}


	wg:=new(sync.WaitGroup) // определяем группу горутин
	for _,val:=range arr{
		wg.Add(1) // увеличиваем счетчик горутин в группе
		go func(num int){
			defer wg.Done() // сообщаем, что элемент группы завершил работу
			fmt.Println(num*num)
		}(val)
	}
	wg.Wait() // ожидаем завершения всех горутин в группе

}
