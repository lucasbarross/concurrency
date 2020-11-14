package protocol

type struct Packet {
	Req Request
	Res Response
}

type struct Request {
	ReqHeader RequestHeader
	ReqBody RequestBody
}

type struct RequestHeader {
	RequestId string
	ResponseExpected bool
	ObjectKey string
	Operation string
}

type struct RequestBody {
	Body []interface{}
}

type struct Response {
	ResHeader ResponseHeader
	ResBody ResponseBody
}

type struct ResponseHeader {
	RequestId int
	Status int
}

type struct ResponseBody {
	OperationResult interface{}
}