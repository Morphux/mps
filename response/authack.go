package response

import "github.com/Morphux/mps/message"

type AuthACK struct {
	message.Message

	MPMMajorVersion uint8
	MPMMinorVersion uint8
}

func Version() []byte {
	return []byte{0, 1}
}

func GetAuthACK() []byte {
	return append([]byte{0x2, 0x2, 0x0}, Version()...)
}
