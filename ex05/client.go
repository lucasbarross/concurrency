package main

import (
	"log"
	"middleware/crh"
	"middleware/impl"
	"middleware/marshaller"
	"middleware/naming"
	"middleware/requestor"
	"time"
)

func main() {
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

	namingProxy := naming.NamingProxy{
		Requestor: requestor,
	}

	catPointer := *namingProxy.Lookup("Cat")
	catProxy := impl.CatProxy(catPointer)

	res, err := catProxy.Echo("echo")
	if err != nil {
		log.Fatal(err)
	}
	log.Print(res)
}
