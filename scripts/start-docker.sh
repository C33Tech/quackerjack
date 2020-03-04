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
  /go/src/quackerjack/quackerjack-docker -training /go/src/quackerjack/static/training/afinn-111.csv -redis "redis:6379"
  nohup /go/src/quackerjack/quackerjack-docker -server -redis "redis:6379" -conf "/go/src/quackerjack/quackerjack-docker.conf"
fi