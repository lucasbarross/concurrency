package main

import (
	"log"
	"middleware/crh"
	"time"
)

func main() {
	client := crh.CRH{
		ServerHost: "localhost",
		ServerPort: 8080,
		Protocol:   "udp",
		Timeout:    time.Duration(30 * time.Second)}

	msg := []byte("hello")
	err, res := client.SendReceive(msg)
	if err != nil {
		log.Fatal(err)
	}
	log.Print(string(res))
}
