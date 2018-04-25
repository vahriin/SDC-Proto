package model

import (
	"encoding/json"
	"fmt"
	"os"
)

type ConnectionConf struct {
	DbName   string `json:"dbname"`
	User     string `json:"user"`
	Password string `json:"password"`
	Host     string `json:"host"`
	Port     string `json:"port"`
}

func DefaultConnectionConf() ConnectionConf {
	var cc ConnectionConf
	cc.DbName = "sqltest"
	cc.User = "sqltest_user"
	cc.Password = "sqltest_password"
	cc.Host = "localhost"
	cc.Port = "5432"

	return cc
}

func ReadConnectionConf(filename string) ConnectionConf {
	confFile, err := os.Open(filename)
	if err != nil {
		fmt.Println(err.Error())
		return DefaultConnectionConf()
	}

	jsonDecoder := json.NewDecoder(confFile)
	var cc ConnectionConf
	jsonDecoder.Decode(&cc)
	return cc
}
