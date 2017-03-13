package main

import (
	"database/sql"
	"fmt"

	"github.com/Nyarum/barrel"
)

type ReqGetPKG struct {
	Message

	ID         uint64
	State      uint8
	NameLen    uint16
	CategLen   uint16
	VersionLen uint16
	Name       string
	Category   string
	Version    string
}

func RequestPackage(data []byte, db *sql.DB) (int, Package, error) {
	var err error
	pkg := Package{}

	req := new(ReqGetPKG)

	req.Unpack(data)

	fmt.Println(req)

	if req.ID != 0 && req.NameLen == 0 && req.CategLen == 0 {
		fmt.Println("search By id")
		pkg, err = QueryPkgID(req.ID, req.State, db)
	} else if req.ID == 0 && req.NameLen != 0 && req.CategLen != 0 {
		fmt.Println("search By category")
		pkg, err = QueryPkgNameAndCat(req.Name, req.Category, req.State, db)
	} else {
		fmt.Println("wtfmate")
	}

	fmt.Println(pkg)
	return 0, pkg, err
}

func (h *ReqGetPKG) Unpack(data []byte) (int, error) {

	var l uint16 = 15

	fmt.Println(data)

	barrel := barrel.NewBarrel()
	load := barrel.Load(h, data, false)

	err := barrel.Unpack(load)
	if err != nil {
		return 0, err
	}

	fmt.Println(h.ID, h.NameLen, h.CategLen, h.VersionLen)

	h.Name = string(data[l : l+h.NameLen])
	h.Category = string(data[l+h.NameLen : l+h.NameLen+h.CategLen])
	h.Version = string(data[l+h.NameLen+h.CategLen : l+h.NameLen+h.CategLen+h.VersionLen])

	return int(l + h.NameLen + h.CategLen + h.VersionLen), nil
}
