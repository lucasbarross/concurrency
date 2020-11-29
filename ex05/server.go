package main

import (
	"log"
	"middleware/clientproxy"
	"middleware/crh"
	"middleware/impl"
	"middleware/marshaller"
	"middleware/naming"
	"middleware/requestor"
	"middleware/srh"
	"os"
	"strconv"
	"time"
)

func main() {
	args := os.Args[1:]
	poolSize, err := strconv.Atoi(args[0])
	if err != nil {
		log.Fatal(err)
	}

	crh := crh.CRH{
		ServerHost: "localhost",
		ServerPort: 8081,
		Protocol:   "tcp",
		Timeout:    time.Duration(30 * time.Second)}

	marshaller := marshaller.JsonMarshaller{}
	requestor := requestor.Requestor{
		Marshaller: marshaller,
		CRH:        crh,
	}
	namingproxy := naming.NamingProxy{
		Requestor: requestor,
	}

	catProxy := createClientProxy()
	result := namingproxy.Register("Cat", catProxy)
	log.Println(result)

	srh := srh.SRH_TCP{
		ServerPort: 8080,
	}
	catInvoker := impl.CatInvoker{
		SRH:        srh,
		Marshaller: marshaller,
		CatPool:    impl.NewCatPool(poolSize),
	}

	log.Println("Starting Cat service")
	catInvoker.Invoke()
}

func createClientProxy() clientproxy.ClientProxy {
	crh := crh.CRH{
		ServerHost: "localhost",
		ServerPort: 8080,
		Protocol:   "tcp",
		Timeout:    time.Duration(30 * time.Second)}

	marshaller := marshaller.JsonMarshaller{}
	requestor := requestor.Requestor{
		Marshaller: marshaller,
		CRH:        crh,
	}
	catProxy := clientproxy.ClientProxy{
		Requestor: requestor,
	}

	return catProxy
}
