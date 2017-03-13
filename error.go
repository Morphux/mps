package main

//An error happened server side
const ERR_SERVER_FAULT uint8 = 0x1

//A packet send by the client is wrong
const ERR_MALFORMED_PACKET uint8 = 0x2

//A request send by the client find no result
const ERR_RES_NOT_FOUND uint8 = 0x3

type Error struct {
	Message

	ErrorType uint8
	ErrorLen  uint16
	Error     string
}

func (p *Error) PackError(err error, errortype uint8) ([]byte, error) {
	p.ErrorType = errortype
	p.Error = err.Error()
	p.ErrorLen = uint16(len(p.Error))

	return p.Pack()
}
