package srh

type SRH interface {
	Send(msgToClient []byte) error
	Receive() (error, []byte)
}
