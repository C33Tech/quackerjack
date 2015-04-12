all: deps format data
		@echo "--> Building"
		@go build

deps:
		@echo "--> Installing build dependencies"
		@go get github.com/mikeflynn/golang-instagram/instagram
		@go get code.google.com/p/google-api-go-client/googleapi/transport
		@go get google.golang.org/api/youtube/v3
		@go get github.com/eaigner/shield

format:
		@echo "--> Running go fmt"
		@gofmt -s -w .

data:
		@echo "--> Importing binary files"
		@go-bindata -o webgui.go static/gui/