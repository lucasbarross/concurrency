package main

import (
	"fmt"
	"math/rand"
	"sync"
	"time"
)

func consumirCarros(sentido string, fila chan int, ponte *sync.Mutex) {
	for {
		hasLock := false
		ponte.Lock()
		fmt.Printf("A ponte está liberada para '%s'\n", sentido)
		hasLock = true
		go func() {
			time.Sleep(600 * time.Millisecond)
			hasLock = false
		}()
		for hasLock {
			if carro, ok := <-fila; ok {
				fmt.Printf("Carro %d está %s pela ponte\n", carro, sentido)
			} else {
				break
			}
		}
		fmt.Printf("A ponte está fechada para '%s'\n", sentido)
		ponte.Unlock()
	}
}

func produzirCarros(sentido string, fila chan int) {
	for {
		carro := rand.Intn(100)
		fmt.Printf("Carro %d está %s\n", carro, sentido)
		fila <- carro
		time.Sleep(100 * time.Duration(carro) * time.Millisecond)
	}
}

func main() {
	filaA := make(chan int, 300)
	filaB := make(chan int, 300)

	ponte := sync.Mutex{}

	go produzirCarros("ino", filaA)
	go produzirCarros("vino", filaB)
	go consumirCarros("ino", filaA, &ponte)
	go consumirCarros("vino", filaB, &ponte)
	for {
	}
}
