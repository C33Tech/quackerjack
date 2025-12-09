#!/bin/sh

if [ ! -f /go/src/quackerjack/quackerjack-docker ]; then
  echo "Missing quackerjack binary."
  exit 1
fi

if [ ! -f /go/src/quackerjack/docker.conf ]; then
  echo "Missing quackerjack config file."
  exit 1
fi

nohup /go/src/quackerjack/quackerjack-docker -server -conf "/go/src/quackerjack/docker.conf"
