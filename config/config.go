// Copyright 2018 panigrahi.kiran@gmail.com. All rights reserved.
// Use of this source code is governed by a MIT
// license that can be found in the LICENSE file.

package config

import (
	"encoding/json"
	"log"
	"os"
	"time"
)

type folder struct {
	In    string `json:"input"`
	Out   string `json:"output"`
	Archv string `json:"archive"`
	Wrkg  string `json:"working"`
}

// Config - structure to hold the entire configuration
type Config struct {
	Folders      folder   `json:"folders"`
	ChunkSize    int      `json:"chunkSize"`
	ExcludeFiles []string `json:"excludeFiles"`

	UUID        string // not part of the config.json
	RunDttmStmp string // not part of the config.json
}

// Load - function to load the configuration from the file and gets a
// Config object
func Load(configFileName string) Config {
	var config Config

	configFile, err := os.Open(configFileName)
	if err != nil {
		log.Fatal(err)
	}
	defer configFile.Close()

	err = json.NewDecoder(configFile).Decode(&config)
	if err != nil {
		log.Fatal(err)
	}
	// get the current time stamp
	sCurDttmStmp := time.Now().Format("20060102150405")
	// set the UUID
	config.UUID = sCurDttmStmp
	// for different format please refer to the below URL
	// https://medium.com/@Martynas/formatting-date-and-time-in-golang-5816112bf098
	// set the run date time
	config.RunDttmStmp = time.Now().Format("2006-Jan-02 15:04:05.000")

	return config
}
