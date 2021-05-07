package main

import (
	"encoding/json"
	"fmt"
	"io/ioutil"
	"log"
	"os"
	"time"
)

func weitejson() {
	fileName := "plugins\\BDSUpdate\\BDSUpdate.json"
	dstFile, err := os.Create(fileName)
	if err != nil {
		fmt.Println(err.Error())
		return
	}
	defer dstFile.Close()
	s := "{\n\"elua\":false,\n\"privacypolicy\":false,\n\"windows\":true,\n\"linux\":true,\n\"times\":20,\n\"downloadDestFolder\":\"plugins\\\\BDSUpdate\\\\\"\n}"
	dstFile.WriteString(s)
	dstFile.Close()
}

type Configs struct {
	Eula               bool          `json:"eula"`
	PrivacyPolicy      bool          `json:"privacypolicy"`
	Windows            bool          `json:"windows"`
	Linux              bool          `json:"linux"`
	Times              time.Duration `json:"times"`
	DownloadDestFolder string        `json:"downloadDestFolder"`
}

var LLConfig Configs

func ReadConfig() Configs {
	f, ferr := os.OpenFile("plugins\\BDSUpdate\\BDSUpdate.json", os.O_RDONLY, 0600)
	defer f.Close()
	if ferr != nil {
		log.Printf("[Error] Error when reading config > %s", ferr.Error())
	}
	config_bytes, _ := ioutil.ReadAll(f)
	var con Configs
	jerr := json.Unmarshal(config_bytes, &con)
	if jerr != nil {
		log.Printf("[Error] Something wrong with config_json > %s", jerr.Error())
	}
	return con
}
