package main

import "github.com/Nyarum/barrel"

type Header struct {
	Type       uint8
	Size       uint16
	NextPkgLen uint8
	NextPkgs   string
}

func (p *Header) Default() {

}
func (p Header) Check(stats *barrel.Stats) bool {
	return true
}

func (h *Header) Unpack(data []byte) error {
	barrel := barrel.NewBarrel()
	load := barrel.Load(h, data, false)

	err := barrel.Unpack(load)
	if err != nil {
		return err
	}

	return nil
}
