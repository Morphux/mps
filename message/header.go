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

package message

import "github.com/Morphux/mps/vendors/Nyarum/barrel"

type Header struct {
	Message

	Type       uint8
	Size       uint16
	NextPkgLen uint8
	NextPkg    string
}

//Unpack []byte to a proper header object, return the number of byte used and an error
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

//Build help to build a correct header
func (h *Header) Build(Type, number uint8, data []byte) {
	h.Type = Type
	h.NextPkgLen = 0
	h.NextPkg = ""

	h.Size = uint16(int(1+2+1+h.NextPkgLen+number) + len(data))

}

//Pack an header to []byte
func (p *Header) Pack() ([]byte, error) {
	barrel := barrel.NewBarrel()
	load := barrel.Load(p, []byte{}, true)

	err := barrel.Pack(load)

	return barrel.Bytes(), err
}

//BuildHeader Is theorically faster than generating header and Packing it
func BuildHeader(Type, number uint8, data []byte) []byte {

	hash := ""
	length := len(data) + 4 + 1

	header := []byte{Type, byte(length & 0xff), byte(length >> 8), byte(len(hash))}

	return header

}
