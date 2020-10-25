package srh

import (
	"errors"
	"io"
	"log"
	"net"
	"strconv"
)

type SRH struct {
	Protocol   string
	ServerPort int
}

var ln net.Listener
var conn net.Conn
var err error

func (srh SRH) Receive() (error, []byte) {
	if ln == nil {
		ln, err = net.Listen(srh.Protocol, ":"+strconv.Itoa(srh.ServerPort))
		if err != nil {
			return err, nil
		}
	}

	conn, err = ln.Accept()
	if err != nil {
		return err, nil
	}

	result := []byte{}
	for {
		buffer := make([]byte, 512)
		log.Print("oi")
		n, err := conn.Read(buffer)
		log.Print(n)
		if n == 0 || err == io.EOF {
			break
		} else if err != nil {
			log.Fatal(err)
		}
		result = append(result, buffer...)
	}

	if err != nil {
		return err, nil
	}

	return nil, result
}

func (srh SRH) Send(msgToClient []byte) error {
	if conn == nil {
		return errors.New("Connection not found")
	}

	conn.Write(msgToClient)
	conn.Close()
	return nil
}
