// Copyright 2018 panigrahi.kiran@gmail.com. All rights reserved.
// Use of this source code is governed by a MIT
// license that can be found in the LICENSE file.

package process

import (
	"log"
	"os"
	"path/filepath"

	"../config"
	"../utils"
)

// PrepareEnv - prepares the environment before processing
func PrepareEnv(config config.Config) {
	// create all the required folders
	utils.MkdirAll(config.Folders.Out)
	// utils.MkdirAll(filepath.Join(config.Folders.Archv, config.UUID))
	utils.MkdirAll(filepath.Join(config.Folders.Wrkg, config.UUID))

	fileList := utils.GetFileList(config.Folders.In, config.ExcludeFiles)
	// // copy files from input to archive folder
	// for _, file := range fileList {
	// 	err := utils.CopyFile(file, filepath.Join(config.Folders.Archv, config.UUID, filepath.Base(file)))
	// 	if err != nil {
	// 		log.Fatal(err)
	// 	}
	// }
	// utils.Archive(filepath.Join(config.Folders.Archv, config.UUID), filepath.Join(config.Folders.Archv, config.UUID+".zip"))

	// // now that archive is complete, we can delete the UUID folder
	// os.RemoveAll(filepath.Join(config.Folders.Archv, config.UUID))

	// move the files from input to working folder
	for _, file := range fileList {
		err := os.Rename(file, filepath.Join(config.Folders.Wrkg, config.UUID, filepath.Base(file)))
		if err != nil {
			log.Fatal(err)
		}
		// if utils.GetFileType(filepath.Join(config.Folders.Wrkg, config.UUID, filepath.Base(file))) == archiveConst {
		// 	fileName := filepath.Join(config.Folders.Wrkg, config.UUID, filepath.Base(file))
		// 	folderName := "zip_" + utils.FileNameOnly(fileName, false)
		// 	// create a folder
		// 	os.MkdirAll(filepath.Join(filepath.Dir(fileName), folderName), os.ModePerm)
		// 	// move the file into filename folder
		// 	os.Rename(fileName, filepath.Join(filepath.Dir(fileName), folderName, filepath.Base(fileName)))
		// }
	}
}
