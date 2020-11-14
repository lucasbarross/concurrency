package main

import (
	"log"
	"middleware/crh"
	"middleware/requestor"
	"middleware/marshaller"
	"middleware/impl"
	"time"
)

func main() {
	crh := crh.CRH{
		ServerHost: "localhost",
		ServerPort: 8080,
		Protocol:   "tcp",
		Timeout:    time.Duration(30 * time.Second)}

	marshaller := marshaller.JsonMarshaller{}
	requestor := requestor.Requestor{
		Marshaller: marshaller,
		CRH: crh,
	}
	catProxy := impl.CatProxy{
		Requestor: requestor,
	}

	res, err := catProxy.Echo("hello")
	if err != nil {
		log.Fatal(err)
	}
	log.Print(res)
}
