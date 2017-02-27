package main

import (
	"database/sql"
	"fmt"
	"log"

	_ "github.com/mattn/go-sqlite3"
)

// MPM package
type Package struct {
	id             uint64
	name           string
	version        string
	category       string
	description    string
	archive        string
	sbu            uint64
	dependencies   string
	archive_size   uint64
	installed_size uint64
	archive_hash   string
}

func queryPackageByName(name string, state int, db *sql.DB) {
	pkg := Package{}
	rows, err := db.Query("SELECT * FROM pkgs where name = ?", name)
	if err != nil {
		log.Fatalln(err)
	}
	for rows.Next() {
		err := rows.Scan(&pkg.id, &pkg.name, &pkg.version, &pkg.category, &pkg.description,
			&pkg.archive, &pkg.sbu, &pkg.dependencies, &pkg.archive_size, &pkg.installed_size, &pkg.archive_hash)
		if err != nil {
			log.Fatalln(err)
		}
		fmt.Printf("%#v\n", pkg)
	}
}

func queryPackageByid(id uint64, state int, db *sql.DB) {
	pkg := Package{}
	rows, err := db.Query("SELECT * FROM pkgs where id = ?", id)
	if err != nil {
		log.Fatalln(err)
	}
	for rows.Next() {
		err := rows.Scan(&pkg.id, &pkg.name, &pkg.version, &pkg.category, &pkg.description,
			&pkg.archive, &pkg.sbu, &pkg.dependencies, &pkg.archive_size, &pkg.installed_size, &pkg.archive_hash)
		if err != nil {
			log.Fatalln(err)
		}
		fmt.Printf("%#v\n", pkg)
	}
}
