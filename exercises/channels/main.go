package main

import(
	"fmt"
	"time"
)

func main() {
	c := make(chan int)

	go func(){
		time.Sleep(time.Second * 3)
		c <- 1 //push số 1 vào channel
	}()
	// <-c //đọc ra từ channel
	fmt.Println("Hello, playground")
	fmt.Println(<-c)
}