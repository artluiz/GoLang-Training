package main

import (
	"fmt"
	"time"
)

func worker(workerId int, msg chan int) {
	for res := range msg {
		fmt.Println("Worker: ", workerId, "Msg: ", res)
		time.Sleep(time.Second)
	}
}

func main() {

	msg := make(chan int)

	go worker(1, msg)
	go worker(2, msg)
	go worker(3, msg)
	go worker(4, msg)
	go worker(5, msg)
	go worker(6, msg)
	go worker(7, msg)
	go worker(8, msg)
	go worker(9, msg)
	go worker(10, msg)
	go worker(11, msg)
	go worker(12, msg)
	go worker(13, msg)
	go worker(14, msg)
	go worker(15, msg)
	go worker(16, msg)
	go worker(17, msg)
	go worker(18, msg)
	go worker(19, msg)
	go worker(20, msg)

	for i := 0; i < 101; i++ {
		msg <- i
	}
}
