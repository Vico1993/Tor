default: build lint

.PHONY: build test lint

build:
	@ go build -a \
			 -o "./tor"
	@ echo "Build done 🛠"

test:
	go test -v ./...

lint:
	@ golangci-lint run ./... -v
	@ echo "Lint done 🪛"