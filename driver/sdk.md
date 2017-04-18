
# Babelfish SDK

The [Babelfish SDK](https://github.com/bblfsh/sdk/) contains tools and libraries
required to create a Babelfish driver.

## Prerequisites

Babelfish has the following prerequisites:

* [Docker](https://www.docker.com/get-docker)
* [Go](https://golang.org/dl/)

## Installing

Install the Babelfish SDK:

```bash
$ go get github.com/bblfsh/sdk/...
```

This will fetch the sdk repository to `$GOPATH/src/github.com/bblfsh/sdk` and
will install the `bblfsh-sdk` CLI to `$GOPATH/bin/bblfsh-sdk`.

## Bootstrapping a new driver

Let's say we're creating a driver for `mylang`. First step is initializing a git
repository for the driver:

```bash
$ cd $GOPATH/src/github.com/bblfsh
$ git init mylang-driver
$ cd mylang-driver
$ touch README.md
$ git add README.md
$ git commit -m 'initialize repository'
```

Now the driver should be bootstrapped with `bblfsh-sdk`. This will create some
directories and files required by every driver. They will be overwritten if they
exist, like the README.md file in the example below.

```bash
$ bblfsh-sdk init mylang alpine
initializing driver "mylang", creating new manifest
creating file "manifest.toml"
creating file "Makefile"
creating file "driver/main.go"
creating file "driver/normalizer/normalizer.go"
creating file ".git/hooks/pre-commit"
creating file ".gitignore"
creating file ".travis.yml"
creating file "Dockerfile.build.tpl"
creating file "driver/normalizer/normalizer_test.go"
creating file "Dockerfile.tpl"
creating file "LICENSE"
managed file "README.md" has changed, discarding changes
```

Note that this adds a pre-commit [git hook](https://git-scm.com/book/en/v2/Customizing-Git-Git-Hooks),
which will verify these files are up to date before every commit. You can
by-pass this with `git commit --no-verify`.

## Creating the native AST parser

The native AST parser should be in the directory `native` in the top level of the
driver repository.

TODO

### AST parser

TODO

### Main loop

TODO

### JSON serialization

TODO

## Creating the UAST conversion

TODO

## Building

```bash
$ bblfsh-sdk prepare-build
$ make build
```

TODO

## Testing

```bash
$ bblfsh-sdk prepare-build
$ make integration-test
```

TODO