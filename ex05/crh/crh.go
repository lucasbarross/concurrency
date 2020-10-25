package crh

import (
	"errors"
	"log"
	"net"
	"strconv"
	"time"
	"bufio"
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
	msgToServer = append(msgToServer, "\n")

	go func() {
		conn, err := net.Dial(crh.Protocol, crh.ServerHost+":"+strconv.Itoa(crh.ServerPort))
		if err != nil {
			errChan <- err
			return
		}
		defer conn.Close()
		conn.Write(msgToServer)

		reader := bufio.NewReader(conn)
		buf, err := reader.ReadBytes('\n')
		if err != nil {
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
