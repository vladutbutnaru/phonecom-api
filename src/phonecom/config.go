package main

import (
	"os"
	"fmt"
	"io/ioutil"
	"encoding/xml"
)

type Query struct {
	ConfigList []Config `xml:"Config"`
}

type Config struct {
	BaseApiPath string
	ApiKeyPrefix string
	ApiKey string
	Type string
}

func getConfig() Config {

	xmlFile, err := os.Open("config.xml")

	var noConfig Config;

	if err != nil {
		fmt.Println("Could not read config.xml", err)
		return noConfig
	}

	defer xmlFile.Close()

	content, _ := ioutil.ReadAll(xmlFile)

	var q Query
	xml.Unmarshal(content, &q)

	for _, config := range q.ConfigList {
		if (config.Type == "main") {
			return config
		}
	}

	return noConfig
}
