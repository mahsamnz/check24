.PHONY: test
test:
	go test -v ./...

INPUT_FILE_PATH ?= input.json
PROVIDER ?= ACME

.PHONY: run
run:
	@echo "Usage: make run INPUT_FILE_PATH=<path-to-json-file> PROVIDER=<provider-name>"
	@echo "Example: make run INPUT_FILE_PATH=./data/request.json PROVIDER=ACME"
	go run cmd/main.go $(INPUT_FILE_PATH) $(PROVIDER)