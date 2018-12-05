// Copyright 2018 panigrahi.kiran@gmail.com. All rights reserved.
// Use of this source code is governed by a MIT
// license that can be found in the LICENSE file.

package main

import (
	"fmt"
	"log"
	"os"
	"path/filepath"

	"./config"
	"./process"
	"./utils"
)

const (
	// http://patorjk.com/software/taag/#p=display&h=2&f=Banner4&t=Check%20dups!
	banner = `
 ######  ##     ## ########  ######  ##    ##   ########  ##     ## ########   ######   ####
##    ## ##     ## ##       ##    ## ##   ##    ##     ## ##     ## ##     ## ##    ##  ####
##       ##     ## ##       ##       ##  ##     ##     ## ##     ## ##     ## ##        ####
##       ######### ######   ##       #####      ##     ## ##     ## ########   ######    ##
##       ##     ## ##       ##       ##  ##     ##     ## ##     ## ##              ##
##    ## ##     ## ##       ##    ## ##   ##    ##     ## ##     ## ##        ##    ##  ####
 ######  ##     ## ########  ######  ##    ##   ########   #######  ##         ######   ####
 `
	line      = "----------------------------------------------------------------------------------------------"
	version   = "1.0."
	copyright = "Copyright 2018 panigrahi.kiran@gmail.com. All rights reserved."
)

func main() {
	fmt.Println(banner)
	fmt.Println("Version:", version, copyright)
	fmt.Println(copyright)
	fmt.Println(line)
	fmt.Println()

	currWrkgDir, _ := os.Getwd()
	configFileLoc := "./config/config.json"

	config := config.Load(configFileLoc)
	// set the fully qualified path if user has provided relative path
	config.Folders.In = utils.GetFullQualifiedPath(currWrkgDir, config.Folders.In)
	config.Folders.Out = utils.GetFullQualifiedPath(currWrkgDir, config.Folders.Out)
	config.Folders.Archv = utils.GetFullQualifiedPath(currWrkgDir, config.Folders.Archv)
	config.Folders.Wrkg = utils.GetFullQualifiedPath(currWrkgDir, config.Folders.Wrkg)

	fileList := utils.GetFileList(config.Folders.In, config.ExcludeFiles)
	if len(fileList) == 0 {
		// no file(s) for processing
		log.Fatal(fmt.Sprintf("no file(s) %s: The system cannot find any file(s) for processing.", config.Folders.In))
	} else {
		process.PrepareEnv(config)
		// by this time we already moved the files from input to working folder
		// expand and then get all the file list using breadth search algorithm
		fileList = process.GetAllFileList(filepath.Join(config.Folders.Wrkg, config.UUID), config.ExcludeFiles)
		log.Println("Total file(s)               :", len(fileList))
		for indx, file := range fileList {
			log.Printf("currently processing (%2d/%2d): %s\n", indx+1, len(fileList), filepath.Base(file))
			process.Process(file)
		}
		process.PostProcess(config)
	}
}
