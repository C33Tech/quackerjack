build: format
		@echo "--> Building"
		@go build

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
