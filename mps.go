package main

import (
	"database/sql"
	"flag"
	"fmt"
	"io/ioutil"
	"log"
	"net"
	"os"

	_ "github.com/mattn/go-sqlite3"
)

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

	message, err := ioutil.ReadAll(conn)
	if err != nil {
		log.Print("Error reading:", err.Error())
	}
	ParseRequest(message, conn)
	conn.Close()
}
