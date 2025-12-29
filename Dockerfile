FROM --platform=$BUILDPLATFORM golang:1.24-alpine AS builder

ARG TARGETOS
ARG TARGETARCH

RUN apk add --no-cache git

WORKDIR /go/src/quackerjack

COPY go.mod go.sum ./
RUN go mod download

COPY . .

RUN CGO_ENABLED=0 GOOS=${TARGETOS} GOARCH=${TARGETARCH} go build -o quackerjack-docker

FROM alpine:latest

RUN apk --no-cache add ca-certificates

WORKDIR /go/src/quackerjack

COPY --from=builder /go/src/quackerjack/quackerjack-docker .
COPY --from=builder /go/src/quackerjack/scripts ./scripts
COPY --from=builder /go/src/quackerjack/static ./static
COPY --from=builder /go/src/quackerjack/quackerjack.example.conf .

EXPOSE 8000

CMD ["./scripts/start-docker.sh"]
