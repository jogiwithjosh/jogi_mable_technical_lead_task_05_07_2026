APP=app

run:
	go run ./cmd/main.go

build:
	CGO_ENABLED=0 \
    GOOS=linux \
    GOARCH=amd64 \
    go build \
        -C api \
        -trimpath \
        -ldflags="-s -w" \
        -o api \
        ./cmd

test:
	go test -race ./api/...
	go test -race ./pipeline/...

fmt:
	go fmt ./api/...
	go fmt ./pipeline/...

vet:
	go vet ./api/...
	go vet ./pipeline/...

tidy:
	go mod tidy

clean:
	rm -rf bin
	rm -f coverage.out

benchmark:
	go test -bench=. -benchmem ./pipeline/...

docker-up:
	docker-compose up --build

docker-down:
	docker-compose down

docker-logs:
	docker-compose logs -f

docker-clean:
	docker-compose down -v