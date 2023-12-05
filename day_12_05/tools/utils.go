package tools

import (
	"fmt"
	"sync"
	"time"
)

func init() {}

var wg sync.WaitGroup
var num = 0

func add() {
	defer wg.Done()
	num++
	time.Sleep(time.Millisecond)
	fmt.Println(num)

}

func MultiTest() {

	for i := 0; i < 20; i++ {
		wg.Add(1)
		go add()
	}

	wg.Wait()

}
