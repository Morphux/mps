package response

import "github.com/Morphux/mps/message"

type AuthACK struct {
	message.Message

	MPMMajorVersion uint8
	MPMMinorVersion uint8
}
