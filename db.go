package main

import (
	"github.com/syndtr/goleveldb/leveldb"
)

var db *leveldb.DB

func init_db() {
	var err error
	db, err = leveldb.OpenFile(Config.DBFile, nil)
	if err != nil {
		log.Fatal(err)
	}
}

func getAllData() []string {
	var data []string

	iter := db.NewIterator(nil, nil)
	for iter.Next() {
		key := iter.Key()
		data = append(data, string(key))
	}

	return data
}
