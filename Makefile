build:
	go build -o dist/ad2bs cmd/ad2bs/main.go
	go build -o dist/bs2ad cmd/bs2ad/main.go

test:
	go test -v ./...

# Go specific make commands

go-get:
	go get ./...

go-clean:
	go clean -cache -modcache

.PHONY: build test go-get go-clean
