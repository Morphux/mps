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
	"errors"
	"fmt"
	"log"
	"strconv"
	"strings"

	"github.com/Morphux/mps/request"
	"github.com/Morphux/mps/response"
	_ "github.com/mattn/go-sqlite3"
)

//Package is a MPS internally used interface to make a bridge between the sql query and the resp_pkg
type Package struct {
	ID            uint64
	Name          string
	State         uint8
	Version       string
	Category      string
	Description   string
	Archive       string
	SBU           float64
	Dependencies  string
	ArchiveSize   float64
	InstalledSize float64
	ArchiveHash   string
	TimeAddPkg    uint64
}

//RequestPackage return a `Package` struct corresponding to the packet passed as data
func RequestPackage(data []byte, db *sql.DB) (int, Package, error) {
	pkg := Package{}

	req := new(request.ReqGetPKG)

	n, err := req.Unpack(data)

	if err != nil {
		return n, pkg, err
	}

	fmt.Println("unpacked :", req)

	if req.ID != 0 && req.NameLen == 0 && req.CategLen == 0 {
		fmt.Println("search By id")
		pkg, err = QueryPkgID(req.ID, req.State, db)
	} else if req.ID == 0 && req.NameLen != 0 && req.CategLen != 0 {
		fmt.Println("search By category")
		pkg, err = QueryPkgNameAndCat(req.Name, req.Category, req.State, db)
	} else {
		err = errors.New("A packet send by the client is wrong")
	}

	return n, pkg, err
}

//PkgtoRespPkg convert a `Package` to the correct reps_pkg packet
func PkgtoRespPkg(pkg Package) (*response.RespPkg, error) {

	dep := strings.Split(pkg.Dependencies, ",")

	var depID []uint64
	for _, v := range dep {
		i, err := strconv.Atoi(v)
		if err != nil {
			continue
		}
		depID = append(depID, uint64(i))
	}

	ret := new(response.RespPkg)
	ret.ID = pkg.ID
	ret.CompTime = float32(pkg.SBU)
	ret.InstSize = float32(pkg.InstalledSize)
	ret.ArchSize = float32(pkg.ArchiveSize)
	ret.State = pkg.State
	ret.NameLen = uint16(len(pkg.Name))
	ret.CategoryLen = uint16(len(pkg.Category))
	ret.VersionLen = uint16(len(pkg.Version))
	ret.ArchiveLen = uint16(len(pkg.Archive))
	ret.ChecksumLen = uint16(len(pkg.ArchiveHash))
	ret.DependenciesSize = uint16(len(depID))
	ret.Name = []byte(pkg.Name)
	ret.Category = []byte(pkg.Category)
	ret.Version = []byte(pkg.Version)
	ret.Archive = []byte(pkg.Archive)
	ret.Checksum = []byte(pkg.ArchiveHash)
	ret.Dependencies = depID

	return ret, nil
}

//QueryPkgNameAndCat Query the database to get a package by its name or categories
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
		err := rows.Scan(&pkg.ID, &pkg.Name, &pkg.State, &pkg.Version, &pkg.Category, &pkg.Description,
			&pkg.Dependencies, &pkg.Archive, &pkg.SBU, &pkg.ArchiveSize, &pkg.InstalledSize, &pkg.ArchiveHash, &pkg.TimeAddPkg)
		if err != nil {
			log.Fatalln(err)
			return pkg, err
		}
		return pkg, nil
	}

	return pkg, nil
}

//QueryPkgID Query the database to get a package by its id
func QueryPkgID(id uint64, state uint8, db *sql.DB) (Package, error) {
	pkg := Package{}
	rows, err := db.Query("SELECT * FROM pkgs where id = ?", id)
	if err != nil {
		log.Fatalln(err)
	}
	for rows.Next() {
		err := rows.Scan(&pkg.ID, &pkg.Name, &pkg.State, &pkg.Version, &pkg.Category, &pkg.Description,
			&pkg.Dependencies, &pkg.Archive, &pkg.SBU, &pkg.ArchiveSize, &pkg.InstalledSize, &pkg.ArchiveHash, &pkg.TimeAddPkg)
		if err != nil {
			log.Fatalln(err)
			return pkg, err
		}
		return pkg, nil
	}
	return pkg, nil
}
