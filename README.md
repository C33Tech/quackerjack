# `Quackerjack` v2.0
## A Comment Thread Analyzer and Report Generator

A CLI script, written in Go, that takes a YouTube video (or Instagram post) URL and generates a report about the content of the comment stream.

![Web GUI](/static/web-gui.png)

The report includes:
* The total number of comments and how many were collected.
* The top keywords from the comments.
* The sentiment analysis of the comments.
* A random sampling of the comments.

## Requirements

* go (`brew`, `apt-get` or https://golang.org/doc/install)
* redis (`brew`, `apt-get` or http://redis.io/topics/quickstart)

## Installation

`go get github.com/mikeflynn/quackerjack`

## Running

The script can be run via the command line, which will return a JSON object of the resulting data:

`quackerjack -stopwords ./static/stopwords.txt[,/comma/delimited/textfiles] -post [post url] [-verbose]`

You can also run quackerjack as a web service. Add the `-server` flag to start a web server that has two endpoints on port 8080: `/` is a simple web interface where you can enter post urls and it displays the results in a visual report. `/api` is the JSON web service route which takes a `?vid=` param and retuns a JSON response.

## Flags

`-post XXXX` _YouTube, or Instagram post URL._

`-ytkey xxxxxxxxxxxxxxxx` _Your Google API key with access to the YouTube Data API._

`-igkey xxxxxxxxxxxxxxxx` _Your Instagram API key with._

`-stopwords /xxx/yy/zz.txt` _A comma delimitated list of additional stop word text files not in config._

`-verbose` _Standard boolean flag for extra logging to std out._

`-training /xxx/yy/zz.txt` _Load a list of training words for the semantic analysis._

`-redis "127.0.0.1:6379"` _A string containing the host and port of the redis server._

`-server` _Start the web interface rather than having the JSON get dumpped to std out._

`-port 8000` _Override the web server port._

## Training

Before you can analyze any comments, you'll need to train your semantic engine by starting up your Redis server and loading the supplied training data.

`quackerjack -training ./static/training/afinn-111.csv`

You can always add in your own training data by creating a csv file with two fields: word, tag (in between -5 and 5)

## Development

If you want to dig in to the code, you can clone this repo...

`git clone https://github.com/mikeflynn/quackerjack.git`

...tweak until your heart's content, and then build your new version with the included Makefile...

`make`

...which installs all the dependencies, formats your code and builds the web GUI HTML file in to go code, then generates a `quackerjack` binary.

## Who is Quackerjack?

He's a villain from Darkwing Duck.

