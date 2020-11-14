package impl

import (
	"log"
	"middleware/srh"
	"middleware/marshaller"
	"middleware/protocol"
)

type struct CatInvoker {
	SRH srh.SRH
	Object Cat
	Marshaller marshaller.Marshaller
}

func (CatInvoker invoker) Invoke() {
	for {
		err, requestBytes := invoker.SRH.Receive()
		if err != nil {
			log.Fatal("Error")
		}

		requestPacket := protocol.Packet{}
		err := invoker.Marshaller.Unmarshal(requestBytes, requestPacket)
		if err != nil {
			log.Fatal("Error unmarhalling")
		}

		objectKey := requestPacket.ReqHeader.ObjectKey
		if objectKey != "Cat" {
			log.Fatal("Wrong object")
		}
		operation := requestPacket.ReqHeader.Operation
		parameters := requestPacket.ReqBody.Body

		protocol.Response response
		switch operation {
		case "Echo":
			result := invoker.Object.Echo(parameters[0])

			response = protocol.Response{
				ResponseHeader{requestPacket.ReqHeader.RequestId, 200}
				ResponseBody{result}
			}
		default:
			response = protocol.Response{
				ResponseHeader{requestPacket.ReqHeader.RequestId, 404}
				ResponseBody{nil}
			}
		}
		
		requestPacket.Res = response
		responseBytes, err := invoker.Marshaller.Marshal(requestPacket)
		if err != nil {
			log.Fatal("Error marshaling")
		}

		invoker.SRH.Send(responseBytes)
	}
}