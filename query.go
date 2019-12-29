package main

import (
	"encoding/json"
	"fmt"
	"io/ioutil"
	"net/http"
	"reflect"
	"strconv"
)

const (
	API string = "https://api.ip.sb/geoip/"
)

type IPInfo struct {
	IP            string  `json:"ip"`
	CountryCode   string  `json:"country_code"`
	Country       string  `json:"country"`
	RegionCode    string  `json:"region_code"`
	Region        string  `json:"region"`
	City          string  `json:"city"`
	PostalCode    string  `json:"postal_code"`
	ContinentCode string  `json:"continent_code"`
	Latitude      float64 `json:"latitude"`
	Longitude     float64 `json:"longitude"`
	Organization  string  `json:"organization"`
	Timezone      string  `json:"timezone"`
}

func queryInfo(address string) {
	var info IPInfo

	ip := lookupIP(address)

	res, err := http.Get(API + ip)
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
		if names[i] == "Latitude" || names[i] == "Longitude" {
			values[i] = strconv.FormatFloat(v.Field(i).Interface().(float64), 'f', -1, 64)
		} else {
			values[i] = v.Field(i).Interface().(string)
		}
	}
	l := getMaxNameLength(names)
	for i := 0; i < v.NumField(); i++ {
		if values[i] != "" {
			fmt.Printf("%-*s:    %s\n", l, names[i], values[i])
		}
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
