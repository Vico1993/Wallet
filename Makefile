default: build lint

.PHONY: build test lint

build:
	@ go build -a \
			 -o "./wallet" "./cmd/wallet"
	@ echo "Build done 🛠"

test:
	DEBUG=1 go test -v ./...

lint:
	@ golangci-lint run ./... -v
	@ echo "Lint done 🪛"