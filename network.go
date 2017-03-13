package main

import (
	"database/sql"
	"errors"
	"fmt"
	"net"
)

func ParseRequest(message []byte, conn net.Conn, db *sql.DB) error {
	var cursor int
	var header = new(Header)

	if len(message) < 5 {
		return errors.New("Header too short")
	}

	headerSize := uint8(message[3])

	c, err := header.Unpack(message[0 : 3+headerSize+2])
	cursor += c

	if header == nil || err != nil {
		return errors.New("Header too short")
	}

	//dismiss payload for now

	cursor = cursor + 1

	switch header.Type {
	case 0x01:
		conn.Write(Version())
	case 0x10:
		c, _, _ := RequestPackage(message[cursor+1:], db)
		cursor += c
	default:
		fmt.Println(header.Type)
	}

	return nil
}
