package main

import (
	"fmt"
	"sync"
	"time"
)

var wg sync.WaitGroup
var num = 0

var mutix sync.Mutex

func add() {
	defer wg.Done()
	mutix.Lock()
	num++
	fmt.Println(num)
	time.Sleep(time.Millisecond)
	mutix.Unlock()
	//wg.Done()

}

func MultiTest() {

	for i := 0; i < 20; i++ {
		wg.Add(1)
		go add()
	}

	wg.Wait()

}
func main() {
	//tools.MultiTest()
	MultiTest()
}
