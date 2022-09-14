package main

import (
	"bufio"
	"fmt"
	"os"
	"strings"
)

func main() {

	var fsFiles = make([]string, 0)
	var dbFiles = make([]string, 0)
	var missing = make([]string, 0)

	files, err := os.Open("files")
	if err != nil {
		fmt.Println(err)
	}
	defer files.Close()

	scannerFiles := bufio.NewScanner(files)
	for scannerFiles.Scan() {
		fsFiles = append(fsFiles, scannerFiles.Text())
	}

	if err := scannerFiles.Err(); err != nil {
		fmt.Println(err)
	}

	filesDB, err := os.Open("filenamesInDb")
	if err != nil {
		fmt.Println(err)
	}
	defer filesDB.Close()

	scannerDB := bufio.NewScanner(filesDB)
	for scannerDB.Scan() {
		dbFiles = append(dbFiles, scannerDB.Text())
	}

	if err := scannerDB.Err(); err != nil {
		fmt.Println(err)
	}

	for _, db := range dbFiles {
		for i, fs := range fsFiles {
			if strings.Trim(db, " ") == strings.Trim(fs, " ") {
				break
			}

			if i+1 == len(fsFiles) {
				missing = append(missing, db)
			}
		}
	}

	fmt.Printf("total %d \n", len(missing))

	for i, miss := range missing {
		fmt.Printf("%d, %s\n", i, miss)
	}

	strings.

}
