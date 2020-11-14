package requestor

type struct Packet {
	Req Request
	Res Response
}

type struct Request {
	ReqHeader RequestHeader
	ReqBody RequestBody
}

type struct RequestHeader {
	RequestId int
	Status int
}

type struct RequestBody {
	Body []interface{}
}

type struct ResponseBody {
	Body interface{}
}