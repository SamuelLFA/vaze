GO=go
MAIN=cmd/main.go
BIN=bin/main

build:
	$(GO) build -o $(BIN) $(MAIN)

run:
	$(GO) run $(MAIN)

