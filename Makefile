.PHONY: build run
build:
	go build -o=/tmp/bin/luxury ./cmd/web
	
run: build
	/tmp/bin/luxury