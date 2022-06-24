default: build lint

.PHONY: build debug_run run test lint

build:
	@ go build -a \
			 -o "./wallet" "./cmd/wallet"
	@ echo "Build done 🛠"

debug:
	DEBUG=1 go run ./...

run:
	DEBUG=0 go run ./...

test:
	go test -v ./...

lint:
	@ golangci-lint run . -v
	@ echo "Lint done 🪛"