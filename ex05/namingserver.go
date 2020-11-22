package main

import (
	"log"
	"middleware/clientproxy"
	"middleware/marshaller"
	"middleware/naming"
	"middleware/srh"
)

func main() {
	srh := srh.SRH_UDP{
		ServerPort: 8081,
	}
	marshaller := marshaller.JsonMarshaller{}
	namingInvoker := naming.NamingInvoker{
		SRH:        srh,
		Marshaller: marshaller,
		Object: naming.NamingService{
			Repository: make(map[string]clientproxy.ClientProxy),
		},
	}

	log.Println("Starting Naming service")
	namingInvoker.Invoke()
}
