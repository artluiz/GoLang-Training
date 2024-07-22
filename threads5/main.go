package main

import (
	"fmt"
	"time"
)

func main() {

	// cria um channel que comunica entre duas threads
	queue := make(chan int)

	go func() {
		i := 0
		for {
			// apenas preenche o queue quando ele é lido em outra thread
			queue <- i
			i++
		}
	}()

	// por o queue esperar ser lido para avançar na sua thread
	// sempre que x for iterado aqui o for da thread também avança
	for x := range queue {
		time.Sleep(time.Second)
		fmt.Println(x)
	}
}
