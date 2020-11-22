package srh

import (
	"errors"
	"net"
	"strconv"
	"middleware/crh"
)

type SRH_UDP struct {
	ServerPort int
}

var connUDP net.PacketConn
var address net.Addr

func (srh SRH_UDP) Receive() (error, []byte) {
	connUDP, err = net.ListenPacket("udp", ":"+strconv.Itoa(srh.ServerPort))
	if err != nil {
		return err, nil
	}

	result := []byte{}
	for {
		buff := make([]byte, 1024)
		n, addr, err := connUDP.ReadFrom(buff)
		address = addr
		if err != nil {
			return err, nil
		}
		
		result = append(result, buff[:n]...)	
		if (buff[n-1] == crh.EOT_CHARACTER) {
			break
		}
	}

	return nil, result
}

func (srh SRH_UDP) Send(msgToClient []byte) error {
	if connUDP == nil {
		return errors.New("connUDPection not found")
	}

	if msgToClient[len(msgToClient)-1] != crh.EOT_CHARACTER{
		msgToClient = append(msgToClient, crh.EOT_CHARACTER)
	}

	connUDP.WriteTo(msgToClient, address)
	connUDP.Close()
	return nil
}
