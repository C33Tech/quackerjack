FROM golang:1.22-alpine

RUN mkdir -p /go/src/quackerjack && \
    apk update && \
    apk add git && \
    apk add gcc

WORKDIR /go/src/quackerjack

COPY . .

RUN go get github.com/jteeuwen/go-bindata/... && \
    go-bindata -o webgui.go static/gui/ && \
    go build -o quackerjack-docker

EXPOSE 8000

CMD ./scripts/start-docker.sh