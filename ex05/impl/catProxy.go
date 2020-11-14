package impl

import {
	"middleware/requestor"
}

type struct CatProxy {
	Requestor requestor.Requestor
}

func (CatProxy cat) Echo(string message) {
	cat.Requestor.Invoke("Cat", "Echo", [message])
}