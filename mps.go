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
	"flag"
	"fmt"
	"log"
	"net"
	"os"

	_ "github.com/mattn/go-sqlite3"
)

const MPM_MAJOR_VERSION uint8 = 0
const MPM_MINOR_VERSION uint8 = 1

func main() {

	databasePtr := flag.String("db", "", "a sqlite database")

	flag.Parse()

	if *databasePtr == "" || len(flag.Args()) == 0 {
		flag.Usage()
		os.Exit(0)
	}

	db, err := sql.Open("sqlite3", *databasePtr)
	if err != nil {
		log.Fatal(err)
	}
	defer db.Close()

	l, err := net.Listen("tcp", flag.Args()[0])
	if err != nil {
		fmt.Println("Error listening:", err.Error())
		os.Exit(1)
	}

	defer l.Close()
	fmt.Println("Listening on " + flag.Args()[0])
	for {

		conn, err := l.Accept()
		if err != nil {
			fmt.Println("Error accepting: ", err.Error())
			os.Exit(1)
		}

		go handleRequest(conn, db)
	}
}

func handleRequest(conn net.Conn, db *sql.DB) {
	buf := make([]byte, 4096)

	//var auth bool = false

	for {
		n, err := conn.Read(buf)
		//fmt.Printf("receive :%#v\n\n", buf[0:n])

		errParse := ParseRequest(buf[0:n], conn, db)

		if errParse != nil {
			log.Print(errParse)
		}

		if err != nil || n == 0 {
			conn.Close()
			break
		}

		if err != nil {
			conn.Close()
			break
		}
	}

}
