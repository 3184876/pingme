package main

import (
	"encoding/json"
	"fmt"
	"io/ioutil"
)

type ConfigData struct {
	DBFile    string `json:"db_file"`
	ServePort int    `json:"serve_port"`
}

var Config ConfigData = ConfigData{
	DBFile:    "test.db",
	ServePort: 5000,
}

func init_config() {
	jsonData, err := ioutil.ReadFile(ConfigFile)
	err = json.Unmarshal(jsonData, &Config)
	if err != nil {
		fmt.Println(err)
	}
}
