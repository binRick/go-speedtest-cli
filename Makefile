GOCMD=go
GO_VERSION=1.16.4
BINARY_DEST_SUB_DIR=bin
BINARY_NAME=go-speedtest-cli

clean:
	rm $(BINARY_DEST_SUB_DIR)/$(BINARY_NAME)
	rm $(BINARY_DEST_SUB_DIR)/$(BINARY_NAME)-static
	
binary:
	$(GOCMD) build -o $(BINARY_DEST_SUB_DIR)/$(BINARY_NAME) .

build: binary

test:
	$(BINARY_DEST_SUB_DIR)/$(BINARY_NAME) --help
	$(BINARY_DEST_SUB_DIR)/$(BINARY_NAME)-static --help

static:
	$(GOCMD) build -a -ldflags "-linkmode external -extldflags '-static' -s -w" -o $(BINARY_DEST_SUB_DIR)/$(BINARY_NAME)-static .

all: binary static
