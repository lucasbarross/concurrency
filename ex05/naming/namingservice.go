package naming

import (
	"middleware/clientproxy"
)

type NamingService struct {
	Repository map[string]clientproxy.ClientProxy
}

func (naming NamingService) Register(name string, proxy clientproxy.ClientProxy) bool {
	_, ok := naming.Repository[name]

	if ok {
		return false
	}

	naming.Repository[name] = proxy

	return true
}

func (naming NamingService) Lookup(name string) *clientproxy.ClientProxy {
	proxy, ok := naming.Repository[name]

	if ok {
		return &proxy
	}

	return nil
}

func (naming NamingService) List() map[string]clientproxy.ClientProxy {
	return naming.Repository
}
