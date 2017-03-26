package message

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

// Return Server Fault Message
func ServerFault() []byte {
	errorMessage := []byte("An error happened server side") //len 0x1D
	return append([]byte{0x03, 0x25, 0x00, 0x00, 0x01, 0x01, 0x1D, 0x00}, errorMessage...)
}

// Return Malformed Malformed Message
func MalformedPacket() []byte {
	errorMessage := []byte("A packet send by the client is wrong") //len 0x24
	return append([]byte{0x03, 0x2C, 0x00, 0x00, 0x01, 0x02, 0x24, 0x00}, errorMessage...)
}

// Return Ressource Not Found Message
func ResNotFound() []byte {
	errorMessage := []byte("A request send by the client find no result") //len 0x2B
	return append([]byte{0x03, 0x33, 0x00, 0x00, 0x01, 0x03, 0x2B, 0x00}, errorMessage...)
}

// Generate a Error Message
func GenerateError(Type uint8, message string) []byte {
	return append([]byte{Type, byte(len(message) + 8), 0x00, 0x00, 0x01, 0x03, byte(len(message)), 0x00}, []byte(message)...)
}
