package message

import "github.com/Morphux/mps/vendors/Nyarum/barrel"

type Payload struct {
	Number uint8
}

func (p *Payload) Default() {

}
func (p Payload) Check(stats *barrel.Stats) bool {
	return true
}

func (p *Payload) Unpack(data []byte) error {
	barrel := barrel.NewBarrel()
	load := barrel.Load(p, data, false)

	err := barrel.Unpack(load)
	if err != nil {
		return err
	}

	return nil
}
