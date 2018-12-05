// Copyright 2018 panigrahi.kiran@gmail.com. All rights reserved.
// Use of this source code is governed by a MIT
// license that can be found in the LICENSE file.

package process

import (
	"bufio"
	"crypto/sha1"
	"encoding/hex"
	"fmt"
	"log"
	"os"
	"path/filepath"
	"strconv"

	"../utils"
	"github.com/schollz/progressbar"
)

// Process - core function to process duplicate rows if any in the file.
func Process(fileName string) {
	origFileName := filepath.Join(filepath.Dir(fileName), "ORIG_"+filepath.Base(fileName))
	dupsFileName := filepath.Join(filepath.Dir(fileName), "DUPS_"+filepath.Base(fileName))

	os.Rename(fileName, origFileName)

	totalLines := utils.GetTotalRows(origFileName)

	// create all the file pointers
	origFile, _ := os.Open(origFileName)
	file, _ := os.Create(fileName)
	dupsFile, _ := os.Create(dupsFileName)

	defer origFile.Close()
	defer file.Close()
	defer dupsFile.Close()

	// Start reading from the file with a reader.
	inputReader := bufio.NewReader(origFile)
	m := make(map[string]int)
	lineNbr := 0

	line, error := inputReader.ReadString('\n')
	if error != nil {
		return
	}
	file.WriteString(line)
	dupsFile.WriteString("DupsLineNbr|OrigLineNbr|" + line)
	// dupsFound := false
	lineNbr++
	dupRows := 0

	bar := progressbar.New(totalLines)
	for {
		line, error = inputReader.ReadString('\n')
		if error != nil {
			break
		}
		lineNbr++
		// convert the line into a HASH code
		h := sha1.New()
		h.Write([]byte(line))
		sHASHCode := hex.EncodeToString(h.Sum(nil))
		if m[sHASHCode] != 0 {
			// dups
			dupsFile.WriteString(strconv.Itoa(lineNbr) + "|" + strconv.Itoa(m[sHASHCode]) + "|" + line)
			// dupsFound = true
			dupRows++
		} else {
			m[sHASHCode] = lineNbr
			file.WriteString(line)
		}
		bar.Add(1)
	}
	bar.Add(1)
	fmt.Println()

	origFile.Close()
	file.Close()
	dupsFile.Close()
	if dupRows > 0 {
		os.MkdirAll(filepath.Join(filepath.Dir(origFileName), "orginals"), os.ModePerm)
		os.MkdirAll(filepath.Join(filepath.Dir(dupsFileName), "duplicates"), os.ModePerm)
		os.Rename(origFileName, filepath.Join(filepath.Dir(origFileName), "orginals", filepath.Base(origFileName)[5:len(filepath.Base(origFileName))]))
		os.Rename(dupsFileName, filepath.Join(filepath.Dir(dupsFileName), "duplicates", filepath.Base(dupsFileName)[5:len(filepath.Base(dupsFileName))]))
	} else {
		os.RemoveAll(dupsFileName)
		os.RemoveAll(fileName)
		os.Rename(origFileName, fileName)
	}
	log.Println("Duplicate rows              :", dupRows)
	log.Println("Total rows                  :", lineNbr)
	fmt.Println("----------------------------------------------------------------------------------------------")
}
