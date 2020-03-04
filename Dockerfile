FROM golang:1.14-alpine

RUN mkdir -p /go/src/quackerjack && \
    apk update && \
    apk add git && \
    apk add gcc

WORKDIR /go/src/quackerjack

COPY . .

RUN go get github.com/jteeuwen/go-bindata/... && \
    go-bindata -o webgui.go static/gui/ && \
    go mod vendor && \
    go build -mod=vendor -o quackerjack-docker

EXPOSE 8000

CMD ./scripts/start-docker.sh