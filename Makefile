.DEFAULT_GOAL := run

# enables things like `make run -- -l 3 --min-word-length 5`
ifeq (run, $(firstword $(MAKECMDGOALS)))
  runargs := $(wordlist 2, $(words $(MAKECMDGOALS)), $(MAKECMDGOALS))
  $(eval $(runargs):;@true)
endif

.PHONY: run
run:
	go run main.go $(runargs)

build:
	go build main.go -o typer

install:
	go install
