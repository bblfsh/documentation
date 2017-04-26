
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

## Creating a driver for a new language

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

You can find the driver skeleton used here in the SDK source code:
[`etc/sekeleton`](https://github.com/bblfsh/sdk/tree/master/etc/skeleton).

## Updating SDK version

Whenever a new version of the SDK is released, drivers might need updates.
`bblfsh-sdk` can be used to perform some of this updates in managed files.
For example, if a new SDK version is released with a new version of the README
template, running `bblfsh-sdk update` will overwrite it.

```bash
$ bblfsh-sdk update
managed file "README.md" has changed, discarding changes
```

## Creating the native AST parser

The native AST parser should be in the directory `native` in the top level of the
driver repository. This directory should contain all source code and support
files (e.g. build system manifests) required to build it.

### AST parser

The core functionality to be implemented in the native side is the AST parser.
This should get the contents of a file and return an AST representation. This
AST representation should be what the AST parser outputs as is. The parser might
be from the standard library of the language, a third party library, or a custom
one (in that order of preference).

### JSON serialization

The driver should contain a JSON serializer which is able to serialize requests
and responses (including the AST) into a single-line JSON.

### Main loop

When the native parser is executed, its entry point should execute this main loop.
It should read requests from standard input, process them and write responses to
standard output. A response for a request should always be written, even if
processing fails.

If the standard input stream is closed, the program should exit with code 0.
If write to standard output fails, the program should exit with code 1.
Here we illustrate its behaviour in Python-like syntax:

```python
while True:
    line = read_line_from_stdin()
    if stdin_is_closed():
        exit(0)
    try:
        request = parse_json_request(line)
        ast = parse_ast(request)
        response = serialize_json_response(ast)
    except:
        write_to_stdout(fatal_response)
        continue
    try:
        write_to_stdout(response)
    except:
        exit(1)
```

## Creating the UAST conversion

The conversion from AST to UAST is written in Go. The main file to be edited is
`driver/normalizer/normalizer.go`.

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
