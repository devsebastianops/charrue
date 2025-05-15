.PHONY: all run build test

all: build

run:
	@go run main.go build

build:
	@go build -o bin/charrue main.go

install:
	@go install

test:
	@go test -v ./...


run-examples:
	@echo "Running examples..."
	@echo "Example 1: Simple template"
	@make run-example-1
	@echo "Example 2: Looping through a list"
	@make run-example-2
	@echo "Example 3: Conditional rendering"
	@make run-example-3

run-example-1:
	@go run main.go build --values examples/01-simple/data.yaml \
		--template-path examples/01-simple \
		--output-path examples/01-simple/dist \
		--verbose

run-example-2:
	@go run main.go build --values examples/02-loop/data.yaml \
		--template-path examples/02-loop \
		--output-path examples/02-loop/dist \
		--verbose

run-example-3:
	@go run main.go build --values examples/03-conditional/data.yaml \
		--template-path examples/03-conditional \
		--output-path examples/03-conditional/dist \
		--verbose

clear:
	@rm -rf examples/*/dist