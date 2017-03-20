package request

import "github.com/Morphux/mps/message"

type Auth struct {
	message.Message

	mpm_major_version uint8
	mpm_minor_version uint8
}
