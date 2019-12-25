package main

import (
	"encoding/json"
	"fmt"
	"io/ioutil"
	"net/http"
	"reflect"
)

const (
	API string = "http://ip-api.com/json/"
)

type IPInfo struct {
	Query   string
	IP      string `json:"query"`
	City    string `json:"city"`
	Country string `json:"country"`
	ISP     string `json:"isp"`
	AS      string `json:"as"`
}

func queryInfo(q string) {
	var info IPInfo

	query := parseInput(q)
	info.Query = query

	res, err := http.Get(API + query)
	if err != nil {
		log.Fatal(err)
	}
	defer res.Body.Close()

	body, err := ioutil.ReadAll(res.Body)
	if err != nil {
		log.Fatal(err)
	}
	err = json.Unmarshal(body, &info)
	if err != nil {
		fmt.Println(err)
	}
	printInfo(info)
}

func printInfo(info IPInfo) {
	v := reflect.ValueOf(info)
	names := make([]string, v.NumField())
	values := make([]string, v.NumField())
	for i := 0; i < v.NumField(); i++ {
		names[i] = v.Type().Field(i).Name
		values[i] = v.Field(i).Interface().(string)
	}
	l := getMaxNameLength(names)
	for i := 0; i < v.NumField(); i++ {
		fmt.Printf("%-*s:    %s\n", l, names[i], values[i])
	}
}

func getMaxNameLength(names []string) int {
	var length int
	for _, val := range names {
		if len(val) > length {
			length = len(val)
		}
	}
	return length
}
