package main

import (
	"github.com/syndtr/goleveldb/leveldb"
)

var db *leveldb.DB

func init_db() {
	var err error
	db, err = leveldb.OpenFile("test.db", nil)
	if err != nil {
		log.Fatal(err)
	}
}
