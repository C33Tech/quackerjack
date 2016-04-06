package main

import (
	"bytes"
	"encoding/csv"
	"fmt"
	"log"
	"os"
	"time"

	"github.com/eaigner/shield"
	//"github.com/kr/pretty"
)

var buf bytes.Buffer
var shieldInstance shield.Shield

// InitShield instantiates the text classifier engine
func InitShield() {
	if shieldInstance == nil {
		shieldInstance = shield.New(
			shield.NewEnglishTokenizer(),
			shield.NewRedisStore(GetConfigString("redis"), "", log.New(&buf, "logger: ", log.Lshortfile), ""),
		)

		// Start process to monitor redis connection
		go func() {
			for {
				shieldInstance.TestConnection()

				d, _ := time.ParseDuration("15m")
				time.Sleep(d)
			}
		}()
	}
}

// LoadTrainingData input the training data in to the text classifier
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

// GetSentiment classifies a single string of text. Returns the tag it matched.
func GetSentiment(text string) string {
	InitShield()

	tag, err := shieldInstance.Classify(text)
	if err == nil {
		return tag
	}

	LogMsg(err.Error())

	return "UNKNOWN"
}

// SentimentTag is a list entry of the tag and the percent of comments that were classified with that tag.
type SentimentTag struct {
	Name    string
	Percent float64
}
