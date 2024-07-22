package main

import "fmt"

func main() {

	forever := make(chan string)

	go func() {
		x := true
		for {
			if x == true {
				continue
			}
		}
	}()

	fmt.Println("Aguardando para sempre")
	// O forever aguarda para sempre pois nada é associado a ele por ninguém
	<-forever
}
