# `Quackerjack` v4

## A Comment Thread Analyzer and Report Generator

A CLI script and web service, written in Go, that takes a YouTube video URL and generates a report about the content of the comment stream.

![Web GUI](/static/web-gui.png)

The report includes:

- The total number of comments and how many were collected.
- The top keywords from the comments.
- The sentiment analysis of the comments.
- Sentiment over the first 30 days from upload.
- Emoji usage analysis.

v4 Changes:

- New theme
- Easier hosting
- Caching

v3 Changes:

- New interface
- Mobile friendly
- No more redis
- Emoji analysis

_NOTE: Submissions of training data is more than welcome! YouTube comments are a hard thing to analyze so the more data we can collect the better! Please submit a pull request or you can reach out to me on Mastodon: [@hydrox@defcon.social](https://defcon.social/@hydrox)_

## Installation

There are two options to use Quackerjack via either the web or cli:

1. You can install the go binary directly via `go get`:

`go get github.com/mikeflynn/quackerjack`

2. You can use docker with the [latest public image.](https://hub.docker.com/r/mikeflynn/quackerjack)

Here's an example docker-compose.yaml file.

```
services:
  quackerjack:
    image: mikeflynn/quackerjack:latest
    restart: unless-stopped
    volumes:
      - ./quackerjack.conf:/go/src/quackerjack/docker.conf
    ports:
      - "8111:8000"
```

Either option will require a config file that lists the required AI keys such as:

```
server = true
port = 8000
ytkey = xxxxxxxxxxxxxx
stopwords = /go/src/quackerjack/static/stopwords.txt
cache_ttl = 3600
```

## Running

The best way to deploy it is via Docker and the docker-compose.yaml file.

To run it locally for development, use the following command:

`go build && ./quackerjack -server -verbose`

This will start the web server on port 8000.

The script can be run via the command line, which will return a JSON object of the resulting data:

`quackerjack -stopwords ./static/stopwords.txt[,/comma/delimited/textfiles] -post [post url] [-verbose]`

You can also run quackerjack as a web service. Add the `-server` flag to start a web server that has two endpoints on port 8000: `/` is a simple web interface where you can enter post urls and it displays the results in a visual report. `/api` is the JSON web service route which takes a `?vid=` param and retuns a JSON response.

## Flags

`-post XXXX` _YouTube, or Instagram post URL._

`-ytkey xxxxxxxxxxxxxxxx` _Your Google API key with access to the YouTube Data API._

`-stopwords /xxx/yy/zz.txt` _A comma delimitated list of additional stop word text files not in config._

`-verbose` _Standard boolean flag for extra logging to std out._

`-server` _Start the web interface rather than having the JSON get dumpped to std out._

`-port 8000` _Override the web server port._

`-html /path/to/html/file.html` _Overrides the built in html interface with a custom one._

## Training

The new internal analysis engine is trained automatically when the Quackerjack service is started. The training documents are located in static/training and are simply a line of text and a 1 (positive) or 0 (negative) separated by a pipe. To improve the analysis, simply add as many training lines as you'd like and restart Quackerjack (from the git repo directory).

```
app_1  | 2020/04/28 15:56:08 Loading conf file: /go/src/quackerjack/quackerjack-docker.conf
app_1  | 2020/04/28 15:56:08 Num of docs in TRAIN dataset: 2730
app_1  | 2020/04/28 15:56:08 Num of docs in TEST dataset: 310
app_1  | 2020/04/28 15:56:09 SentimentClassifier.totalWords: 17925
app_1  | 2020/04/28 15:56:09 Accuracy on TEST dataset is 83.5 with 1.9 unknowns
app_1  | 2020/04/28 15:56:09
app_1  | Accuracy on TRAIN dataset is 94.0 with 0.0 unknowns
app_1  | 2020/04/28 15:56:09 Web server running on 8000
```

## Development

If you want to dig in to the code, you can clone this repo...

`git clone https://github.com/mikeflynn/quackerjack.git`

The way to run it locally for development is:

`go build && ./quackerjack -server -verbose`

I've included a Docker image and docker-compose config that will run the web service in verbose mode by default. You can also run `make docker-shell` to get a shell on the Docker instance for active development.

## To Do:

- Save a copy of the resulting model for easy `go get` download and startup.

## Who is Quackerjack?

He's a villain from Darkwing Duck.
