package main

import (
	"errors"
	"strings"

	"github.com/syndtr/goleveldb/leveldb"
	"github.com/syndtr/goleveldb/leveldb/util"
)

var db *leveldb.DB

func init_db() {
	var err error
	db, err = leveldb.OpenFile(Config.DBFile, nil)
	if err != nil {
		log.Fatal(err)
	}
}

func parseRecord(s string) map[string]interface{} {
	rec := make(map[string]interface{})

	a := strings.Split(s, ":")
	rec["type"] = a[0]
	rec["timestamp"] = a[1]
	rec["query"] = a[2]
	rec["ip"] = a[3]
	rec["latency"] = a[4]

	return rec
}

func getRecords(start string, end string) ([]map[string]interface{}, error) {
	data := make([]map[string]interface{}, 0)

	if start > end {
		return data, errors.New("Invalid query string")
	}

	s1 := "ICMP:" + start
	s2 := "ICMP:" + end

	iter := db.NewIterator(&util.Range{Start: []byte(s1), Limit: []byte(s2)}, nil)
	for iter.Next() {
		key := iter.Key()
		data = append(data, parseRecord(string(key)))
	}
	iter.Release()
	err := iter.Error()

	return data, err
}
