package main

//#1 способ

//func main(){
//	quit:=make(chan bool)
//	go func(){
//		fmt.Println("Start")
//		for{
//			select {
//			case quit<-true:
//				fmt.Println("Stop")
//				return
//			}
//		}
//	}()
//	<-quit
//}



// #2 способ
//func main(){
//	quit:=make(chan int)
//	go func() {
//		n:=1
//		for {
//			select {
//			case quit<- n:
//				n++
//			case <-quit: // При попадании данных в канал выходим из горутины
//				fmt.Println("Stop")
//				return
//			}
//		}
//	}()
//	for i:=0; i<3; i++{
//		fmt.Println(<-quit)
//	}
//	quit<-0 // kill
//
//}
