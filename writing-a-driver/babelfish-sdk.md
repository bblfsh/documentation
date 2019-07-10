# Babelfish SDK

The [Babelfish SDK](https://github.com/bblfsh/sdk) contains the tools and libraries required to create a Babelfish driver for a programming language.

## Preparations

### Dependencies

The Babelfish SDK has the following dependencies:

* [Docker](https://www.docker.com/get-docker)
* [Go](https://golang.org/dl/)

Make sure that you've correctly set your [GOROOT and GOPATH](https://golang.org/doc/code.html#Workspaces) environment variables
and that your project is _not_ inside the `GOPATH` (to enable [Go modules](https://github.com/golang/go/wiki/Modules)).

### Installing

Install the Babelfish SDK. Native drivers and the external tools can be developed in any programming language but the normalizer is developed using Go so you must get it as any other go package:

```bash
$ go get -u -v github.com/bblfsh/sdk/v3/...
```

This will fetch the SDK repository to the modules cache and will install the `bblfsh-sdk` CLI to `$GOPATH/bin/bblfsh-sdk`.

### Creating the driver's initial structure

Let's say we're creating a driver for `mylang`. The driver should be bootstrapped with `bblfsh-sdk`. This will create a new git repo with some content required by every driver.

```bash
$ bblfsh-sdk init mylang alpine
initializing driver "mylang", creating new manifest
initializing new repo "mylang-driver"
creating file "mylang-driver/driver/fixtures/fixtures_test.go"
creating file "mylang-driver/driver/main.go"
creating file "mylang-driver/driver/normalizer/annotation.go"
creating file "mylang-driver/driver/normalizer/transforms.go"
creating file "mylang-driver/test.go"
creating file "mylang-driver/.gitignore"
creating file "mylang-driver/driver/normalizer/normalizer.go"
creating file "mylang-driver/go.mod"
creating file "mylang-driver/native/README.md"
creating file "mylang-driver/native/native.sh"
creating file "mylang-driver/update.go"
creating file "mylang-driver/.travis.yml"
creating file "mylang-driver/LICENSE"
creating file "mylang-driver/driver/impl/impl.go"
creating file "mylang-driver/driver/sdk_test.go"
creating file "mylang-driver/README.md"
creating file "mylang-driver/build.go"
+ go mod tidy
+ go mod vendor
managed file "Dockerfile" has changed, overriding changes
+ git add Dockerfile go.mod go.sum
$ cd mylang-driver
```

Note that this adds a pre-commit [git hook](https://git-scm.com/book/en/v2/Customizing-Git-Git-Hooks), which will verify these files are up to date before every commit and will disallow commits if some of the managed files are changed. You can by-pass this with `git commit --no-verify`.

You can find the driver skeleton used here in the SDK source code: [`etc/skeleton`](https://github.com/bblfsh/sdk/tree/master/etc/skeleton).

### Building an "echo" driver

The skeleton provides a functional driver that echoes the source file back as a response.

To build it, make sure you have [Docker](https://www.docker.com/get-docker) installed and run:

```bash
$ go run ./build.go mylang:dev
sha256:3b199c4c35dab3535a20cf4a3c0583482c8cb796f0c67affad2755ad223a1381
```

This will execute a `build.go` script, which in turn generates a `Dockerfile` from the driver build manifest (`build.yml`)
and builds a container image with a `mylang:dev` name.

Let's start the new driver:

```bash
$ docker run -it --rm -p 9432:9432 mylang:dev
[2019-07-10T13:42:16Z]  INFO mylang-driver version: dev-unknown-dirty (build: 2019-07-10T13:35:49Z)
[2019-07-10T13:42:16Z]  INFO server listening in 0.0.0.0:9432 (tcp)
```

We can now test it with any of the [Babelfish clients](../using-babelfish/clients.md).
For example, using [Go client CLI](https://github.com/bblfsh/go-client#cli):

```bash
$ echo "test source" > test.txt
$ bblfsh-cli -l mylang test.txt
{
   Encoding: "utf8",
   content: "test source\n",
}
```

The output is the fake AST that our echo driver emits. It corresponds to the JSON request received by the native driver
via `stdin`. The [`./native/native.sh`](https://github.com/bblfsh/sdk/blob/master/etc/skeleton/native/native.sh) file
is responsible for parsing the request and writing the JSON response back to `stdout`.

This script is automatically executed by the driver server written in Go, which exposes it via gRPC protocol to `bblfshd`
and runs the [transformation DSL](./transform-dsl.md) engine for converting native AST to UAST.

It is also possible to debug the native driver by executing it directly:

```bash
$ docker run -it --rm -p 9432:9432 --entrypoint=/opt/driver/bin/native mylang:dev
> {"content":"test"}
< {"status":"ok", "ast": {"content":"test"}}
```

### Updating SDK version

Whenever a new version of the SDK is released, drivers might need updates. `bblfsh-sdk` can be used to perform some of this updates in managed files.
For example, if a new SDK version is released with a new version of the README template, running `bblfsh-sdk update` will overwrite it.

```bash
$ bblfsh-sdk update
managed file "README.md" has changed, discarding changes
```

`bblfsh-sdk` only updates driver's SDK version up to the version of the tool, so if you want to use a newer version of the SDK you'll also have to update the Go package **first** with:

```bash
go get -u github.com/bblfsh/sdk/v3/...
```

## Implementing the Driver

A driver has several sub-components. On this document and the next one we'll see how to implement all these parts to create a new driver.

The components are:

* Tests: source code, native AST, and UAST files for the UAST transformation and integration tests.
* Native parser: parses source code in the driver's target language and generates a language- or library-specific \(non-universal\) AST.
* Driver normalizer: annotates and normalizes the native AST and generates the final UAST.

### Creating the native parser

The native source code parser implementation should be in the directory `native` in the top level of the driver repository.
This directory should contain all source code and support files \(e.g. build system manifests\) required to build it.

The core functionality to be implemented on the parser side is the source code parser. This should get the source contents of a file and return an AST representation.
This AST representation should be what the native or library AST parser outputs as is. The native parser might be from the standard library of the language, a third-party library, or a custom one \(in that order of preference\).
Any language can be used to implement the parser so use whatever is easier for the specific language you're working on.

#### Native parser operation

When the driver executes the parser, the former entry point should run a main loop similar to the example below.
It should read requests from standard input \(which contain a string with the contents of the source to parse\), parse it generating the native AST and write a response to the standard output.

Requests and responses must be serialized in JSON format on a single line finished with a Unix newline character \(`\n`\) so the driver can iterate over the standard input line by line.
This means so the language selected to implement the parser should support it in the standard library or any third party module.

A response for a request should always be written, even if processing fails \(in this case an error response will be returned\).
If the standard input stream is closed or the program can write to it for any other reason, the program should exit with code 0.
If the writing to standard output fails, the program should exit with code 1 \(additionally it could try to log the error to `stderr`\).

Here we illustrate its behaviour in Python-like pseudocode:

```python
while True:
    line = read_line_from_stdin()
    if stdin_is_closed():
        exit(0)

    try:
        request = parse_json_request(line)
        ast = parse_code(request["content"])
        response = serialize_json_response(ast)
    except Exception, e:
        try:
            write_to_stdout(generate_error_response(e))
        except:
            exit(1)
        continue

    try:
        write_to_stdout(response)
    except:
        exit(1)
```

In order to avoid maintainability burdens, the parser should avoid, as much as possible, to change the AST that is generated by the language library or third party module in use.
One exception to this would be to improve the information kept in the AST to kept more information as explained in the ["From Code to AST"](../uast/code-to-ast.md) page, for example, for including comments and whitespace information.
Use your own discretion when balancing the improved information versus the development and maintainability effort.

The fields in the JSON requests and responses are detailed on the [internal protocol](internal-protocol.md) page.

_Note:_ When extracting the native AST you should try to keep as much information as possible following the guidelines explained in the ["From Code to AST"](../uast/code-to-ast.md) page.

#### Native parser unit tests

The parser should provide unit tests checking that the parsing and the AST conversion work well.
Since the parser can be written in many different languages, the location and form of the tests should try to be as idiomatic as possible to avoid future maintainability problems.

Native tests are executed during the driver build (`go run ./build.go`) and are script to use for the tests can be set
in the `build.yml`'s `native.test` section:

```bash
  test:
    run:
      # TODO: native driver tests
      - 'echo tests'
```

**TODO:** Document the `test` section better (e.g. what is the source image for the test environment, etc). For now, see `build.yml` in [existing drivers](../languages.md).

### Creating the Converter and Annotator

The conversion from AST to UAST is written in Go. Main files to be edited are `annotation.go` and `normalizer.go` in `./driver/normalizer/`. The details are explained in the [annotations section](adding-uast-annotations.md).

### Updating the driver build instructions

The `build.yml` must contain instructions to build the parser and install it into a preconfigured path:

```text
native:
  # TODO: an image used as a driver runtime
  image: 'debian:latest'
  # TODO: files copied from the source to the driver image
  static:
  - path: 'native.sh'
    dest: 'native'
  build:
    # TODO: an image to build a driver
    image: 'debian:latest'
    # TODO: build system dependencies, can not use the source
    deps:
      - 'echo dependencies'
    # TODO: build steps
    run:
      - 'echo build'
# TODO: files copied from the builder to the driver image
#    artifacts:
#      - path: '/native/native-binary'
#        dest: 'native-binary'

```

### The Dockerfile

The driver build system is based on a single multi-stage Dockerfile. The `bblfsh-sdk` tools acts as a code generator for
the main Dockerfile. Thus, you can debug the driver build by running it manually:

```bash
$ go run ./update.go
$ bblfsh-sdk release
$ docker build -t mylang:dev .
```

The first command will read the `manifest.toml` and `build.yml` to generate the main `Dockerfile`.

The second command creates a file with build-related metadata, required by the driver build stages.

And the last command executes Docker build normally. It produces the same images as `go run ./build.go` would.

### Creating the UAST tests

The integration tests will test all the process of the driver from the request to the parser to the annotated UAST generation.
They work by comparing the UAST and native output with the previously existing one, and will fail if any difference is found.

For this to work, the driver developer will need to provide files with source code examples in the `./fixtures/` directory with the native source file extension of the target language (e.g. `.js`).

Also, if you are starting with a driver template, you may need to change the [./driver/fixtures/fixtures_test.go](https://github.com/bblfsh/sdk/blob/master/etc/skeleton/driver/fixtures/fixtures_test.go.tpl#L16)
file, specifically the `Ext` field, that must match the file extension of used for the test files.

Once this is done, you need to generate the `.native`, `.uast` and `.sem.uast` files. For this, run the command from the driver top directory:

```text
go run ./test.go
```

The first time \(or every time you add a new source file and regenerate the native/uast files\) you'll need to manually and carefully check the `.uast` and `.sem.uast` files because they'll be used as reference on for the integration tests.
Once you're happy with them, you can re-run the the UAST transformation tests alone:

```bash
go test -v ./driver/...
```

This will generate the `.uast` and `.sem.uast` files from the `.native` files in the `./fixtures` directory and compare them with the previously generated files, failing and printing a diff if it detects any changes.
It won't generate the `.native` files, so you will need to run a full test pass (`go run ./test.go`) each time you update the native driver or the test source files.

It's advisable to create very small source files just testing the annotation of one language feature, or even several if the feature is complex. This will make the tests more atomic and help immensely with your sanity when trying to debug failed integration tests.

### Officially adding your driver to Babelfish

The last step if to have your language added to the [Babelfish Github's Project](https://github.com/bblfsh) so more people can use/break/debug or improve it. For this you need to request a new repository to be created which you can make by opening a PR or talking with us in the [public Babelfish Slack channel](https://sourced-community.slack.com/join/shared_invite/MTkwNTM0ODEyODIzLTE0OTYxMzc5NTMtODRhMDYyNzAyYQ). Once the driver subproject has been added you'll make a PR with your code.

