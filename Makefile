build: format
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

docker-run:
		docker-compose up

docker-build:
		docker-compose up --build

docker-shell-start:
		docker-compose run --service-ports --rm app /bin/ash

docker-shell-running:
		docker-compose exec app /bin/ash

docker-clean:
		docker-compose down

docker-status:
		docker-compose ps
