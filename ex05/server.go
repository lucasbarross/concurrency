package main

import (
	"log"
	"middleware/srh"
	"middleware/impl"
	"middleware/marshaller"
)

func main() {
	srh := srh.SRH_TCP{
		ServerPort: 8080,
	}
	marshaller := marshaller.JsonMarshaller{}
	catInvoker := impl.CatInvoker{
		SRH: srh,
		Marshaller: marshaller,
		Object: impl.CatImpl{},
	}

	log.Println("Starting Cat service")
	catInvoker.Invoke()
}
