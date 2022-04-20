package main

import (
	"fmt"
	"sync"
	"time"
)

/*
	3. Чем отличаются RWMutex от Mutex?
	Методы Lock() и Unlock() из пакета Mutex обеспечивают экслюзивную блокировку горутины,
	тем самым помещая другие горутины в очередь и увеличивая при этом время выполнения программы.
	Однако при чтении данных это делать необязательно,в таком случае рекомендуется использовать методы  RLock() и RUnlock().
	Т.е. RWMutex следует использовать в том случае, когда мы только читаем данные и при этом ничего не записываем, так мы
	добьемся более быстрого выполнения работы горутины.

 */
func workWithRWMutex(){
	var (
		wg sync.WaitGroup
		mu sync.Mutex
		rw sync.RWMutex

	)
	start:= time.Now() // определяем текущее время
	counter:= 0

	wg.Add(100)

	for i := 0; i < 50; i++ {

		go func() { //горутина, которая только читает данные из счетчика
			defer wg.Done()
			rw.RLock()
			_ = counter
			time.Sleep(1 * time.Nanosecond)
			rw.RUnlock()
		}()

		go func() { //горутина, которая меняет значение счетчика
			defer wg.Done()
			mu.Lock()
			counter++
			time.Sleep(1 * time.Nanosecond)
			mu.Unlock()
		}()
	}
	wg.Wait()

	fmt.Println(counter)
	fmt.Println(time.Now().Sub(start).Seconds()) //рассчитываем время выполнения программы
}
func workWithoutRWMutex()  {
	var (
		wg sync.WaitGroup
		mu sync.Mutex
	)
	start:= time.Now()
	counter:= 0
	//wg.Add(100)

	for i := 0; i < 50; i++ {
		wg.Add(2)
		go func() {
			defer wg.Done()
			mu.Lock()
			_ = counter
			time.Sleep(1 * time.Nanosecond)
			mu.Unlock()
		}()

		go func() {
			defer wg.Done()
			mu.Lock()
			counter++
			time.Sleep(1 * time.Nanosecond)
			mu.Unlock()
		}()
	}
	wg.Wait()

	fmt.Println(counter)
	fmt.Println(time.Now().Sub(start).Seconds())
}
func main(){
	workWithRWMutex()
	workWithoutRWMutex()
}
