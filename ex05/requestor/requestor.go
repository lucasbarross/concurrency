package requestor

import (
	"errors"
	"log"
	"middleware/crh"
	"middleware/marshaller"
	"middleware/protocol"
	"strconv"

	"github.com/google/uuid"
)

type Requestor struct {
	Marshaller marshaller.Marshaller
	CRH        crh.CRH
}

func (requestor Requestor) Invoke(objectName string, methodName string, parameters []interface{}) (interface{}, error) {
	requestPacket := createRequestPacket(objectName, methodName, parameters)

	requestBytes, err := requestor.Marshaller.Marshal(requestPacket)
	if err != nil {
		log.Println("Marshal")
		return nil, err
	}
	err, responseBytes := requestor.CRH.SendReceive(requestBytes)
	if err != nil {
		log.Println("SendReceive")
		return nil, err
	}
	// log.Println(string(responseBytes))

	responsePacket := protocol.Packet{}
	err = requestor.Marshaller.Unmarshal(responseBytes, &responsePacket)
	if err != nil {
		return nil, err
	}

	if responsePacket.Req.ReqHeader.RequestId != responsePacket.Res.ResHeader.RequestId {
		return nil, errors.New("Unexpected requestId on response")
	}

	if responsePacket.Res.ResHeader.Status != 200 {
		return nil, errors.New("Unexpected response status " +
			strconv.Itoa(responsePacket.Res.ResHeader.Status))
	}

	// log.Println(string(responseBytes))

	return responsePacket.Res.ResBody.OperationResult, nil
}

func createRequestPacket(objectName string, methodName string, parameters []interface{}) protocol.Packet {
	return protocol.Packet{
		protocol.Request{
			protocol.RequestHeader{
				uuid.New().String(),
				true,
				objectName,
				methodName,
			},
			protocol.RequestBody{
				parameters,
			},
		},
		protocol.Response{
			protocol.ResponseHeader{
				"",
				0,
			},
			protocol.ResponseBody{
				nil,
			},
		},
	}
}
