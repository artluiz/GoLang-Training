package main

import (
	"fmt"
	"time"
)

func main() {

	// cria um channel que comunica entre duas threads
	hello := make(chan string)

	go func() {
		time.Sleep(time.Second)
		hello <- "Hello World"
	}()

	select {
	// espera um valor ser atribuído a hello e recebido por x
	case x := <-hello:
		fmt.Println(x)
	// se não tiver um default ele espera até o case 1 receber algo
	default:
		fmt.Println("deu ruim")
	}
}
