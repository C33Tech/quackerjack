package main

import (
	"bytes"
	"encoding/csv"
	"fmt"
	"log"
	"math"
	"os"
	//"strconv"

	"github.com/eaigner/shield"
)

var buf bytes.Buffer
var shieldInstance shield.Shield

func InitShield() {
	if shieldInstance == nil {
		shieldInstance = shield.New(
			shield.NewEnglishTokenizer(),
			shield.NewRedisStore(*RedisServer, "", log.New(&buf, "logger: ", log.Lshortfile), ""),
		)
	}
}

func LoadTrainingData(path string) {
	csvfile, err := os.Open(path)
	if err != nil {
		fmt.Println(err)
		os.Exit(1)
	}

	defer csvfile.Close()

	reader := csv.NewReader(csvfile)
	reader.FieldsPerRecord = -1

	csvData, err := reader.ReadAll()
	if err != nil {
		fmt.Println(err)
		os.Exit(1)
	}

	InitShield()

	fmt.Println("Learning started.")

	for _, row := range csvData {
		// score, _ := strconv.ParseInt(row[1], 10, 0)
		shieldInstance.Learn(row[1], row[0])
	}

	fmt.Println("Learning complete!")
}

func GetSentiment(text string) string {
	tag, err := shieldInstance.Classify(text)
	if err == nil {
		return tag
	}

	return "UNKNOWN"
}

type SentimentTag struct {
	Name    string
	Percent float64
}

func GetSentimentSummary(comments []Comment) []SentimentTag {
	InitShield()

	tags := map[string]int{}

	for _, comment := range comments {
		tag := GetSentiment(comment.Content)

		tags[tag]++
	}

	result := []SentimentTag{}

	for tag, count := range tags {
		st := SentimentTag{
			Name:    tag,
			Percent: math.Ceil((float64(count) / float64(len(comments))) * float64(100)),
		}

		result = append(result, st)
	}

	return result
}
