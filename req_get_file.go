package main

type ReqGetFile struct {
	Message

	ID      uint64
	PathLen uint16
	Path    string
}
