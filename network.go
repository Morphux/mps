/*********************************** LICENSE **********************************\
*                            Copyright 2017 Morphux                            *
*                                                                              *
*        Licensed under the Apache License, Version 2.0 (the "License");       *
*        you may not use this file except in compliance with the License.      *
*                  You may obtain a copy of the License at                     *
*                                                                              *
*                 http://www.apache.org/licenses/LICENSE-2.0                   *
*                                                                              *
*      Unless required by applicable law or agreed to in writing, software     *
*       distributed under the License is distributed on an "AS IS" BASIS,      *
*    WITHOUT WARRANTIES OR CONDITIONS OF ANY KIND, either express or implied.  *
*        See the License for the specific language governing permissions and   *
*                       limitations under the License.                         *
\******************************************************************************/

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
