package response

import "github.com/Morphux/mps/message"

type RespCat struct {
	message.Message

	ID       uint64
	ParentID uint64
	NameLen  uint64
	Name     string
}
