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

	result := []byte{}
	for {
		buff := make([]byte, 1024)
		n, addr, err := connUDP.ReadFrom(buff)
		address = addr
		if err != nil {
			return err, nil
		}
		
		result = append(result, buff...)	
		if (buff[n-1] == '\n') {
			break
		}
	}
	
	result = removeEOFCharacter(result)

	return nil, result
}

func removeEOFCharacter(buffer []byte) []byte {
	result := buffer
	
	for i := 0; i < len(result); i++ {
		if (result[i] == '\n') {
			result[i] = 0
			return result
		}
	}
	
	return result
}

func (srh SRH_UDP) Send(msgToClient []byte) error {
	if connUDP == nil {
		return errors.New("connUDPection not found")
	}

	connUDP.WriteTo(msgToClient, address)
	connUDP.Close()
	return nil
}
