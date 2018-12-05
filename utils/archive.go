// Copyright 2018 panigrahi.kiran@gmail.com. All rights reserved.
// Use of this source code is governed by a MIT
// license that can be found in the LICENSE file.

package utils

import (
	"archive/zip"
	"fmt"
	"io"
	"io/ioutil"
	"log"
	"os"
	"path/filepath"
	"strings"
)

// Archive - archives file(s) and folder(s)
func Archive(source, archvFileName string) {
	log.Println("archiving is in progress...")
	// Get a Buffer to Write To
	outFile, err := os.Create(archvFileName)
	if err != nil {
		log.Fatal(err)
	}
	defer outFile.Close()
	// Create a new zip archive.
	w := zip.NewWriter(outFile)
	// Add some files to the archive.
	addFiles(w, source, "")
	if err != nil {
		log.Fatal(err)
	}
	// Make sure to check the error on Close.
	err = w.Close()
	if err != nil {
		log.Fatal(err)
	}
	log.Println("archiving completed successfully...")
}

func addFiles(w *zip.Writer, basePath, baseInZip string) {
	// Open the Directory
	files, err := ioutil.ReadDir(basePath)
	if err != nil {
		log.Fatal(err)
	}

	for _, file := range files {
		if !file.IsDir() {
			dat, err := ioutil.ReadFile(filepath.Join(basePath, file.Name()))
			if err != nil {
				log.Fatal(err)
			}
			// Add some files to the archive.
			f, err := w.Create(filepath.Join(baseInZip, file.Name()))
			if err != nil {
				log.Fatal(err)
			}
			_, err = f.Write(dat)
			if err != nil {
				log.Fatal(err)
			}
		} else if file.IsDir() {
			newBase := filepath.Join(basePath, file.Name())
			addFiles(w, newBase, file.Name()+"/")
		}
	}
}

// UnArchive will decompress a zip archive, moving all files and folders
// within the zip file (parameter 1) to an output directory (parameter 2).
func UnArchive(src string, dest string) error {

	var filenames []string

	r, err := zip.OpenReader(src)
	if err != nil {
		return err
	}
	defer r.Close()

	for _, f := range r.File {

		rc, err := f.Open()
		if err != nil {
			return err
		}
		defer rc.Close()

		// Store filename/path for returning and using later on
		fpath := filepath.Join(dest, f.Name)

		// Check for ZipSlip. More Info: http://bit.ly/2MsjAWE
		if !strings.HasPrefix(fpath, filepath.Clean(dest)+string(os.PathSeparator)) {
			return fmt.Errorf("%s: illegal file path", fpath)
		}

		filenames = append(filenames, fpath)

		if f.FileInfo().IsDir() {

			// Make Folder
			os.MkdirAll(fpath, os.ModePerm)

		} else {

			// Make File
			if err = os.MkdirAll(filepath.Dir(fpath), os.ModePerm); err != nil {
				return err
			}

			outFile, err := os.OpenFile(fpath, os.O_WRONLY|os.O_CREATE|os.O_TRUNC, f.Mode())
			if err != nil {
				return err
			}

			_, err = io.Copy(outFile, rc)

			// Close the file without defer to close before next iteration of loop
			outFile.Close()

			if err != nil {
				return err
			}

		}
	}
	return nil
}
