package main

type RespCat struct {
	Message

	ID       uint64
	ParentID uint64
	NameLen  uint64
	Name     string
}
