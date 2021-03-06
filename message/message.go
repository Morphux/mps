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

//Message is the basic structure to send and receive Message all the methods can be overwritted if needed
type Message struct {
}

func (p *Message) Default() {
	//default value
}

func (p Message) Check(stats *barrel.Stats) bool {

	return true
}

func (p *Message) Unpack(data []byte) error {
	barrel := barrel.NewBarrel()
	load := barrel.Load(p, data, false)

	err := barrel.Unpack(load)

	return err
}

func (p *Message) Pack() ([]byte, error) {
	barrel := barrel.NewBarrel()
	load := barrel.Load(p, []byte{}, true)

	err := barrel.Pack(load)

	return barrel.Bytes(), err
}
