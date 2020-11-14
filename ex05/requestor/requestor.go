package requestor

import (
	"middleware/marshaller"
	"middleware/crh"
	"middleware/protocol"
)

type Requestor struct {
	Marshaller marshaller.Marshaller
	CRH crh.CRH
}

func (requestor Requestor) Invoke(objectName string, methodName string, parameters []interface{}) (interface{}, error) {
	requestPacket := createRequestPacket(objectName, methodName, parameters)

	requestBytes, err := requestor.Marshaller.Marshal(requestPacket)
	if err != nil {
		return nil, err
	}
	err, responseBytes := requestor.CRH.SendReceive(requestBytes)
	if err != nil {
		return nil, err
	}

	responsePacket := protocol.Packet{}
	err = requestor.Marshaller.Unmarshal(responseBytes, responsePacket)
	if err != nil {
		return nil, err
	}
	// checar request id é igual??
	// checar status da resposta???
	return responsePacket.Res.ResBody.OperationResult, nil
}

func createRequestPacket(objectName string, methodName string, parameters []interface{}) protocol.Packet {
	return protocol.Packet{
		protocol.Request{
			protocol.RequestHeader{
				"uuid",
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