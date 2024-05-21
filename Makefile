build:
	goreleaser release --clean --skip=publish

clean:
	rm -rf dist

# Go specific make commands

go-get:
	go get ./...

go-clean:
	go clean -cache -modcache

.PHONY: build clean go-get go-clean