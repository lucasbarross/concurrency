package srh

import (
	"errors"
	"net"
	"strconv"
)

type SRH_TCP struct {
	ServerPort int
}

var ln net.Listener
var conn net.Conn
var err error

func (srh SRH_TCP) Receive() (error, []byte) {
	if ln == nil {
		ln, err = net.Listen("tcp", ":"+strconv.Itoa(srh.ServerPort))
		if err != nil {
			return err, nil
		}
	}
	conn, err = ln.Accept()
	if err != nil {
		return err, nil
	}

	buffer := make([]byte, 1024)
	_, err := conn.Read(buffer)

	if err != nil {
		return err, nil
	}

	return nil, buffer
}

func (srh SRH_TCP) Send(msgToClient []byte) error {
	if conn == nil {
		return errors.New("Connection not found")
	}

	conn.Write(msgToClient)
	conn.Close()
	return nil
}
