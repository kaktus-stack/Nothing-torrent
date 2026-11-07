BINARY_NAME=nothing-torrent
BUILD_DIR=.
CMD_DIR=./cmd/nothing-torrent

.PHONY: all build build-windows build-all install uninstall clean test

all: build

build:
	go build -ldflags="-s -w" -o $(BUILD_DIR)/$(BINARY_NAME) $(CMD_DIR)

build-windows:
	CGO_ENABLED=0 GOOS=windows GOARCH=amd64 go build -ldflags="-s -w" -o $(BUILD_DIR)/$(BINARY_NAME).exe $(CMD_DIR)

build-all: build build-windows
	CGO_ENABLED=0 GOOS=darwin GOARCH=arm64 go build -ldflags="-s -w" -o $(BUILD_DIR)/$(BINARY_NAME)-darwin-arm64 $(CMD_DIR)
	CGO_ENABLED=0 GOOS=darwin GOARCH=amd64 go build -ldflags="-s -w" -o $(BUILD_DIR)/$(BINARY_NAME)-darwin-amd64 $(CMD_DIR)

install: build
	./$(BINARY_NAME) --install

uninstall:
	./$(BINARY_NAME) --uninstall

clean:
	rm -f $(BUILD_DIR)/$(BINARY_NAME) $(BUILD_DIR)/$(BINARY_NAME).exe $(BUILD_DIR)/$(BINARY_NAME)-darwin-*

test:
	go test -v ./...
