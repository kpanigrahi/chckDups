// Copyright 2018 panigrahi.kiran@gmail.com. All rights reserved.
// Use of this source code is governed by a MIT
// license that can be found in the LICENSE file.

package utils

import (
	"bufio"
	"io"
	"io/ioutil"
	"log"
	"os"
	"path/filepath"
	"strings"
)

const (
	folderConst  = "folder"
	archiveConst = "archive"
	fileConst    = "file"
)

// GetFileList - returns the file(s) and folder(s) under a specified folder
func GetFileList(folder string, excludeFileList []string) []string {
	files, err := ioutil.ReadDir(folder)
	if err != nil {
		log.Fatal(err)
	}
	var fileList []string
	for _, file := range files {
		excludeFile := false
		for _, exFile := range excludeFileList {
			if file.Name() == exFile {
				excludeFile = true
			}
		}
		if excludeFile == false {
			fileList = append(fileList, filepath.Join(folder, file.Name()))
		}
	}
	return fileList
}

// MkdirAll - creates a folder
func MkdirAll(folderName string) {
	err := os.MkdirAll(folderName, os.ModePerm)
	if err != nil {
		log.Fatal(err)
	}
}

// GetFullQualifiedPath - returns the fully qualified path of the
// specified folder name
func GetFullQualifiedPath(qualifiedPath, folder string) string {
	if folder[0:1] == "." {
		return filepath.Join(qualifiedPath, folder)
	}
	return filepath.Join(folder)
}

// CopyFile - copies file from source location to a target location
func CopyFile(src, tgt string) error {
	if GetFileType(src) == folderConst {
		os.MkdirAll(tgt, os.ModePerm)
		var excludeFileList []string
		excludeFileList = append(excludeFileList, "")
		fileList := GetFileList(src, excludeFileList)
		for _, file := range fileList {
			err := CopyFile(file, filepath.Join(tgt, filepath.Base(file)))
			if err != nil {
				return err
			}
		}
	} else {
		srcFile, _ := os.Open(src)
		defer srcFile.Close()

		tgtFile, err := os.Create(tgt)
		if err != nil {
			return err
		}
		defer tgtFile.Close()

		_, err = io.Copy(tgtFile, srcFile)
		if err != nil {
			return err
		}
		tgtFile.Close()
	}
	return nil
}

// GetFileType - gets the file type
func GetFileType(fileName string) string {
	fileInfo, _ := os.Stat(fileName)
	switch mode := fileInfo.Mode(); {
	case mode.IsDir():
		return folderConst
	case mode.IsRegular():
		if strings.ToUpper(filepath.Ext(fileName)) == ".ZIP" {
			return archiveConst
		}
	}
	return fileConst
}

// FileNameOnly function to return just the file name
func FileNameOnly(fileName string, withExt bool) string {
	if withExt {
		return filepath.Base(fileName)
	}
	return filepath.Base(fileName)[0 : len(filepath.Base(fileName))-len(filepath.Ext(fileName))]
}

// GetTotalRows function to return total number of lines in a given file
func GetTotalRows(fileName string) int {
	file, _ := os.Open(fileName)
	defer file.Close()
	// Start reading from the file with a reader.
	inputReader := bufio.NewReader(file)
	lineNbr := 0
	for {
		_, error := inputReader.ReadString('\n')
		if error != nil {
			break
		}
		lineNbr++
	}
	return lineNbr
}
