package invoker

import (
	"middleware/srh"
	"middleware/marshaller"
)

type Invoker interface {
	Invoke()
}