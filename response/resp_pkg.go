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

package response

import (
	"fmt"

	"github.com/Morphux/mps/message"
	"github.com/Morphux/mps/vendors/Nyarum/barrel"
)

type RespPkg struct {
	message.Message

	ID               uint64
	CompTime         float32
	InstSize         float32
	ArchSize         float32
	State            uint8
	NameLen          uint16
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

func (p *RespPkg) Pack() ([]byte, error) {

	fmt.Printf("%#v\n", p)

	barrel := barrel.NewBarrel()
	load := barrel.Load(p, []byte{}, true)

	err := barrel.Pack(load)

	return barrel.Bytes(), err
}
