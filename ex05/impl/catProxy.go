package impl

import (
	"middleware/requestor"
	"errors"
)

type CatProxy struct {
	Requestor requestor.Requestor
}

func (cat CatProxy) Echo(message string) (string, error) {
	parameters := []interface{}{message}

	res, err := cat.Requestor.Invoke("Cat", "Echo", parameters)
	if err != nil {
		return "", err
	}

	val, ok := res.(string)
	if !ok {
		err = errors.New("Unexpected type " + val)
	}

	return val, err
}