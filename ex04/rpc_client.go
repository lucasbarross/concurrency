package main

import (
	"fmt"
	"log"
	"os"
	"strconv"
	"sync"
	"time"

	"github.com/google/uuid"
	"github.com/streadway/amqp"
)

func catRPC(clientId int, addr, text string, wg *sync.WaitGroup) (res string, err error) {
	defer wg.Done()

	conn, err := amqp.Dial(addr)
	if err != nil {
		log.Fatal(err)
	}
	defer conn.Close()

	ch, err := conn.Channel()
	if err != nil {
		log.Fatal(err)
	}
	defer ch.Close()

	q, err := ch.QueueDeclare(
		"",    // name
		false, // durable
		false, // delete when unused
		true,  // exclusive
		false, // noWait
		nil,   // arguments
	)
	if err != nil {
		log.Fatal(err)
	}

	msgs, err := ch.Consume(
		q.Name, // queue
		"",     // consumer
		true,   // auto-ack
		false,  // exclusive
		false,  // no-local
		false,  // no-wait
		nil,    // args
	)
	if err != nil {
		log.Fatal(err)
	}

	for i := 0; i < 10000; i++ {
		start := time.Now()

		corrId := uuid.New().String()
		err = ch.Publish(
			"",          // exchange
			"rpc_queue", // routing key
			false,       // mandatory
			false,       // immediate
			amqp.Publishing{
				ContentType:   "text/plain",
				CorrelationId: corrId,
				ReplyTo:       q.Name,
				Body:          []byte(text),
			})
		if err != nil {
			log.Fatal(err)
		}

		for d := range msgs {
			if corrId == d.CorrelationId {
				res = string(d.Body)
				break
			}
		}

		elapsed := time.Since(start)

		fmt.Printf("%d:%s\n", clientId, elapsed)
	}

	return
}

func main() {
	args := os.Args[1:]
	threads, err := strconv.Atoi(args[0])
	addr := args[1] // "amqp://guest:guest@localhost:5672/"
	text := args[2]

	if err != nil {
		log.Fatal(err)
	}
	var wg sync.WaitGroup

	for i := 0; i < threads; i++ {
		wg.Add(1)
		go catRPC(i, addr, text, &wg)
	}

	wg.Wait()
}
