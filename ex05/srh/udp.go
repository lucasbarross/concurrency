package srh

import (
	"errors"
	"net"
	"strconv"
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

	buff := make([]byte, 1024)
	_, address, err = connUDP.ReadFrom(buff)
	if err != nil {
		return err, nil
	}

	return nil, buff
}

func (srh SRH_UDP) Send(msgToClient []byte) error {
	if connUDP == nil {
		return errors.New("connUDPection not found")
	}

	connUDP.WriteTo(msgToClient, address)
	connUDP.Close()
	return nil
}
