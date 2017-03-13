package main

type RespFile struct {
	Message

	ID       uint64
	Type     uint8
	ParentID uint64
	PathLen  uint16
	Path     string
}
