all: build

build:
	go build -o bin/starvo-backend ./app

test:
	go test -v ./review/... && \
	go test -v ./ocpp/... && \
	go test -v ./domain/... && \
	go test -v ./transaction/... && \
	go test -v ./external/... && \
	go test -v ./machinecheck/... && \
	go test -v ./ads/... && \
	go test -v ./kvstore/... && \
	go test -v ./price_settings/... && \
	go test -v ./station/... 

dockerbuild:
	go build ./app/main.go
	cp main app/app
	docker build --platform linux/amd64 -f DockerfileLocal . -t pos
	docker compose -f docker-compose-local.yaml up --build -d
	docker logs -f pos
	
dockerbuildmac:
	env GOOS=linux GOARCH=amd64 go build ./app/main.go
	cp main app/app
	docker buildx build --platform linux/amd64 -f DockerfileLocal . -t starvo
	docker compose -f docker-compose-local.yaml up --build -d