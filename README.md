# `Quackerjack`
## A YouTube Comment Thread Analyzer

A CLI script, written in Go, that takes a YouTube video ID and generates a report about the content of the comment stream.

![Quackerjack](/static/quackerjack.png)

The report includes:
* The total number of comments.
* The average number of comments a day.
* The top commenters (based on subscribers)
* The top keywords from the comments.
* The sentiment analysis of the comments.

## Installation
`export YOUTUBE_KEY=xxxxxxxxxxxxxxx`
`go get github.com/mikeflynn/quackerjack`

## Running

`quackerjack -stopwords ./static/stopwords.txt[,/comma/delimited/textfiles] -video [video id] [-verbose]`

## To Do

* Optional sentiment analysis using Google's Prediction API.
* Top commenters list.
* More comment stats (avg per day, biggest commenting day, ...)
