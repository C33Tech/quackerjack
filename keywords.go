package main

import (
	"bufio"
	"os"
	"regexp"
	"sort"
	"strings"
)

// GetKeywords takes an array of Comments and returns a list of the top 10 most frequent words.
func GetKeywords(comments []Comment) []string {
	idx := make(map[string]int)

	for _, comment := range comments {
		words := getWords(comment.Content)

		for _, word := range words {
			if !isStopWord(word) {
				idx[strings.ToLower(word)]++
			}
		}
	}

	sorted := sortedKeys(idx)

	return sorted[:10]
}

func getWords(text string) []string {
	words := regexp.MustCompile("\\w+")
	return words.FindAllString(text, -1)
}

func isStopWord(needle string) bool {
	if stopWords[strings.ToLower(needle)] {
		return true
	}

	return false
}

var stopWords = make(map[string]bool)

// LoadStopWords takes list of paths and reads the contents in to the stopWords array.
func LoadStopWords(path string) {
	file, err := os.Open(path)
	if err != nil {
		return
	}
	defer file.Close()

	scanner := bufio.NewScanner(file)
	for scanner.Scan() {
		stopWords[scanner.Text()] = true
	}
}

// Sorting a map via: https://gist.github.com/ikbear/4038654

type sortedMap struct {
	m map[string]int
	s []string
}

func (sm *sortedMap) Len() int {
	return len(sm.m)
}

func (sm *sortedMap) Less(i, j int) bool {
	return sm.m[sm.s[i]] > sm.m[sm.s[j]]
}

func (sm *sortedMap) Swap(i, j int) {
	sm.s[i], sm.s[j] = sm.s[j], sm.s[i]
}

func sortedKeys(m map[string]int) []string {
	sm := new(sortedMap)
	sm.m = m
	sm.s = make([]string, len(m))
	i := 0
	for key := range m {
		sm.s[i] = key
		i++
	}
	sort.Sort(sm)
	return sm.s
}
