REST_MAIN = cmd/rest.go

.PHONY: test run build docker-build
test:
	 go test ./...

run:
	go run $(REST_MAIN)

build:
	GOOS=linux GOARCH=amd64 CGO_ENABLED=0 go build -o main $(REST_MAIN)

docker-build:
	docker build .