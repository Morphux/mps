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

//An error happened server side
const ERR_SERVER_FAULT uint8 = 0x1

//A packet send by the client is wrong
const ERR_MALFORMED_PACKET uint8 = 0x2

//A request send by the client find no result
const ERR_RES_NOT_FOUND uint8 = 0x3

type Error struct {
	Message

	ErrorType uint8
	ErrorLen  uint16
	Error     string
}

//PackError take a error and transform it to an mps.Error and pack it to []byte
func (p *Error) PackError(err error, errortype uint8) ([]byte, error) {
	p.ErrorType = errortype
	p.Error = err.Error()
	p.ErrorLen = uint16(len(p.Error))

	return p.Pack()
}

//ServerFault Return Server Fault Message
func ServerFault() []byte {
	errorMessage := []byte("An error happened server side") //len 0x1D
	return append([]byte{0x03, 0x25, 0x00, 0x00, 0x01, 0x01, 0x1D, 0x00}, errorMessage...)
}

//MalformedPacket Return Malformed Malformed Message
func MalformedPacket() []byte {
	errorMessage := []byte("A packet send by the client is wrong") //len 0x24
	return append([]byte{0x03, 0x2C, 0x00, 0x00, 0x01, 0x02, 0x24, 0x00}, errorMessage...)
}

//ResNotFound Return Ressource Not Found Message
func ResNotFound() []byte {
	errorMessage := []byte("A request send by the client find no result") //len 0x2B
	return append([]byte{0x03, 0x33, 0x00, 0x00, 0x01, 0x03, 0x2B, 0x00}, errorMessage...)
}

//GenerateError Generate a Error Message
func GenerateError(Type uint8, message string) []byte {
	return append([]byte{Type, byte(len(message) + 8), 0x00, 0x00, 0x01, 0x03, byte(len(message)), 0x00}, []byte(message)...)
}
