.DEFAULT: all

.PHONY: all
all: build run

.PHONY: gen
gen:
	mkdir -p mocks
	go generate ./...

.PHONY: build
build:
	go build

.PHONY: test
test:
	go test -v ./...

.PHONY: run
run:
	./swagger-demo