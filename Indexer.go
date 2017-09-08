package main

import (
	"fmt"
	"io/ioutil"
	"log"
)

func createInvertedIndex(packageToScan string, err error) {
	files, err := ioutil.ReadDir(packageToScan)
	if err != nil {
		log.Fatal(err)
	}

	for _, file := range files {
		fmt.Println(file.Name())
	}
}