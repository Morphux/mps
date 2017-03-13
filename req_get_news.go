package main

type ReqGetNews struct {
	Message

	LastRequest uint32
	PkgsIDsSize uint16
	PkgsIDs     []uint64
}
