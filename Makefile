RUN_ARGS := $(wordlist 2,$(words $(MAKECMDGOALS)),$(MAKECMDGOALS))

# If the first argument is "run"...
ifeq (run,$(firstword $(MAKECMDGOALS)))
  # ...and turn them into do-nothing targets
  $(eval $(RUN_ARGS):;@:)
endif

# If the first argument is "debug"...
ifeq (debug,$(firstword $(MAKECMDGOALS)))
  # ...and turn them into do-nothing targets
  $(eval $(RUN_ARGS):;@:)
endif


default: build lint

.PHONY: build debug_run run test lint

build:
	@ go build -a \
			 -o "./wallet" "./cmd/wallet"
	@ echo "Build done 🛠"

debug:
	DEBUG=1 go run ./... $(RUN_ARGS)

run:
	DEBUG=0 go run ./... $(RUN_ARGS)

test:
	go test -v ./...

lint:
	@ golangci-lint run ./... -v
	@ echo "Lint done 🪛"