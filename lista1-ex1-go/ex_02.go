package main

import (
	"fmt"
	"math/rand"
	"time"
)

//func consumir(n, id int) {
//fmt.Printf("C%d: %d\n", id, n)
//}

//func consumidor(canal chan int) {
//for {
//n := <-canal
//for i := 0; i < 10; i++ {
//go consumir(n, i)
//}
//}
//}

func consumir(canal chan int) int {
	return <-canal
}

func consumidor(canal chan int, id int) {
	for {
		n := consumir(canal)
		fmt.Printf("C%d: %d\n", id, n)
	}
}

func produzir(canal chan int, n int) {
	canal <- n
}

func produtor(canal chan int) {
	for {
		n := rand.Intn(100)
		fmt.Printf("P: %d\n", n)
		go produzir(canal, n)
		time.Sleep(600 * time.Millisecond)
	}
}

func main() {
	canal := make(chan int)
	go produtor(canal)
	//go consumidor(canal)
	for i := 0; i < 10; i++ {
		go consumidor(canal, i)
	}
	for {
	}
}
