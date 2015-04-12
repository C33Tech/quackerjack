#!/bin/bash

if [ -z $1 ]; then
  RT="-server"
else
  RT="-post $1"
fi

if [ -z $2 ]; then
  SW=""
else
  SW=",$2"
fi

YOUTUBE_KEY=xxxxxxxxxxxxxxx
INSTAGRAM_KEY=xxxxxxxxxxxxxxx

./quackerjack -verbose -ytkey $YOUTUBE_KEY -igkey $INSTAGRAM_KEY -stopwords ./static/stopwords.txt$SW $RT
