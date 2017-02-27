package main

import (
	"errors"
	"net"
)

func ParseRequest(message []byte, conn net.Conn) error {
	var header *Header

	if len(message) < 5 {
		return errors.New("Header too short")
	}

	headerSize := uint8(message[3])

	header.Unpack(message[0 : 3+headerSize])

	switch header.Type {
	case 0x10:
		//handle request
	}

	return nil
}
