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
	"crypto/rand"
	"crypto/tls"
	"database/sql"
	"flag"
	"fmt"
	"log"
	"net"
	"os"

	_ "github.com/mattn/go-sqlite3"
)

// const MajorVersion uint8 = 0
// const MinorVersion uint8 = 1

func main() {

	databasePtr := flag.String("db", "", "A sqlite database")

	tlsPtr := flag.Bool("notls", false, "Do not use tls")
	PublicKeyPtr := flag.String("pub", "", "Public key path")
	PrivateKeyPtr := flag.String("priv", "", "Private key path")

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

	var listener net.Listener

	if *tlsPtr == false && *PrivateKeyPtr != "" && *PublicKeyPtr != "" {
		cert, err := tls.LoadX509KeyPair(*PrivateKeyPtr, *PublicKeyPtr)
		if err != nil {
			log.Fatalf("server: loadkeys: %s", err)
		}
		config := tls.Config{Certificates: []tls.Certificate{cert}}
		config.Rand = rand.Reader

		fmt.Println("Successfully load key pair ", *PrivateKeyPtr, *PublicKeyPtr)

		listener, err = tls.Listen("tcp", flag.Args()[0], &config)
	} else {
		if (*tlsPtr == true) {
			listener, err = net.Listen("tcp", flag.Args()[0])
		} else {
			fmt.Println("Cannot launch server without tls support")
			os.Exit(1)
		}
	}

	if listener == nil {
		fmt.Println("Error listening: listener == nil (Maybe unbindable address)")
		os.Exit(1)
	}

	if err != nil {
		fmt.Println("Error listening:", err.Error())
		os.Exit(1)
	}

	defer listener.Close()
	if *tlsPtr == false {
		fmt.Println("Listening on "+flag.Args()[0], "with TLS support")
	} else {
		fmt.Println("Listening on " + flag.Args()[0])
	}

	for {
		conn, err := listener.Accept()
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
