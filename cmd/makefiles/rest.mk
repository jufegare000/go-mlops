APP_NAME=rest-client
ENTRYPOINT=./cmd/rest-client
OUTPUT=./bin/$(APP_NAME)

.PHONY: build-rest run-rest

build-rest:
	go build -o $(OUTPUT) $(ENTRYPOINT)

run-rest: build-rest
	$(OUTPUT)