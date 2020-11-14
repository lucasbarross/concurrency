package protocol

type Packet struct {
	Req Request
	Res Response
}

type Request struct {
	ReqHeader RequestHeader
	ReqBody RequestBody
}

type RequestHeader struct {
	RequestId string
	ResponseExpected bool
	ObjectKey string
	Operation string
}

type RequestBody struct {
	Body []interface{}
}

type Response struct {
	ResHeader ResponseHeader
	ResBody ResponseBody
}

type ResponseHeader struct {
	RequestId string
	Status int
}

type ResponseBody struct {
	OperationResult interface{}
}