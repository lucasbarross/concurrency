package naming

import (
	"log"
	"middleware/clientproxy"
	"middleware/marshaller"
	"middleware/protocol"
	"middleware/srh"
)

type NamingInvoker struct {
	SRH        srh.SRH
	Object     NamingService
	Marshaller marshaller.Marshaller
}

func (invoker NamingInvoker) Invoke() {
	for {
		err, requestBytes := invoker.SRH.Receive()
		// log.Println("Handling request ")
		// log.Println(string(requestBytes))

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
		if objectKey != "NamingService" {
			log.Fatal("Wrong object")
		}
		operation := requestPacket.Req.ReqHeader.Operation
		parameters := requestPacket.Req.ReqBody.Body

		var response protocol.Response
		switch operation {
		case "Register":
			name, okName := parameters[0].(string)
			proxy := clientproxy.FromMap(parameters[1].(map[string]interface{}))
			if okName {
				result := invoker.Object.Register(name, proxy)
				log.Println("RequestId: ", requestPacket.Req.ReqHeader.RequestId)

				response = protocol.Response{
					protocol.ResponseHeader{RequestId: requestPacket.Req.ReqHeader.RequestId, Status: 200},
					protocol.ResponseBody{result},
				}
			} else {
				log.Println("error paremeters type: ", parameters, okName)
			}
		case "Lookup":
			name, ok := parameters[0].(string)
			if ok {
				result := invoker.Object.Lookup(name)

				response = protocol.Response{
					protocol.ResponseHeader{requestPacket.Req.ReqHeader.RequestId, 200},
					protocol.ResponseBody{result},
				}
			}
		case "List":
			result := invoker.Object.List()

			response = protocol.Response{
				protocol.ResponseHeader{requestPacket.Req.ReqHeader.RequestId, 200},
				protocol.ResponseBody{result},
			}
		default:
			response = protocol.Response{
				protocol.ResponseHeader{requestPacket.Req.ReqHeader.RequestId, 404},
				protocol.ResponseBody{nil},
			}
		}

		requestPacket.Res = response
		responseBytes, err := invoker.Marshaller.Marshal(requestPacket)
		if err != nil {
			log.Fatal("Error marshaling")
		}

		// log.Println("Responding request ")
		// log.Println(string(responseBytes))
		invoker.SRH.Send(responseBytes)
	}
}
