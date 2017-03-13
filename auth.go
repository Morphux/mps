package main

type Auth struct {
	Message

	mpm_major_version uint8
	mpm_minor_version uint8
}

func Version() []byte {
	return []byte{0, 1}
}
