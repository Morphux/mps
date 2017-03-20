package request

import "github.com/Morphux/mps/message"

type ReqGetUPD struct {
	message.Message

	PkgsLen uint64
	Pkgs    []uint64
}
