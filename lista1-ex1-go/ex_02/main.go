package main

import (
	"fmt"
	"sync"
	"time"
)
var rw = sync.RWMutex{}
var consumidorCond = sync.NewCond(rw.RLocker())
var produtorCond = sync.NewCond(&sync.Mutex{})
var val = -1
var consumersCount = 0

func consumir() int {
	consumidorCond.L.Lock()
	produtorCond.Broadcast()
	consumersCount += 1
	consumidorCond.Wait()
	return val
}

func produzir(n int) {
	produtorCond.L.Lock()
	if consumersCount > 0 {
		val = n
		consumidorCond.Broadcast()
		consumersCount = 0
		produtorCond.L.Unlock()
	} else {
		produtorCond.Wait()
		produzir(n)
	}
}

func main(){
	go func()  {
		fmt.Printf("C1 %d\n", consumir())
		fmt.Printf("C1 %d\n", consumir())
	}()

	go func() {
		fmt.Printf("C2 %d\n", consumir())
		fmt.Printf("C2 %d\n", consumir())
	}()

	go func() {
		time.Sleep(50 * time.Millisecond)
		produzir(10)
		produzir(20)
	}()

	// go func() {
	// 	fmt.Printf("Producing 10\n")
	// 	produzir(10)
	// 	fmt.Printf("Produced 10\n")
	// }()

	// time.Sleep(100 * time.Millisecond)
	
	// go func()  {
	// 	fmt.Printf("C1 Consuming\n")
	// 	fmt.Printf("C1 %d\n", consumir())
	// 	fmt.Printf("C1 Consumed\n")
	// }()
	
	// time.Sleep(100 * time.Millisecond)
	
	// go func()  {
	// 	fmt.Printf("C2 Consuming\n")
	// 	fmt.Printf("C2 %d\n", consumir())
	// 	fmt.Printf("C2 Consumed\n")
	// }()
	

	for {}
}
