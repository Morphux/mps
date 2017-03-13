package main

type ReqGetCat struct {
	Message

	CategoriesLen uint16
	Categories    []uint64
}
