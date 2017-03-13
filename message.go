package main

import "github.com/Nyarum/barrel"

//Message is the basic structure to send and receive Message all the methods can be overwritted if needed
type Message struct {
}

func (p *Message) Default() {
	//default value
}

func (p Message) Check(stats *barrel.Stats) bool {

	return true
}

func (p *Message) Unpack(data []byte) error {
	barrel := barrel.NewBarrel()
	load := barrel.Load(p, data, false)

	err := barrel.Unpack(load)

	return err
}

func (p *Message) Pack() ([]byte, error) {
	barrel := barrel.NewBarrel()
	load := barrel.Load(p, []byte{}, true)

	err := barrel.Pack(load)

	return barrel.Bytes(), err
}
