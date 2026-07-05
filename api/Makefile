APP=api

run:
	go run ./cmd/server

build:
	go build -o bin/$(APP) ./cmd/server

test:
	go test -race ./...

fmt:
	go fmt ./...

vet:
	go vet ./...

tidy:
	go mod tidy

clean:
	rm -rf bin
	rm -f coverage.out