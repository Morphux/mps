package response

import "github.com/Morphux/mps/message"

type RespNews struct {
	message.Message

	ID            uint64
	ParentID      uint64
	AuthorLen     uint16
	AuthorMailLen uint16
	TextLen       uint16
	Author        string
	AuthorMail    string
	Text          string
}
