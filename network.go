package main

import (
	"database/sql"
	"errors"
	"fmt"
	"net"

	"github.com/Morphux/mps/message"
	"github.com/Morphux/mps/response"
)

func ParseRequest(data []byte, conn net.Conn, db *sql.DB) error {
	var cursor int
	var header = new(message.Header)

	if len(data) < 5 {
		return errors.New("Header too short")
	}

	headerSize := uint8(data[3])

	c, err := header.Unpack(data[0 : 3+headerSize+2])
	cursor += c

	if header == nil || err != nil {
		return errors.New("Header too short")
	}

	//payload
	cursor += 1

	switch header.Type {
	case 0x01:
		conn.Write(response.GetAuthACK())
	case 0x10:

		c, pkg, err := RequestPackage(data[cursor+1:], db)

		if err != nil {
			conn.Write(message.MalformedPacket())
			return err
		}

		resp, err := PkgtoRespPkg(pkg)

		if err != nil {
			return err
		}

		resp_data, err := resp.Pack()

		if err != nil {
			return err
		}

		fmt.Println(resp_data)

		resp_header := new(message.Header)

		resp_header.Build(0x20, 1, resp_data)

		header_data, err := resp_header.Pack()

		fmt.Println("header", header_data, resp_header, message.BuildHeader(0x20, 1, resp_data))

		tosend := append(message.BuildHeader(0x20, 1, resp_data), 0x1)
		tosend = append(tosend, resp_data...)

		_, err = conn.Write(tosend)

		if err != nil {
			return err
		}

		cursor += c
	default:
		fmt.Println(header.Type)
	}

	return nil
}
