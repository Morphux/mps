package request

import "github.com/Morphux/mps/message"

type ReqGetFile struct {
	message.Message

	ID      uint64
	PathLen uint16
	Path    string
}
