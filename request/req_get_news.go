package request

import "github.com/Morphux/mps/message"

type ReqGetNews struct {
	message.Message

	LastRequest uint32
	PkgsIDsSize uint16
	PkgsIDs     []uint64
}
