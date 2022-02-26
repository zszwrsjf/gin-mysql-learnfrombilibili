package utils

import (
	"fmt"

	"gopkg.in/ini.v1"
)

var (
	AppMode  string
	HttpPort string

	Db         string
	DbHost     string
	DbPort     string
	DbUser     string
	DbPassWord string
	DbName     string
)

func init() {
	file, err := ini.Load("config/config.ini")
	if err != nil {
		fmt.Println("error config: ", err)
	}
	LoadServer(file)
	LoadData(file)
}

func LoadServer(file *ini.File) {
	AppMode = file.Section("server").Key("AppMode").MustString("debug")
	HttpPort = file.Section("server").Key("HttpPort").MustString("8080")
}

// TODO muststring setting from config.ini
func LoadData(file *ini.File) {
	Db = file.Section("database").Key("Db").MustString("debug")
	DbHost = file.Section("database").Key("DbHost").MustString("")
	DbPort = file.Section("database").Key("DbPort").MustString("")
	DbUser = file.Section("database").Key("DbUser").MustString("")
	DbPassWord = file.Section("database").Key("DbPassWord").MustString("")
	DbName = file.Section("database").Key("DbName").MustString("")
}
