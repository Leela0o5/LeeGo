build:
    go build -o bin/ws-tester main.go

run:
    go run main.go

test:
    go test ./...

clean:
    rm -rf bin/
