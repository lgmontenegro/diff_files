package main

import (
	"bufio"
	"fmt"
	"log"
	"os"
	"strings"
)

type FileInfo struct {
	FileArray []string
}

func (f *FileInfo) ReadData(fileToOpen string) {
	files, err := os.Open(fileToOpen)
	if err != nil {
		fmt.Println(err)
	}
	defer files.Close()

	scannerFiles := bufio.NewScanner(files)
	for scannerFiles.Scan() {
		f.FileArray = append(f.FileArray, scannerFiles.Text())
	}

	if err := scannerFiles.Err(); err != nil {
		log.Fatal(err)
	}
}

func (fi *FileInfo) DiffFrom(fileToCompare []string) (missing []string) {
	for _, fc := range fileToCompare {
		for i, f := range fi.FileArray {
			if strings.Trim(fc, " ") == strings.Trim(f, " ") {
				break
			}

			if i+1 == len(fi.FileArray) {
				missing = append(missing, fc)
			}
		}
	}

	return missing
}
