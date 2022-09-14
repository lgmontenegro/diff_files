package main

import (
	"fmt"
	"log"
	"os"
)

func main() {

	fileFs := FileInfo{}
	fileDb := FileInfo{}

	if len(os.Args) < 3 {
		log.Fatal("Need 2 files to compare")
	}

	fileDb.ReadData(os.Args[1])
	fileFs.ReadData(os.Args[2])
	
	missing := fileFs.DiffFrom(fileDb.FileArray)

	fmt.Printf("total %d \n", len(missing))

	for i, miss := range missing {
		fmt.Printf("%d, %s\n", i, miss)
	}
}
