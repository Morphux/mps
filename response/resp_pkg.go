package response

import (
	"fmt"

	"github.com/Morphux/mps/message"
	"github.com/Morphux/mps/vendors/Nyarum/barrel"
)

type RespPkg struct {
	message.Message

	ID               uint64
	CompTime         float64
	InstSize         float64
	ArchSize         float64
	State            uint8
	NameLen          uint64
	CategoryLen      uint16
	VersionLen       uint16
	ArchiveLen       uint16
	ChecksumLen      uint16
	DependenciesSize uint16
	Name             string
	Category         string
	Version          string
	Archive          string
	Checksum         string
	Dependencies     []uint64
}

func (p *RespPkg) Pack() ([]byte, error) {

	fmt.Printf("%#v\n", p)

	barrel := barrel.NewBarrel()
	load := barrel.Load(p, []byte{}, true)

	err := barrel.Pack(load)

	return barrel.Bytes(), err
}
