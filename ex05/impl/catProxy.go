package impl

import (
	"middleware/requestor"
)

type CatProxy struct {
	Requestor requestor.Requestor
}

func (cat CatProxy) Echo(message string) {
	parameters := []interface{}{message}

	cat.Requestor.Invoke("Cat", "Echo", parameters)
}