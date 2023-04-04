all: run

run:
	go run ./main.go

build:
	env GOFLAGS=-buildvcs=false GOOS=linux GOARCH=amd64 go build -ldflags="-s -w" -o ./local/backend_linux .
