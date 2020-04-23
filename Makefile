IS_RUNNING := $(shell docker ps -q --no-trunc | grep $(shell docker-compose ps -q app))

build: deps format data
		@echo "--> Building"
		@go build

mod-init:
		@echo "--> Installing build dependencies"
		@go get github.com/google/google-api-go-client/googleapi/transport
		@go get google.golang.org/api/youtube/v3
		@go get github.com/cdipaolo/sentiment
		@go get go get -u github.com/jteeuwen/go-bindata/...
		@go mod vendor

format:
		@echo "--> Running go fmt"
		@gofmt -s -w .

data:
		@echo "--> Importing binary files"
		@go-bindata -o webgui.go static/gui/

docker-run:
		docker-compose up

docker-build:
		docker-compose up --build

docker-shell:
ifeq ($(IS_RUNNING),)
		docker-compose run --service-ports --rm app /bin/ash
else
		docker-compose exec app /bin/ash
endif
		@echo "Shell closed."

docker-clean:
		docker-compose down

docker-status:
		docker-compose ps
