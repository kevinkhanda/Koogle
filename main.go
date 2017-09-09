package main

import (
	"path/filepath"
)

func main() {
	createInvertedIndex(filepath.Abs("../resources"))
}
