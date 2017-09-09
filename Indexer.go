package main

import (
	"fmt"
	"os"
	"io/ioutil"
	//"github.com/caneroj1/stemmer"
	"strings"
	"text/scanner"
	"strconv"
	//"time"
	//"regexp/syntax"
	"regexp"
	"bytes"
	//"github.com/caneroj1/stemmer"
	//"github.com/emirpasic/gods/sets/treeset"
	//"time"
)


var IsValidString = regexp.MustCompile("[a-z]+$|[0-9]+$").MatchString

var stemmingMap = make(map[string]map[string]int) // Map of type "STEM" -> {"term1" -> Position in Indexes file}
var indexesMap = make(map[string]map[int]int)     // Map of type "term" -> map{docId -> termFrequency}
var sortedIndexesMap = make(map[string][]Posting) // Map of type "term" -> []{Posting1 (docId, termFrequency)}
var offsetsMap = make(map[string]int)

var scan scanner.Scanner

func checkError(err error)  {
	if err != nil {
		panic(err)
	}
}

func createInvertedIndex(packageToScan string, err error) {
	files, err := ioutil.ReadDir(packageToScan)
	checkError(err)

	for _, file := range files {
		analyzedDocuments := analyzeDocuments(packageToScan, file)
		tokenizeDocuments(analyzedDocuments)
	}
	for key, value := range indexesMap {
		sortedIndexesMap[key] = sortMapByValue(value) // Sort indexes by their term frequencies
	}
	//for key, value := range sortedIndexesMap {
	//	fmt.Printf("Key = %s; Value = ", key)
	//	fmt.Println(value)
	//}
	createInvertedIndexFile()
}

func analyzeDocuments(packageName string, file os.FileInfo) map[int]string {
	result := make(map[int]string)
	re := regexp.MustCompile("[0-9]+")
	fileData, err := ioutil.ReadFile(packageName + "/" + file.Name())
	checkError(err)
	documents := strings.Split(string(fileData), "********************************************")
	for _, document := range documents {
		trimmedDocument := strings.TrimSpace(document)
		title := strings.Split(trimmedDocument, "\n")[0]
		content := strings.Join(strings.Split(trimmedDocument, "\n")[1:], "\n")
		if title != "" {
			docId, err := strconv.ParseInt(re.FindString(title), 10, 64)
			checkError(err)
			result[int(docId)] = content
		}
	}
	return result
}

func tokenizeDocuments(analyzedDocuments map[int] string) {
	for key, value := range analyzedDocuments {
		scan.Init(strings.NewReader(value))
		for token := scan.Scan(); token != scanner.EOF; token = scan.Scan() {
			term := strings.ToLower(scan.TokenText())
			if IsValidString(term) {
				if val, ok := indexesMap[term]; ok {  // If term exists as key
					if value, newOk := val[key]; newOk {  // If docId exists as a key
						val[key] = value + 1
					} else {
						val[key] = 1
					}
					indexesMap[term] = val
				} else {
					tfMap := make(map[int]int)
					tfMap[key] = 1
					indexesMap[term] = tfMap
				}
			}
		}
	}
}


func createStemFile() {
	//stem := stemmer.Stem(term)
	//if val, ok := stemmingMap[stem]; ok {
	//	stemmingMap[stem] = append(val, term)  // Appending to existing
	//} else {
	//	stemmingMap[stem] = []string{term}  // Creating new
	//}
}

func createInvertedIndexFile() {
	indexesFile, err := os.Create("index/invertedIndex")
	checkError(err)

	defer indexesFile.Close()

	for key, value := range sortedIndexesMap {
		position := 0
		var buffer bytes.Buffer
		buffer.WriteString(key + ":")
		for _, pair := range value {
			buffer.WriteString(fmt.Sprintf("<%d:%d>", pair.Key, pair.Value))
		}
		buffer.WriteString("\n")
		offsetsMap[key] = position
		position += buffer.Len()
		indexesFile.Write(buffer.Bytes())
	}
}