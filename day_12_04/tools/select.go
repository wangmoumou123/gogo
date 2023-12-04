package tools

import (
	"fmt"
	"time"
)

func SelectUse() {

	chInt := make(chan int, 10)
	chStr := make(chan string, 5)
	for i := 0; i < 10; i++ {
		chInt <- i
	}

	for i := 0; i < 5; i++ {
		chStr <- "ws123"
	}

	for {
		select {
		case i := <-chInt:
			fmt.Println(i)
			time.Sleep(time.Millisecond * 10)
		case str := <-chStr:
			fmt.Println(str)
			time.Sleep(time.Millisecond * 60)
		default:
			fmt.Println("all  chs  empty!")
			return
		}

	}
}
