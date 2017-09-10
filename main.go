package main

import (
	"path/filepath"
	"fmt"
	"time"
	"bufio"
	"os"
	"strings"
)

func main() {
	fmt.Println("█░█ █▀▀█ █▀▀█ █▀▀▀ █░░ █▀▀   █▀▀ █▀▀ █▀▀█ █▀▀█ █▀▀ █░░█   █▀▀ █▀▀▄ █▀▀▀ ░▀░ █▀▀▄ █▀▀")
	fmt.Println("█▀▄ █░░█ █░░█ █░▀█ █░░ █▀▀   ▀▀█ █▀▀ █▄▄█ █▄▄▀ █░░ █▀▀█   █▀▀ █░░█ █░▀█ ▀█▀ █░░█ █▀▀")
	fmt.Println("▀░▀ ▀▀▀▀ ▀▀▀▀ ▀▀▀▀ ▀▀▀ ▀▀▀   ▀▀▀ ▀▀▀ ▀░░▀ ▀░▀▀ ▀▀▀ ▀░░▀   ▀▀▀ ▀░░▀ ▀▀▀▀ ▀▀▀ ▀░░▀ ▀▀▀")
	fmt.Println("Hello! Koogle is going to index files from the resource folder in a second!")
	time.Sleep(1000 * time.Millisecond)
	createInvertedIndex(filepath.Abs("resources"))
	fmt.Println("Files were indexed successfully! Indexes are stored in 'index' folder.")
	time.Sleep(1000 * time.Millisecond)
	fmt.Println("Now you can type your query or exit (type 'Koogle exit').")
	scanner := bufio.NewScanner(os.Stdin)
	for {
		fmt.Print("Type your query: ")
		scanner.Scan()
		userQuery := scanner.Text()
		if userQuery == "Koogle exit" {
			os.Exit(0)
		}
		userQuery = strings.TrimSpace(userQuery)
		userQuery = strings.ToLower(userQuery)
		queryResults := search(userQuery)
		fmt.Println(queryResults)
	}
}
