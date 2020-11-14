package invoker

import (
	"middleware/srh"
	"middleware/marshaller"
)

type struct Invoker {
	SRH srh.SRH
	Marshaller marshaller.Marshaller
	ObjectMap map[string]RemoteObject
}

type interface RemoteObject {
	
}