#!/bin/sh

if [ ! -f /go/src/quackerjack/quackerjack-docker ]; then
  echo "Missing quackerjack binary."
  exit 1
fi

if [ ! -f /go/src/quackerjack/quackerjack-docker.conf ]; then
  echo "Missing quackerjack config file."
  exit 1
fi

if ! pgrep "quackerjack" > /dev/null; then
  go build -mod=vendor -o quackerjack-docker
  nohup /go/src/quackerjack/quackerjack-docker -server -verbose -conf "/go/src/quackerjack/quackerjack-docker.conf"
fi