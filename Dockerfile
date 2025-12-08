FROM golang:1.24-alpine

RUN mkdir -p /go/src/quackerjack && \
  apk update && \
  apk add git && \
  apk add gcc && \
  apk add libc-dev

WORKDIR /go/src/quackerjack

COPY . .

RUN go build -o quackerjack-docker

EXPOSE 8000

CMD ./scripts/start-docker.sh
