package main

import (
	"sort"
	"log"
)

type errorString struct {
	s string
}

func (e *errorString) Error() string {
	return e.s
}

func checkError(err error)  {
	if err != nil {
		log.Fatal(err)
	}
}

func createStemPairsList(stemmingMap map[string]map[string]int) map[string]StemPairList {
	result := make(map[string]StemPairList)
	for k, wordPositions := range stemmingMap {
		stemList := make(StemPairList, len(wordPositions))
		i := 0
		for key, value := range wordPositions {
			stemList[i] = StemPair{key, value}
			i++
		}
		result[k] = stemList
	}
	return result
}

type StemPair struct {
	Key string
	Value int
}

type StemPairList []StemPair

func sortPostingsByTermFrequency(wordFrequencies map[int]int) PostingsList {
	pl := make(PostingsList, len(wordFrequencies))
	i := 0
	for k, v := range wordFrequencies {
		pl[i] = Posting{k, v}
		i++
	}
	sort.Sort(sort.Reverse(pl))
	return pl
}

type Posting struct {
	Key int
	Value int
}

type PostingsList []Posting

func (p PostingsList) Len() int           { return len(p) }
func (p PostingsList) Less(i, j int) bool { return p[i].Value < p[j].Value }
func (p PostingsList) Swap(i, j int)      { p[i], p[j] = p[j], p[i] }


type DeserializedPosting struct {
	Key   int
	Value int
}

func sortDeserializedPostingsByDocId(postings DeserializedPostingsList) DeserializedPostingsList {
	sort.Sort(postings)
	return postings
}

type DeserializedPostingsList []DeserializedPosting

func (p DeserializedPostingsList) Len() int           { return len(p) }
func (p DeserializedPostingsList) Less(i, j int) bool { return p[i].Key < p[j].Key }
func (p DeserializedPostingsList) Swap(i, j int)      { p[i], p[j] = p[j], p[i] }