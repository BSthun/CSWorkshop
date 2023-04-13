all: run

run:
	go run .

build:
	GOOS=linux GOARCH=amd64 go build -o local/backend_linux .

test:
	go test -v ./...

test-integration:
	go test -v ./tests -tags=integration

firebaseemu:
	(cd ../firebase/ && firebase emulators:start --only auth --import=./emulator-data --export-on-exit)

chrome:
	"/Applications/Google Chrome.app/Contents/MacOS/Google Chrome" --remote-debugging-port=9222
