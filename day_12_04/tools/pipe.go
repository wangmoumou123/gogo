package tools

import (
	"fmt"
	"sync"
	"time"
)

func init() {}
func read(ch <-chan int, wg *sync.WaitGroup) {
	defer wg.Done()
	for i := range ch {
		//for i := range ch {
		fmt.Println(i)
		time.Sleep(time.Microsecond * 10)
	}
}

func write(ch chan<- int, nums int, chStatus chan bool, wg *sync.WaitGroup) {
	defer wg.Done()
	for i := 0; i < nums; i++ {
		ch <- i
		time.Sleep(time.Microsecond * 50)
	}
	//chStatus <- true
	close(ch)

}

func PipeTest() {
	var ch chan int
	var wg sync.WaitGroup
	ch = make(chan int, 100)
	chStatus := make(chan bool, 2)

	wg.Add(1)
	go read(ch, &wg)
	wg.Add(1)
	go write(ch, 100, chStatus, &wg)

	//wg.Add(1)
	//go func() {
	//	defer wg.Done()
	//	for i := 0; i < 1; i++ {
	//		<-chStatus
	//	}
	//	close(ch)
	//}()
	time.Sleep(time.Second * 2)
	wg.Wait()
	fmt.Println("stop")

}
