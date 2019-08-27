.PHONY: gitbook

gitbook:
	npm install -g gitbook-cli

node_modules: gitbook
	gitbook install

build: node_modules
	gitbook build

serve: node_modules
	gitbook serve

roles:
	go run _tools/roles/main.go > uast/roles.md

languages:
	go run _tools/languages/main.go languages.md languages.json

update: languages
	go run _tools/ci-updater/main.go

types:
	GO111MODULE=on go run _tools/types/main.go > uast/types.md

clean:
	rm -rf node_modules

all: build
