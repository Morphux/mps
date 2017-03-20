package message

import "github.com/Nyarum/barrel"

type Header struct {
	Message

	Type       uint8
	Size       uint16
	NextPkgLen uint8
	NextPkg    string
}

func (h *Header) Unpack(data []byte) (int, error) {

	barrel := barrel.NewBarrel()
	load := barrel.Load(h, data, false)

	err := barrel.Unpack(load)
	if err != nil {
		return 0, err
	}

	h.NextPkg = string(data[4 : 4+h.NextPkgLen])

	return 3 + int(h.NextPkgLen), nil
}
