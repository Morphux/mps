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
