package tools

import (
	"fmt"
	"sync"
	"time"
)

func init() {}

var wg sync.WaitGroup
var mtx sync.Mutex
var rwMutex sync.RWMutex

var num = 10
var m = make(map[int]int, 0)

func add() {
	defer wg.Done()
	num++
	time.Sleep(time.Millisecond)
	//fmt.Println(num)

}

func seq(num int) {
	defer wg.Done()
	mtx.Lock()
	var sum = 1
	for i := 1; i <= num; i++ {
		sum *= i
	}
	m[num] = sum
	mtx.Unlock()

}

func MultiTest() {

	for i := 0; i < 40; i++ {
		wg.Add(1)
		go seq(i)
	}

	wg.Wait()
	fmt.Println(m)
}

func RWTest() {

	for i := 1; i < 10; i++ {
		wg.Add(2)
		go writeMethod(i)
		go readMethod(i)

	}
	wg.Wait()

}

func writeMethod(i int) {
	defer wg.Done()
	rwMutex.Lock()
	num += i
	time.Sleep(time.Second * 2)
	fmt.Println("write====", num)
	rwMutex.Unlock()

}

func readMethod(i int) {
	defer wg.Done()
	rwMutex.RLock()
	time.Sleep(time.Second * 3)
	fmt.Println("read====", num)
	rwMutex.RUnlock()
}

func func1() {
	//其他操作...
	wg.Done()
}

func test() {
	for i := 1; i < 10; i++ {
		wg.Add(1)
		go func1()
	}
	wg.Wait()
}

//这样写有问题吗?
