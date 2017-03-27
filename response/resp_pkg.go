package response

import (
	"fmt"

	"github.com/Morphux/mps/message"
	"github.com/Morphux/mps/vendors/Nyarum/barrel"
)

type RespPkg struct {
	message.Message

	ID               uint64
	CompTime         float32
	InstSize         float32
	ArchSize         float32
	State            uint8
	NameLen          uint16
	CategoryLen      uint16
	VersionLen       uint16
	ArchiveLen       uint16
	ChecksumLen      uint16
	DependenciesSize uint16
	Name             []byte
	Category         []byte
	Version          []byte
	Archive          []byte
	Checksum         []byte
	Dependencies     []uint64
}

func (p *RespPkg) Pack() ([]byte, error) {

	fmt.Printf("%#v\n", p)

	barrel := barrel.NewBarrel()
	load := barrel.Load(p, []byte{}, true)

	err := barrel.Pack(load)

	return barrel.Bytes(), err
}
