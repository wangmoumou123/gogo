package tools

import (
	"fmt"
	"sync"
	"time"
)

var wg sync.WaitGroup

func init() {}

func putNum(intChan chan int, endless int) {
	defer wg.Done()
	for i := 0; i < endless; i++ {
		intChan <- i
	}
	close(intChan)

}

func computePrime(intChan chan int, primeChan chan int, exitChan chan bool) {
	defer wg.Done()
	for num := range intChan {
		flag := false
		for i := 2; i < num; i++ {
			if num%i == 0 {
				flag = false
				break
			}
			flag = true
		}
		if flag {
			primeChan <- num
		}
	}
	exitChan <- true

}

func printPrimeChan(primeChan chan int) {
	defer wg.Done()
	for _ = range primeChan {
		//fmt.Println(num)
	}
}

func Compute(endless int) {
	startTime := time.Now().UnixMilli()
	fmt.Println("compute")
	threads := 8
	intChan := make(chan int, 1000)
	primeChan := make(chan int, 50000)
	exitChan := make(chan bool, threads)

	wg.Add(1)
	go putNum(intChan, endless)

	wg.Add(threads)
	for i := 0; i < threads; i++ {
		go computePrime(intChan, primeChan, exitChan)
	}

	wg.Add(1)
	go printPrimeChan(primeChan)

	wg.Add(1)
	go func() {
		defer wg.Done()
		for i := 0; i < threads; i++ {
			<-exitChan
		}
		close(primeChan)

	}()

	wg.Wait()
	endTime := time.Now().UnixMilli()
	fmt.Printf("done====%v", endTime-startTime)
}
