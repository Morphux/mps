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

package request

import (
	"errors"
	"fmt"

	"github.com/Morphux/mps/message"
	"github.com/Morphux/mps/vendors/Nyarum/barrel"
)

type ReqGetPKG struct {
	message.Message

	ID         uint64
	State      uint8
	NameLen    uint16
	CategLen   uint16
	VersionLen uint16
	Name       string
	Category   string
	Version    string
}

func (h *ReqGetPKG) Unpack(data []byte) (int, error) {

	fmt.Printf("=====\ndata %#v\n\n", data)

	var l uint16 = 15

	fmt.Println(data)

	barrel := barrel.NewBarrel()
	load := barrel.Load(h, data, false)

	err := barrel.Unpack(load)
	if err != nil {
		return 0, err
	}

	//fmt.Println(h.ID, h.NameLen, h.CategLen, h.VersionLen)

	if len(data) < int(l+h.NameLen+h.CategLen+h.VersionLen) {
		return 0, errors.New("A packet send by the client is wrong")
	}

	h.Name = string(data[l : l+h.NameLen])
	h.Category = string(data[l+h.NameLen : l+h.NameLen+h.CategLen])
	h.Version = string(data[l+h.NameLen+h.CategLen : l+h.NameLen+h.CategLen+h.VersionLen])

	return int(l + h.NameLen + h.CategLen + h.VersionLen), nil
}
