

GITBOOK_CMD := $(shell command -v gitbook 2> /dev/null)
GITBOOK_BUILD := $(GITBOOK_CMD) build
GITBOOK_SERVE := $(GITBOOK_CMD) serve
GITBOOK_INSTALL := $(GITBOOK_CMD) install

node_modules:
ifndef GITBOOK_CMD
	$(error "Please, install gitbook-cli: https://toolchain.gitbook.com/setup.html")
endif
	$(GITBOOK_INSTALL)

build: node_modules
	$(GITBOOK_BUILD)

serve: node_modules
	$(GITBOOK_SERVE)

roles:
	go run _tools/roles/main.go > uast/roles.md

languages:
	go run _tools/languages/main.go languages.md languages.json

types:
	GO111MODULE=on go run _tools/types/main.go > uast/types.md

clean:
	rm -rf node_modules

all: build
