package main

import (
	"fmt"
	"log"
	"middleware/crh"
	"middleware/impl"
	"middleware/marshaller"
	"middleware/naming"
	"middleware/requestor"
	"os"
	"strconv"
	"strings"
	"sync"
	"time"
)

func ping(clientId int, protocol, addr, text string, wg *sync.WaitGroup) {
	defer wg.Done()

	s := strings.Split(addr, ":")
	host := s[0]
	port, err := strconv.Atoi(s[1])

	if err != nil {
		log.Fatal(err)
	}

	crh := crh.CRH{
		ServerHost: host,
		ServerPort: port,
		Protocol:   protocol,
		Timeout:    time.Duration(30 * time.Second),
	}

	marshaller := marshaller.JsonMarshaller{}
	requestor := requestor.Requestor{
		Marshaller: marshaller,
		CRH:        crh,
	}

	namingProxy := naming.NamingProxy{
		Requestor: requestor,
	}

	catPointer := *namingProxy.Lookup("Cat")
	catProxy := impl.CatProxy(catPointer)

	for i := 0; i < 10000; i++ {
		start := time.Now()
		_, err := catProxy.Echo(text)
		if err != nil {
			log.Fatal(err)
		}
		elapsed := time.Since(start)

		fmt.Printf("%d:%s\n", clientId, elapsed)
		time.Sleep(1000)
	}
}

func main() {
	args := os.Args[1:]
	threads, err := strconv.Atoi(args[0])
	protocol := args[1]
	addr := args[2]
	text := args[3]

	if err != nil {
		log.Fatal(err)
	}
	var wg sync.WaitGroup

	for i := 0; i < threads; i++ {
		wg.Add(1)
		go ping(i, protocol, addr, text, &wg)
	}

	wg.Wait()
}
