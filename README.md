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

`go get github.com/mikeflynn/quackerjack`

## Flags

`-video XXXX` _YouTube video id_

`-ytkey xxxxxxxxxxxxxxxx` _Your Google API key with access to the YouTube Data API_

`-stopwords /xxx/yy/zz.txt` _A comma delimitated list of additional stop word text files not in config_

`-verbose` _Standard boolean flag for extra logging to std out._

`-training /xxx/yy/zz.txt` _Load a list of training words for the semantic analysis._

`-redis "127.0.0.1:6379"` _A string containing the host and port of the redis server._

`-server` _Start the web interface rather than having the JSON get dumpped to std out._

`-port 8000` _Override the web server port._

## Running

`quackerjack -stopwords ./static/stopwords.txt[,/comma/delimited/textfiles] -video [video id] [-verbose]`

## Future Additions

* Top commenters list.
* More comment stats (avg per day, biggest commenting day, ...)
