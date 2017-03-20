package response

import "github.com/Morphux/mps/message"

type AuthACK struct {
	message.Message

	MPMMajorVersion uint8
	MPMMinorVersion uint8
}

func Version() []byte {
	return []byte{1, 0}
}

func GetAuthACK() []byte {
	return append([]byte{0x02, 0x07, 0x00, 0x00, 0x01}, Version()...)
}
