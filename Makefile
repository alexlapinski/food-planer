BINARY_NAME=menu-planner

all: build
build:
	go build -o $(BINARY_NAME) ./cmd/generator
clean:
	go clean
	rm -f $(BINARY_NAME)