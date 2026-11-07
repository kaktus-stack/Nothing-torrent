BINARY_NAME=nothing-torrent
BUILD_DIR=.
CMD_DIR=./cmd/nothing-torrent

.PHONY: all build install uninstall clean test

all: build

build:
	go build -ldflags="-s -w" -o $(BUILD_DIR)/$(BINARY_NAME) $(CMD_DIR)

install: build
	./$(BINARY_NAME) --install

uninstall:
	./$(BINARY_NAME) --uninstall

clean:
	rm -f $(BUILD_DIR)/$(BINARY_NAME)

test:
	go test -v ./...
