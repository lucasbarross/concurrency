package main

import (
	"log"
	"middleware/srh"
)

func main() {
	server := srh.SRH{
		Protocol:   "tcp",
		ServerPort: 8080,
	}

	for {
		err, buff := server.Receive()
		if err != nil {
			log.Fatal(err)
		}
		msg := string(buff)
		log.Print(msg)

		log.Print("Send")
		err = server.Send(buff)
		log.Print("Log")
		if err != nil {
			log.Fatal(err)
		}
	}
}
