package request

import "github.com/Morphux/mps/message"

type ReqGetCat struct {
	message.Message

	CategoriesLen uint16
	Categories    []uint64
}
