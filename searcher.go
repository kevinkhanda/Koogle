package main

import (
	"github.com/caneroj1/stemmer"
	"os"
	"strings"
	"bufio"
	"io"
	"strconv"
	"sort"
	"fmt"
	"regexp"
)

var re = regexp.MustCompile("[0-9]+")


func search(query string) string {
	var postingsLists []DeserializedPostingsList
	queryTerms := strings.Split(query, " ")
	for _, term := range queryTerms {
		postingList, err := findTermPostings(term)
		if err == nil {
			postingsLists = append(postingsLists, postingList)
		}
	}
	var resultPostings DeserializedPostingsList
	if len(postingsLists) > 1 {
		i := 0
		for i < len(postingsLists) - 1 {
			resultPostings = mergePostingsLists(postingsLists[0], postingsLists[1])
			postingsLists[1] = resultPostings
			postingsLists = postingsLists[1:]
		}
	} else if len(postingsLists) == 1 {
		resultPostings = postingsLists[0]
	}
	if len(resultPostings) != 0 {
		var result string
		if len(resultPostings) > 20 {
			result = returnPostings(resultPostings[:20])
			result += strconv.Itoa(len(resultPostings)) + " results were found, but top 20 are displayed\n"
		} else {
			result = returnPostings(resultPostings)
			result += strconv.Itoa(len(resultPostings)) + " results were found\n"
		}
		return result
	} else {
		return "Nothing was found! Try to change your query."
	}
}

func mergePostingsLists(first DeserializedPostingsList, second DeserializedPostingsList) (result DeserializedPostingsList) {
	first = sortDeserializedPostingsByDocId(first)
	second = sortDeserializedPostingsByDocId(second)
	small, big := first, second
	if len(small) > len(big) {
		small, big = second, first
	}

	for _, val1 := range small {
		for j, val2 := range big {
			if val1.Key == val2.Key {
				result = append(result, val2)
				big = big[:j+copy(big[j:], big[j+1:])]
				break
			}
		}
	}
	return
}

func returnPostings(deserializedPostingsLists DeserializedPostingsList) (result string) {
	var postingsList PostingsList

	for _, deserializedPosting := range deserializedPostingsLists {
		postingsList = append(postingsList, Posting{Key:deserializedPosting.Key, Value:deserializedPosting.Value})
	}
	sort.Sort(sort.Reverse(postingsList))
	for _, posting := range postingsList {
		documentsFile, err := os.Open("index/documentsIndexes")
		checkError(err)
		documentsReader := bufio.NewReader(documentsFile)
		var docTitle string
		for {
			requiredLine, err := documentsReader.ReadString('|')
			if err != nil {
				break
			}
			documentContent := strings.Split(requiredLine, "->")
			docNumber, err := strconv.ParseInt(re.FindString(documentContent[0]), 10, 64)
			if int(docNumber) == posting.Key {
				docBody := strings.TrimSpace(documentContent[1])
				docTitle = strings.Split(docBody, ".\r")[0]
				break
			}
		}
		result += "Document #" + fmt.Sprintf("%d", posting.Key) + ": " + docTitle + "\n"
		documentsFile.Close()
	}
	return
}

func findTermPostings(term string) (DeserializedPostingsList, error) {
	var result DeserializedPostingsList
	stemList, err := searchForStem(term)
	if err != nil {
		return nil, err
	}
	termIndex, err := findTermIndex(stemList, term)
	if err != nil {
		return nil, err
	}
	result = findPostings(int(termIndex))
	return result, nil
}

func searchForStem(term string) (string, error) {
	stem := stemmer.Stem(term)
	stemFile, err := os.Open("index/stemmingData")
	checkError(err)

	defer stemFile.Close()

	stemFileDataReader := bufio.NewReader(stemFile)
	for {
		fileLine, err := stemFileDataReader.ReadString('\n')
		if err == io.EOF {
			break
		}
		line := fileLine
		lineArray := strings.Split(line, "->")
		if lineArray[0] == stem {
			return lineArray[1], nil
		}
	}
	return "", &errorString{"Nothing is found"}
}

func findTermIndex(rawStemList string, term string) (int64, error) {
	rawStemList = rawStemList[1:len(rawStemList) - 2]  // Substring on pairs by "><"
	for _, stemPair := range strings.Split(rawStemList, "><") {
		stemPairArray := strings.Split(stemPair, ":")
		pairTerm := stemPairArray[0]
		termIndex := stemPairArray[1]
		if pairTerm == term {
			return strconv.ParseInt(termIndex, 10, 64)
		}
	}
	return 0, &errorString{"Nothing is found"}
}

func findPostings(position int) DeserializedPostingsList {
	var result DeserializedPostingsList
	invertedIndexFile, err := os.Open("index/invertedIndex")

	defer invertedIndexFile.Close()

	checkError(err)
	indexReader := bufio.NewReader(invertedIndexFile)
	i := 0
	var requiredLine string
	for i <= position {
		requiredLine, err = indexReader.ReadString('\n')
		checkError(err)
		i++
	}
	rawPostings := strings.Split(requiredLine, "->")[1]
	rawPostings = rawPostings[1:len(rawPostings) - 2]
	for _, rawPosting := range strings.Split(rawPostings, "><") {
		postingValues := strings.Split(rawPosting, ":")
		docId, err := strconv.ParseInt(postingValues[0], 10, 64)
		checkError(err)
		termFrequency, err := strconv.ParseInt(postingValues[1], 10, 64)
		checkError(err)
		posting := DeserializedPosting{int(docId), int(termFrequency)}
		result = append(result, posting)
	}
	return result
}