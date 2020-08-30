package main

import (
	"fmt"
	"time"
	"sync"
	"sync/atomic"
	"math/rand"
)

var rwLock = sync.RWMutex{}
var consCond = sync.NewCond(rwLock.RLocker())
var prodCond = sync.NewCond(&sync.Mutex{})
var val = 0
var counter int32 = 0
var wg sync.WaitGroup

func consumir() int {
	consCond.L.Lock()
	atomic.AddInt32(&counter, 1) // increment consumer count
	prodCond.Broadcast()
	consCond.Wait()
	atomic.AddInt32(&counter, -1)

	defer consCond.L.Unlock()
	defer wg.Done() // signal that it is done reading
	return val
}

func produzir(n int) {
	fmt.Printf("Producing %d\n", n)
	if counter > 0 {
		wg.Add(int(counter)) // adding consumers to WaitGroup
		val = n
		consCond.Broadcast()
		wg.Wait() // waiting for consumers
	} else {
		prodCond.L.Lock()
		prodCond.Wait() // wait for consumers
		prodCond.L.Unlock()
		produzir(n) // try produce again
	}
}

func main() {

	//case1()

	time.Sleep(100 * time.Millisecond)
	
	//case2()
	case3()
	// case4()

	for {
	}
}


func case1() {
	go func() {
		fmt.Printf("C1 %d\n", consumir())

	}()

	go func() {
		fmt.Printf("C2 %d\n", consumir())
	}()

	go func()  {
		time.Sleep(100 * time.Millisecond)
		fmt.Printf("producing %d\n", 13)

		produzir(13)

		fmt.Printf("produced %d\n", 13)

	}()
}

func case2() {
	go func ()  {
		fmt.Printf("producing %d\n", 23)

		produzir(23)

		fmt.Printf("produced %d\n", 23)

	}()

	time.Sleep(100 * time.Millisecond)
	go func() {
		fmt.Printf("C1 %d\n", consumir())
	}()
	
	time.Sleep(100 * time.Millisecond)
	go func() {
		fmt.Printf("C2 %d\n", consumir())
	}()
}

func case3() {
	go func() {
		fmt.Printf("C1 %d\n", consumir())
		time.Sleep(500 * time.Millisecond)
		fmt.Printf("C1 %d\n", consumir())

	}()
	
	go func() {
		fmt.Printf("C2 %d\n", consumir())
		time.Sleep(500 * time.Millisecond)
		fmt.Printf("C2 %d\n", consumir())
	}()

	go func ()  {
		time.Sleep(200 * time.Millisecond)
		produzir(23)
		produzir(50)
	}()

}

func case4() {
	for i := 0; i < 10; i++ {
		fmt.Printf("Starting C%d\n", i)
		v := i
		go func() {
			for {
				fmt.Printf("C%d %d\n",v, consumir())
				time.Sleep(time.Duration(rand.Intn(10000)) * time.Millisecond)
			}
		}()
	}
	time.Sleep(500 * time.Millisecond)

	go func (){
		for {
			produzir(rand.Intn(50))
		}
	}()
}