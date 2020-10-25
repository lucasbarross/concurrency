package crh

import (
	"errors"
	"log"
	"net"
	"strconv"
	"time"
)

type CRH struct {
	ServerHost string
	ServerPort int
	Protocol   string
	Timeout    time.Duration
}

func (crh CRH) SendReceive(msgToServer []byte) (error, []byte) {
	resultChan := make(chan []byte, 1)
	errChan := make(chan error, 1)

	go func() {
		conn, err := net.Dial(crh.Protocol, crh.ServerHost+":"+strconv.Itoa(crh.ServerPort))
		if err != nil {
			errChan <- err
			return
		}
		defer conn.Close()
		conn.Write(msgToServer)

		// for {
		// 	log.Print("Oi")
		// 	n, err := conn.Read(buffer)
		// 	log.Print(n)
		// 	if n == 0 || err == io.EOF {
		// 		break
		// 	} else if err != nil {
		// 		errChan <- err
		// 		return
		// 	}
		// 	result = append(result, buffer...)
		// }

		buffer := make([]byte, 512)
		log.Print("read")
		_, err = conn.Read(buffer)
		log.Print("done")

		if err != nil {
			errChan <- err
			return
		}
		resultChan <- buffer
	}()

	select {
	case res := <-resultChan:
		return nil, res
	case <-time.After(crh.Timeout):
		return errors.New("timeout"), nil
	case err := <-errChan:
		return err, nil
	}
}
