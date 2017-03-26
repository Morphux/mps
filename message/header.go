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

//BuildWithHeader Is theorically faster than generating header and Packing it
// func BuildWithHeader(Type, number uint8, data []byte) []byte {

// 	hash := ""

// 	header := []byte{Type, , len(hash), hash}

// 	return append()

// }
