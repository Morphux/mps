package main

type header struct {
	T      uint8
	Number uint8
}

type getPkg_ struct {
	H header

	ID         uint64
	State      uint8
	NameLen    uint16
	CategLen   uint16
	VersionLen uint16
}

type respPkg struct {
	H header

	ID               uint64
	CompTime         float64
	InstSize         float64
	ArchSize         float64
	NameLen          uint64
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

type respPkg_ struct {
	H header

	ID               uint64
	CompTime         float64
	InstSize         float64
	ArchSize         float64
	NameLen          uint64
	CategoryLen      uint16
	VersionLen       uint16
	ArchiveLen       uint16
	ChecksumLen      uint16
	DependenciesSize uint16
}

type resp_file_t struct {
	ID       uint64
	kind     uint8
	parentID uint64
	pathLen  uint16
	path     []byte
}

type resp_news_t struct {
	ID            uint64
	parentID      uint64
	authorLen     uint16
	authorMailLen uint16
	textLen       uint16
	author        []byte
	authorMail    []byte
	text          []byte
}
