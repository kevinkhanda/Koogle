package main

import "sort"

func sortMapByValue(wordFrequencies map[int]int) PostingsList {
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
