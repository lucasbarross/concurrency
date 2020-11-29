package impl

import (
	"log"
	"middleware/marshaller"
	"middleware/protocol"
	"middleware/srh"
)

type CatInvoker struct {
	SRH        srh.SRH
	CatPool    CatPool
	Marshaller marshaller.Marshaller
}

func (invoker CatInvoker) Invoke() {
	for {
		err, requestBytes := invoker.SRH.Receive()

		if err != nil {
			log.Fatal("Error")
		}

		requestPacket := protocol.Packet{}
		err = invoker.Marshaller.Unmarshal(requestBytes, &requestPacket)
		if err != nil {
			log.Println(err)
			log.Fatal("Error unmarhalling")
		}

		objectKey := requestPacket.Req.ReqHeader.ObjectKey
		if objectKey != "Cat" {
			log.Fatal("Wrong object")
		}
		operation := requestPacket.Req.ReqHeader.Operation
		parameters := requestPacket.Req.ReqBody.Body

		object := invoker.CatPool.Get()
		log.Println(object)
		var response protocol.Response
		switch operation {
		case "Echo":
			message, ok := parameters[0].(string)
			if ok {
				result := object.Echo(message)

				response = protocol.Response{
					protocol.ResponseHeader{requestPacket.Req.ReqHeader.RequestId, 200},
					protocol.ResponseBody{result},
				}
			}
		default:
			response = protocol.Response{
				protocol.ResponseHeader{requestPacket.Req.ReqHeader.RequestId, 404},
				protocol.ResponseBody{nil},
			}
		}

		invoker.CatPool.Add(object)
		requestPacket.Res = response
		responseBytes, err := invoker.Marshaller.Marshal(requestPacket)
		if err != nil {
			log.Fatal("Error marshaling")
		}

		invoker.SRH.Send(responseBytes)
	}
}
