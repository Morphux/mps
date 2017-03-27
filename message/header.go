package message

import "github.com/Morphux/mps/vendors/Nyarum/barrel"

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

func (h *Header) Build(Type, number uint8, data []byte) {
	h.Type = Type
	h.NextPkgLen = 0
	h.NextPkg = ""

	h.Size = uint16(int(1+2+1+h.NextPkgLen+number) + len(data))

}

func (p *Header) Pack() ([]byte, error) {
	barrel := barrel.NewBarrel()
	load := barrel.Load(p, []byte{}, true)

	err := barrel.Pack(load)

	return barrel.Bytes(), err
}

//BuildHeader Is theorically faster than generating header and Packing it
func BuildHeader(Type, number uint8, data []byte) []byte {

	hash := ""
	length := len(data) + 4 + 1;

	header := []byte{Type, byte(length & 0xff), byte(length >> 8), byte(len(hash))}

	return header

}
