package main

import (
	"log"
	"middleware/srh"
	"os"
)

func main() {
	args := os.Args[1:]
	protocol := args[0]

	var server srh.SRH
	if protocol == "tcp" {
		server = srh.SRH_TCP{
			ServerPort: 8080,
		}
	} else {
		server = srh.SRH_UDP{
			ServerPort: 8080,
		}
	}

	for {
		err, buff := server.Receive()
		if err != nil {
			log.Fatal(err)
		}
		msg := string(buff)
		log.Print(msg)

		err = server.Send(buff)
		if err != nil {
			log.Fatal(err)
		}
	}
}
