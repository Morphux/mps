package main

import (
	"database/sql"
	"log"

	_ "github.com/mattn/go-sqlite3"
)

// MPM package
type Package struct {
	ID            uint64
	Name          string
	Version       string
	Category      string
	Description   string
	Archive       string
	SBU           uint64
	Dependencies  string
	ArchiveSize   uint64
	InstalledSize uint64
	ArchiveHash   string
}

func QueryPkgNameAndCat(name string, category string, state uint8, db *sql.DB) (Package, error) {
	pkg := Package{}

	var err error

	var rows *sql.Rows

	if name != "" && category != "" {
		rows, err = db.Query("SELECT * FROM pkgs where name = ? AND category = ?", name, category)
	} else if name != "" {
		rows, err = db.Query("SELECT * FROM pkgs where name = ?", name)
	} else {
		rows, err = db.Query("SELECT * FROM pkgs where category = ?", category)
	}

	if err != nil {
		log.Fatalln(err)
	}
	for rows.Next() {
		err := rows.Scan(&pkg.ID, &pkg.Name, &pkg.Version, &pkg.Category, &pkg.Description,
			&pkg.Archive, &pkg.SBU, &pkg.Dependencies, &pkg.ArchiveSize, &pkg.InstalledSize, &pkg.ArchiveHash)
		if err != nil {
			return pkg, err
			log.Fatalln(err)
		}
		return pkg, nil
	}

	return pkg, nil
}

func QueryPkgID(id uint64, state uint8, db *sql.DB) (Package, error) {
	pkg := Package{}
	rows, err := db.Query("SELECT * FROM pkgs where id = ?", id)
	if err != nil {
		log.Fatalln(err)
	}
	for rows.Next() {
		err := rows.Scan(&pkg.ID, &pkg.Name, &pkg.Version, &pkg.Category, &pkg.Description,
			&pkg.Archive, &pkg.SBU, &pkg.Dependencies, &pkg.ArchiveSize, &pkg.InstalledSize, &pkg.ArchiveHash)
		if err != nil {
			return pkg, err
			log.Fatalln(err)
		}
		return pkg, nil
	}
	return pkg, nil
}
