// Copyright 2018 panigrahi.kiran@gmail.com. All rights reserved.
// Use of this source code is governed by a MIT
// license that can be found in the LICENSE file.

package process

import (
	"io/ioutil"
	"log"
	"os"
	"path/filepath"

	"../utils"
)

const (
	folderConst  = "folder"
	archiveConst = "archive"
	fileConst    = "file"
)

// GetAllFileList - function to get all the file list
func GetAllFileList(source string, excludeFileList []string) []string {
	files, err := ioutil.ReadDir(source)
	if err != nil {
		log.Fatal(err)
	}
	var fileList []string
	for _, file := range files {
		excludeFile := false
		for _, exFl := range excludeFileList {
			if file.Name() == exFl {
				excludeFile = true
				break
			}
		}
		if excludeFile == false {
			fileName := filepath.Join(source, file.Name())
			fileType := utils.GetFileType(fileName)
			if fileType == folderConst {
				for _, fl := range GetAllFileList(fileName, excludeFileList) {
					fileList = append(fileList, fl)
				}
			}
			if fileType == archiveConst {
				err = utils.UnArchive(fileName, filepath.Dir(fileName))
				if err != nil {
					log.Fatal(err)
				}
				os.RemoveAll(fileName)
				for _, fl := range GetAllFileList(filepath.Dir(fileName), excludeFileList) {
					fileList = append(fileList, fl)
				}
			}
			if fileType == fileConst {
				fileList = append(fileList, fileName)
			}
		}
	}
	return fileList
}
