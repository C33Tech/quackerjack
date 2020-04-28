# `Quackerjack` v3
## A Comment Thread Analyzer and Report Generator

A CLI script and web service, written in Go, that takes a YouTube video URL and generates a report about the content of the comment stream.

![Web GUI](/static/web-gui.png)

The report includes:
* The total number of comments and how many were collected.
* The top keywords from the comments.
* The sentiment analysis of the comments.
* Sentiment over the first 30 days from upload.
* Emoji usage analysis.

v3 Changes:
* New interface!
* No more redis!
* Emoji analysis!

*NOTE: Submissions of training data is more than welcome! YouTube comments are a hard thing to analyze so the more data we can collect the better! Please submit a pull request or you can reach out to me on Twitter: @thatmikeflynn*

## Installation

`go get github.com/mikeflynn/quackerjack`

## Development

* go 1.14
* Docker

## Running

The script can be run via the command line, which will return a JSON object of the resulting data:

`quackerjack -stopwords ./static/stopwords.txt[,/comma/delimited/textfiles] -post [post url] [-verbose]`

You can also run quackerjack as a web service. Add the `-server` flag to start a web server that has two endpoints on port 8080: `/` is a simple web interface where you can enter post urls and it displays the results in a visual report. `/api` is the JSON web service route which takes a `?vid=` param and retuns a JSON response.


## Flags

`-post XXXX` _YouTube, or Instagram post URL._

`-ytkey xxxxxxxxxxxxxxxx` _Your Google API key with access to the YouTube Data API._

`-stopwords /xxx/yy/zz.txt` _A comma delimitated list of additional stop word text files not in config._

`-verbose` _Standard boolean flag for extra logging to std out._

`-server` _Start the web interface rather than having the JSON get dumpped to std out._

`-port 8000` _Override the web server port._

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

...tweak until your heart's content, and then build your new version with the included Makefile...

`make`

...which installs all the dependencies, formats your code and builds the web GUI HTML file in to go code, then generates a `quackerjack` binary.

## To Do:

* Save a copy of the resulting model for easy `go get` download and startup.


## Who is Quackerjack?

He's a villain from Darkwing Duck.

