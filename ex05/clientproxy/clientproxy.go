package clientproxy

import (
	"middleware/crh"
	"middleware/marshaller"
	"middleware/requestor"
	"time"
)

type ClientProxy struct {
	Requestor requestor.Requestor
}

func FromMap(mapVar map[string]interface{}) ClientProxy {
	requestorMap := mapVar["Requestor"]

	// Before you judge me walk a mile in my shoes
	crhMap := requestorMap.(map[string]interface{})["CRH"].(map[string]interface{})

	crh := crh.CRH{
		ServerHost: crhMap["ServerHost"].(string),
		ServerPort: int(crhMap["ServerPort"].(float64)),
		Protocol:   crhMap["Protocol"].(string),
		Timeout:    time.Duration(crhMap["Timeout"].(float64)),
	}

	return ClientProxy{
		Requestor: requestor.Requestor{
			Marshaller: marshaller.JsonMarshaller{},
			CRH:        crh,
		},
	}
}
