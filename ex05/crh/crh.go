package crh

import (
	"log"
	"errors"
	"net"
	"strconv"
	"time"
	"bufio"
)

const (
	EOT_CHARACTER = '\n'
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
	msgToServer = append(msgToServer, EOT_CHARACTER)

	go func() {
		conn, err := net.Dial(crh.Protocol, crh.ServerHost+":"+strconv.Itoa(crh.ServerPort))
		if err != nil {
			log.Println("Dial")
			errChan <- err
			return
		}
		defer conn.Close()
		conn.Write(msgToServer)
		reader := bufio.NewReader(conn)
		buf, err := reader.ReadBytes(EOT_CHARACTER)
		if err != nil {
			log.Println("ReadBytes")
			errChan <- err
			return
		}

		resultChan <- buf
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
