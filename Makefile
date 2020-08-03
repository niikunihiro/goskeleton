GO_CMD=go
GO_BUILD=$(GO_CMD) build
GO_CLEAN=$(GO_CMD) clean
GO_TEST=$(GO_CMD) test
GO_GET=$(GO_CMD) get
BIN_NAME=goskeleton

all: test build
build:
	$(GO_BUILD) -i -o $(BIN_NAME) -v .
test:
	$(GO_TEST) -v ./...
setup:
	cp etc/go/conf.yaml.sample etc/go/conf.yaml
clean:
	$(GO_CLEAN)
	rm -f $(BIN_NAME)
run:
	$(GO_BUILD) -o $(BIN_NAME) -v .
	./$(BIN_NAME)
