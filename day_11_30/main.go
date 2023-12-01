package main

import "fmt"

func main() {
	start()
}

func start() {
	//var ch = make(chan int, 5)
	//ch := make(chan int, 5)

	var ch chan int
	ch = make(chan int, 5)

	ch <- 15
	ch <- 16
	ch <- 17
	fmt.Printf("name: %v ,len:%v ,cap : %v\n", ch, len(ch), cap(ch))
	fmt.Println(<-ch)
}
