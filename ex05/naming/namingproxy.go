package naming

import (
	"log"
	"middleware/clientproxy"
	"middleware/requestor"
)

type NamingProxy struct {
	Requestor requestor.Requestor
}

func (proxy NamingProxy) Register(name string, clientproxy clientproxy.ClientProxy) bool {
	parameters := []interface{}{name, clientproxy}

	result, err := proxy.Requestor.Invoke("NamingService", "Register", parameters)

	if err != nil {
		log.Println(err)
		return false
	}

	return result.(bool)
}

func (proxy NamingProxy) Lookup(name string) *clientproxy.ClientProxy {
	parameters := []interface{}{name}

	result, err := proxy.Requestor.Invoke("NamingService", "Lookup", parameters)
	resultMap, ok := result.(map[string]interface{})

	if err != nil || !ok {
		return nil
	}

	clientproxy := clientproxy.FromMap(resultMap)

	return &clientproxy
}

func (proxy NamingProxy) List() map[string]clientproxy.ClientProxy {
	result, err := proxy.Requestor.Invoke("NamingService", "List", nil)

	resultMap, ok := result.(map[string]interface{})

	if err != nil || !ok {
		return nil
	}

	returnMap := make(map[string]clientproxy.ClientProxy)
	for key, element := range resultMap {
		returnMap[key] = clientproxy.FromMap(element.(map[string]interface{}))
	}

	return returnMap
}
