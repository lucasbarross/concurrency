package main

import (
	"log"
	"middleware/crh"
	"os"
	"time"
)

func main() {
	args := os.Args[1:]
	protocol := args[0]

	var client crh.CRH
	if protocol == "tcp" {
		client = crh.CRH{
			ServerHost: "localhost",
			ServerPort: 8080,
			Protocol:   "tcp",
			Timeout:    time.Duration(30 * time.Second)}

	} else {
		client = crh.CRH{
			ServerHost: "localhost",
			ServerPort: 8080,
			Protocol:   "udp",
			Timeout:    time.Duration(30 * time.Second)}
	}

	msg := []byte("hello")
	err, res := client.SendReceive(msg)
	if err != nil {
		log.Fatal(err)
	}
	log.Print(string(res))
}
