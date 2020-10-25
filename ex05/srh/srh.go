package srh

import (
	"errors"
	"log"
	"net"
	"bufio"
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

	reader := bufio.NewReader(conn)
	buf, err := reader.ReadBytes('\n')
	if err != nil {
		return err, nil
	}

	return nil, buf
}

func (srh SRH) Send(msgToClient []byte) error {
	if conn == nil {
		return errors.New("Connection not found")
	}

	conn.Write(msgToClient)
	conn.Close()
	return nil
}
