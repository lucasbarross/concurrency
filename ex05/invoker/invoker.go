package invoker

import (
	"middleware/marshaller"
	"middleware/srh"
)

type Invoker interface {
	Invoke()
}
