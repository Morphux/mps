package response

import "github.com/Morphux/mps/message"

type RespFile struct {
	message.Message

	ID       uint64
	Type     uint8
	ParentID uint64
	PathLen  uint16
	Path     string
}
