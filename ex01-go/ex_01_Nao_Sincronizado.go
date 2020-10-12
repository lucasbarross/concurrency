package main

import (
	"fmt"
	"math/rand"
	"time"
)

func consumirCarros(sentido string, fila chan int) {
	for {
		carro := <-fila
		fmt.Printf("Carro %d está %s pela ponte\n", carro, sentido)
		time.Sleep(200 * time.Millisecond)
		fmt.Printf("Carro %d terminou a passagem pela ponte\n", carro)
	}
}

func produzirCarros(sentido string, fila chan int) {
	for {
		carro:= rand.Intn(100)
		fmt.Printf("Carro %d está %s\n", carro, sentido)
		fila <- carro
		time.Sleep(100 * time.Duration(10) * time.Millisecond)
	}
}

func main() {
	filaA := make(chan int, 300)
	filaB := make(chan int, 300)

	go produzirCarros("ino", filaA)
	go produzirCarros("vino", filaB)
	go consumirCarros("ino", filaA)
	go consumirCarros("vino", filaB)
	for {}
}
