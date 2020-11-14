package requestor

import (
	"middleware/marshaller"
	"middleware/crh"
	"middleware/protocol"
)

type struct Requestor {
	Marshaller marshaller.Marshaller
	CRH crh.CRH
}

func (requestor Requestor) Invoke(objectName string, methodName string, parameters []interface{}) interface{} {
	requestPacket := createRequestPacket(objectName, methodName, parameters)

	requestBytes, err := requestor.Marshaller.Marshal(requestPacket)
	if err != nil {
		return nil, err
	}
	err, responseBytes = requestor.CRH.SendReceive(requestBytes)
	if err != nil {
		return nil, err
	}

	responsePacket := protocol.Packet{}
	err := requestor.Marshaller.Unmarshal(responseBytes, responsePacket)
	if err != nil {
		return err
	}
	// checar request id é igual??
	// checar status da resposta???
	return responsePacket.Res.ResponseBody.OperationResult
}

func createRequestPacket(objectName string, methodName string, parameters []interface{}) protocol.Packet {
	return protocol.Packet{
		Request{
			RequestHeader{
				"uuid",
				true,
				objectName,
				methodName
			},
			RequestBody{
				parameters
			}
		}
		nil
	}
}