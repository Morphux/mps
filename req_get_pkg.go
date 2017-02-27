package main

import "github.com/Nyarum/barrel"

type GetPkg struct {
	ID         uint64
	State      uint8
	NameLen    uint16
	CategLen   uint16
	VersionLen uint16
	Name       []byte
	Category   []byte
	Version    []byte
}

func (p *GetPkg) Default() {

}

func (p GetPkg) Check(stats *barrel.Stats) bool {
	return true
}
