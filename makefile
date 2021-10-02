VERSION := $(shell awk '/VERSION/{ gsub(/"/,"",$$4); print $$4 }' main.go)


COLOR_RED := $(shell echo -e "\033[0;31m")
COLOR_YELLOW := $(shell echo -e "\033[0;33m")
COLOR_END := $(shell echo -e "\033[0m")

build: clean _version test
	@echo -e "$(COLOR_YELLOW)Building the project $(COLOR_END)"
	mkdir build || true
	go build -o build/go-pass-keeper main.go

test:
	@echo -e "$(COLOR_YELLOW)Building the project $(COLOR_END)"
	@#go test --json -v  ./..
	 go test -v  ./src/*

_version:
	@echo -e "go-pass-keeper version: $(VERSION)-$(RELEASE)"

help: _version
	@echo -e "$(COLOR_YELLOW)List of actions that can be executed $(COLOR_END)"
	@echo " -> build"
	@echo " -> verison"
	@echo " -> help (this action)"

clean:
	@echo -e "$(COLOR_YELLOW)Removed old files $(COLOR_END)"
	rm build/go-pass-keeper || true

