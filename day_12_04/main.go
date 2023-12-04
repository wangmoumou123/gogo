package main

import "day_12_04/tools"

func main() {
	//ch := make(chan int, 10)
	//for i := 0; i < 10; i++ {
	//	ch <- i
	//
	//}
	//
	////close(ch)
	////for i := range ch {
	////	fmt.Println(i)
	////
	////}
	//chLen := len(ch)
	//for i := 0; i < chLen; i++ {
	//	fmt.Println(<-ch)
	//
	//}
	//tools.Compute(30000)
	//tools.PipeTest()
	tools.SelectUse()

}
