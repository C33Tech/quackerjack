# `Quackerjack`
## A YouTube Comment Thread Analyzer

A CLI script, written in Go, that takes a YouTube video ID and generates a report about the content of the comment stream.

![Web GUI](/static/web-gui.png)

The report includes:
* The total number of comments and how many were collected.
* The top keywords from the comments.
* The sentiment analysis of the comments.

## Requirements

* go (`brew`, `apt-get` or https://golang.org/doc/install)
* redis (`brew`, `apt-get` or http://redis.io/topics/quickstart)

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

## Training

Before you can analyze any comments, you'll need to train your semantic engine by starting up your Redis server and loading the supplied training data.

`quackerjack -training ./static/training/afinn-111.csv`

You can always add in your own training data by creating a csv file with two fields: word, tag (in between -5 and 5)

## Running

`quackerjack -stopwords ./static/stopwords.txt[,/comma/delimited/textfiles] -video [video id] [-verbose]`

## Who is Quackerjack?

He's a villain from Darkwing Duck.

