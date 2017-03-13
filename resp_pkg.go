package main

type RespPkg struct {
	Message

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
	Name             string
	Category         string
	Version          string
	Archive          string
	Checksum         string
	Dependencies     []uint64
}
