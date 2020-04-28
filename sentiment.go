package main

// Modified from original at: https://github.com/sausheong/gonb

import (
	"bufio"
	"fmt"
	"math/rand"
	"os"
	"regexp"
	"sort"
	"strings"
	"time"

	"github.com/kljensen/snowball"
)

var SentimentClassifier Classifier
var cleaner = regexp.MustCompile(`[^\w\s]`)

type document struct {
	sentiment string
	text      string
}

type sorted struct {
	category    string
	probability float64
}

// Classifier is what we use to classify documents
type Classifier struct {
	words               map[string]map[string]int
	totalWords          int
	categoriesDocuments map[string]int
	totalDocuments      int
	categoriesWords     map[string]int
	threshold           float64
}

// create and initialize the classifier
func createClassifier(categories []string, threshold float64) (c Classifier) {
	c = Classifier{
		words:               make(map[string]map[string]int),
		totalWords:          0,
		categoriesDocuments: make(map[string]int),
		totalDocuments:      0,
		categoriesWords:     make(map[string]int),
		threshold:           threshold,
	}

	for _, category := range categories {
		c.words[category] = make(map[string]int)
		c.categoriesDocuments[category] = 0
		c.categoriesWords[category] = 0
	}
	return
}

// Train the classifier
func (c *Classifier) Train(category string, document string) {
	for word, count := range countWords(document) {
		c.words[category][word] += count
		c.categoriesWords[category] += count
		c.totalWords += count
	}
	c.categoriesDocuments[category]++
	c.totalDocuments++
}

// Classify a document
func (c *Classifier) Classify(document string) (category string) {
	// get all the probabilities of each category
	prob := c.Probabilities(document)

	// sort the categories according to probabilities
	var sp []sorted
	for c, p := range prob {
		sp = append(sp, sorted{c, p})
	}
	sort.Slice(sp, func(i, j int) bool {
		return sp[i].probability > sp[j].probability
	})

	// if the highest probability is above threshold select that
	if sp[0].probability/sp[1].probability > c.threshold {
		category = sp[0].category
	} else {
		category = "unknown"
	}

	return
}

// Probabilities of each category
func (c *Classifier) Probabilities(document string) (p map[string]float64) {
	p = make(map[string]float64)
	for category := range c.words {
		p[category] = c.pCategoryDocument(category, document)
	}
	return
}

// p (document | category)
func (c *Classifier) pDocumentCategory(category string, document string) (p float64) {
	p = 1.0
	for word := range countWords(document) {
		p = p * c.pWordCategory(category, word)
	}
	return p
}

func (c *Classifier) pWordCategory(category string, word string) float64 {
	return float64(c.words[category][stem(word)]+1) / float64(c.categoriesWords[category])
}

// p (category)
func (c *Classifier) pCategory(category string) float64 {
	return float64(c.categoriesDocuments[category]) / float64(c.totalDocuments)
}

// p (category | document)
func (c *Classifier) pCategoryDocument(category string, document string) float64 {
	return c.pDocumentCategory(category, document) * c.pCategory(category)
}

// clean up and split words in document, then stem each word and count the occurrence
func countWords(document string) (wordCount map[string]int) {
	cleaned := cleaner.ReplaceAllString(document, "")
	words := strings.Split(cleaned, " ")
	wordCount = make(map[string]int)
	for _, word := range words {
		if !stopWords[word] {
			key := stem(strings.ToLower(word))
			wordCount[key]++
		}
	}
	return
}

// stem a word using the Snowball algorithm
func stem(word string) string {
	stemmed, err := snowball.Stem(word, "english", true)
	if err == nil {
		return stemmed
	}
	fmt.Println("Cannot stem word:", word)
	return word
}

// set up data for training and testing
func setupData(file string, testPercentage float64, train []document, test []document) ([]document, []document) {
	rand.Seed(time.Now().UTC().UnixNano())
	data, err := readLines(file)
	if err != nil {
		fmt.Println("Cannot read file", err)
		os.Exit(1)
	}
	for _, line := range data {
		s := strings.Split(line, "|")
		doc, sentiment := s[0], s[1]

		if rand.Float64() > testPercentage {
			train = append(train, document{sentiment, doc})
		} else {
			test = append(test, document{sentiment, doc})
		}
	}

	return train, test
}

// read the file line by line
func readLines(path string) ([]string, error) {
	file, err := os.Open(path)
	if err != nil {
		return nil, err
	}
	defer file.Close()

	var lines []string
	scanner := bufio.NewScanner(file)
	for scanner.Scan() {
		lines = append(lines, scanner.Text())
	}
	return lines, scanner.Err()
}

func InitClassifier() {
	var (
		testPercentage = 0.1
		threshold      = 1.1
		categories     = []string{"1", "0"}
	)

	var train []document
	var test []document

	train, test = setupData("static/training/amazon.txt", testPercentage, train, test)
	train, test = setupData("static/training/imdb.txt", testPercentage, train, test)
	train, test = setupData("static/training/yelp.txt", testPercentage, train, test)
	train, test = setupData("static/training/youtube.txt", testPercentage, train, test)

	LogMsg(fmt.Sprintf("Num of docs in TRAIN dataset: %d", len(train)))
	LogMsg(fmt.Sprintf("Num of docs in TEST dataset: %d", len(test)))

	SentimentClassifier = createClassifier(categories, threshold)

	// Train on train dataset
	for _, doc := range train {
		SentimentClassifier.Train(doc.sentiment, doc.text)
	}

	LogMsg(fmt.Sprintf("SentimentClassifier.totalWords: %d", SentimentClassifier.totalWords))

	// Validate on test dataset
	count, accurates, unknowns := 0, 0, 0
	for _, doc := range test {
		count++
		sentiment := SentimentClassifier.Classify(doc.text)
		if sentiment == doc.sentiment {
			accurates++
		}
		if sentiment == "unknown" {
			unknowns++
		}
	}
	LogMsg(fmt.Sprintf("Accuracy on TEST dataset is %2.1f with %2.1f unknowns", float64(accurates)*100/float64(count), float64(unknowns)*100/float64(count)))

	// validate on the first 100 docs in the train dataset
	count, accurates, unknowns = 0, 0, 0
	for _, doc := range train[0:100] {
		count++
		sentiment := SentimentClassifier.Classify(doc.text)
		if sentiment == doc.sentiment {
			accurates++
		}
		if sentiment == "unknown" {
			unknowns++
		}
	}
	LogMsg(fmt.Sprintf("\nAccuracy on TRAIN dataset is %2.1f with %2.1f unknowns", float64(accurates)*100/float64(count), float64(unknowns)*100/float64(count)))
}
