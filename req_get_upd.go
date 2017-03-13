package main

type ReqGetUPD struct {
	Message

	PkgsLen uint64
	Pkgs    []uint64
}
