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

	requestor := requestor.Requestor{
		Marshaller: marshaller
	}
	catProxy := impl.CatProxy{}
}
