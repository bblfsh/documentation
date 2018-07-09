# Babelfish SDK

The [Babelfish SDK](https://gopkg.in/bblfsh/sdk.v2) contains the tools and libraries required to create a Babelfish driver for a programming language.

## Preparations

### Dependencies

The Babelfish SDK has the following dependencies:

* [Docker](https://www.docker.com/get-docker)
* [Go](https://golang.org/dl/)

Make sure that you've correctly set your [GOROOT and GOPATH](https://golang.org/doc/code.html#Workspaces) environment variables.

### Installing

Install the Babelfish SDK. Native drivers and the external tools can be developed in any programming language but the normalizer is developed using Go so you must get it as any other go package:

```bash
$ go get -u -v gopkg.in/bblfsh/sdk.v2/...
```

This will fetch the SDK repository to `$GOPATH/src/gopkg.in/bblfsh/sdk.v2` and will install the `bblfsh-sdk` CLI to `$GOPATH/bin/bblfsh-sdk`.

### Creating the driver's initial structure

Let's say we're creating a driver for `mylang`. The driver should be bootstrapped with `bblfsh-sdk`. This will create a new git repo with some content required by every driver.

```bash
$ cd $GOPATH/src/github.com/bblfsh
$ bblfsh-sdk init mylang alpine
initializing driver "mylang", creating new manifest
initializing new repo "mylang-driver"
creating file "mylang-driver/manifest.toml"
creating file "mylang-driver/.travis.yml"
creating file "mylang-driver/Dockerfile.tpl"
creating file "mylang-driver/README.md"
creating file "mylang-driver/.git/hooks/pre-commit"
creating file "mylang-driver/.gitignore"
creating file "mylang-driver/Dockerfile.build.tpl"
creating file "mylang-driver/LICENSE"
creating file "mylang-driver/Makefile"
creating file "mylang-driver/driver/main.go"
creating file "mylang-driver/driver/normalizer/annotation.go"
$ cd mylang-driver
$ git add -A
$ git commit -m 'initialize repository'
```

Note that this adds a pre-commit [git hook](https://git-scm.com/book/en/v2/Customizing-Git-Git-Hooks), which will verify these files are up to date before every commit and will disallow commits if some of the managed files are changed. You can by-pass this with `git commit --no-verify`.

You can find the driver skeleton used here in the SDK source code: [`etc/skeleton`](https://github.com/bblfsh/sdk/tree/master/etc/skeleton).

### Updating SDK version

Whenever a new version of the SDK is released, drivers might need updates. `bblfsh-sdk` can be used to perform some of this updates in managed files. For example, if a new SDK version is released with a new version of the README template, running `bblfsh-sdk update` will overwrite it.

```bash
$ bblfsh-sdk update
managed file "README.md" has changed, discarding changes
```

`bblfsh-sdk` doesn't update the SDK itself so if you want to use a new version of the SDK you'll also have to update the Go package **first** with:

```bash
go get -u gopkg.in/bblfsh/sdk.v2/...
```

If the update gives you any problem you can try to delete the `$GOPATH/src/gopkg.in/bblfsh/sdk.v2` manually and run the go get command again to get a fresh copy.

## Implementing the Driver

A driver has several subcomponents. On this document and the next one we'll see how to implement all these parts to create a new driver.

The components are:

* Tests: source code, native AST, and UAST files for the integration tests.
* Native parser: parses source code in the driver's target language and generates a language or library specific \(non-universal\) AST.
* Driver normalizer: annotates and generates the final UAST.

### Creating the native parser

The native source code parser implementation should be in the directory `native` in the top level of the driver repository. This directory should contain all source code and support files \(e.g. build system manifests\) required to build it.

The core functionality to be implemented on the parser side is the source code parser. This should get the source contents of a file and return an AST representation. This AST representation should be what the native or library AST parser outputs as is. The native parser might be from the standard library of the language, a third-party library, or a custom one \(in that order of preference\). Any language can be used to implement the parser so use whatever is easier for the specific language you're working on.

#### Native parser operation

When the driver executes the parser, the former entry point should run a main loop similar to the example below. It should read requests from standard input \(which contain a string with the contents of the source to parse\), parse it generating the native AST and write a response to the standard output.

Requests and responses must be serialized in JSON format on a single line finished with a Unix newline character \('\n'\) so the driver can iterate over the standard input line by line. This means so the language selected to implement the parser should support it in the standard library or any third party module.

A response for a request should always be written, even if processing fails \(in this case an error response will be returned\). If the standard input stream is closed or the program can write to it for any other reason, the program should exit with code 0. If the writing to standard output fails, the program should exit with code 1 \(additionally it could try to log the error to stderr\).

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

In order to avoid maintainability burdens, the parser should avoid, as much as possible, to change the AST that is generated by the language library or third party module in use. One exception to this would be to improve the information kept in the AST to kept more information as explained in the ["From Code to AST"](../uast/code-to-ast.md) page, for example, for including comments and whitespace information. Use your own discretion when balancing the improved information versus the development and maintainability effort.

The fields in the JSON requests and responses are detailed on the [internal protocol](internal-protocol.md) page.

_Note:_ When extracting the native AST you should try to keep as much information as possible following the guidelines explained in the ["From Code to AST"](../uast/code-to-ast.md) page.

#### Native parser unit tests

The parser should provide unit tests checking that the parsing and the AST conversion work well. Since the parser can be written in many different languages, the location and form of the tests should try to be as idiomatic as possible to avoid future maintainability problems.

Once you've created your unit tests, you must modify the Makefile in the root of the driver filesystem adding the commands to run them under the Makefile target `test-parser-internal`. For example, for a Python driver using the Python default test framework it could look like:

```text
test-parser-internal:
    cd native/python_package/test; \
    python3 -m unittest discover
```

Once you've done it, you can run your tests with `make test-parser-internal`. Please note that these tests aren't run inside a Docker container so you'll need to have any needed dependency installed on your environment to run the tests.

### Creating the Converter and Annotator

The conversion from AST to UAST is written in Go. The main file to be edited is `driver/normalizer/annotation.go`. The details are explained in the [annotations section](adding-uast-annotations.md).

### Updating the Makefile with the build instructions

The `Makefile` must contain a target with instructions to build the parser and install it into a preconfigured path. The `bblfsh` tool has already added that target with this code:

```text
build-native-internal:
    cd native; \
    echo "not implemented"
    echo -e "#!/bin/bash\necho 'not implemented'" > $(BUILD_PATH)/bin/native
    chmod +x $(BUILD_PATH)/bin/native
```

You need to remove the placeholder `echo` commands and add there any commands needed to compile \(for non-interpreted languages\) and generate a binary or script with the name `native`, copy it to `$(BUILD\_PATH)` and assign execution permissions.

For example, for a typical C program the target could be:

```text
build-native-internal:
    cd native; \
    gcc parser.c -o $(BUILD_PATH)/bin/native
    chmod +x $(BUILD_PATH)/bin/native
```

This makefile target will usually run inside a Docker container \(see the next section\), but you can test it running `make build-native-internal` and checking the resulting file on the `build` directory.

### Configuring the build Dockerfile

The building of the driver is done in a Docker container so you must adapt the Docker template `Dockerfile.build.tpl` that the `bblfsh` command created for you to install all the dependencies needed to build the `native` command detailed in the previous step.

The templated dockerfile looks like this \(without comments\):

```text
FROM alpine:3.6
RUN mkdir -p /opt/driver/src && \
    adduser ${BUILD_USER} -u ${BUILD_UID} -D -h /opt/driver/src
RUN apk add --no-cache make git curl ca-certificates
WORKDIR /opt/driver/src
```

As the comments in the file say, Alpine is the preferred distribution for space reasons but you can use Debian slim or other slim versions if for any reason Alpine don't work for your driver.

You should add any packages you need to the `apk add` line. Avoid adding more lines with more `RUN apk add` commands instead because that would create more intermediate layers \(taking more space\).

You can also install dependencies using the language package manager instead \(pip, cargo, dub, npm, etc. if you prefer, after installing it with apk. Another option is downloading and compiling the dependencies. If you go that way you'll probably need the `build-base` or `alpine-sdk` packages.

You can see a list and search for Alpine packages [here](https://pkgs.alpinelinux.org/packages).

You don't need to run any command to build the `native` file; the SDK will execute `make build-native-internal` target once the Docker has been created.

To trigger the building and test your Dockerfile and building process you can issue the commands:

```bash
$ bblfsh-sdk prepare-build
$ make build
```

_Note_: if you are on macOS and you see an error `envsubst: command not found` try installing the `gettext` package with homebrew or similar.

```bash
$ brew install gettext
$ brew link --force gettext
```

You'll see a sequence of `docker build` and `docker run` commands execute and any building error will show on the standard output.

### Configuring the run Dockerfile

The second dockerfile template that you have to edit is the one that the Babelfish server will use to run new instances of drivers. Like the build one it will also be generated by the `bblfsh-sdk` command but with the filename `Dockerfile.tpl` and this initial content without comments:

```text
FROM alpine:3.6
ENTRYPOINT ["/opt/driver/bin/driver"]
```

As with the build container, the first step is to add any packages needed to **run** the `native` binary produced by the build process, if any is needed like dynamically loading libraries for compiled languages \(not included in Alpine base\) or runtimes for interpreted or compiled-to-bytecode languages.

Then you will need to copy the `build` directory into `/opt/driver/bin`, so that both the \(automatically build by the SDK\) `driver` and `native` binaries will rest in that directory.

The final `CMD` command should remain at the end of the file and you should not change it since it'll execute the driver.

### Creating the integration tests

The integration tests will test all the process of the driver from the request to the parser to the annotated UAST generation. They work by comparing the UAST and native output with the previously existing one, and will fail if any difference is found.

For this to work, the driver developer will need to provide files with source code examples in the `tests/` directory with the `.source` file extension. It's recommended that you keep the original extension before the `.source` so the command below can autodetect the language.

Once this is done, you need to generate the `.native` and `.uast` files, but first you should have either a [bblfshd server with your driver added](../using-babelfish/advanced-usage.md#adding-drivers-from-the-local-docker-daemon) or a [single driver listening on the 9432 port](../using-babelfish/advanced-usage.md#running-a-driver-without-the-server), the latest option being probably the easiest.

Then, run the command from the driver top directory:

```text
bblfsh-sdk-tool fixtures fixtures/*.source
```

If the command fails with the error `could not detect language` it could be because the source files doesn't have the right extension \(like when the only extension is `.source`\). If you positively know that there is a driver installed for that language, add the `--language=mylang` to the command above to skip the autodetection.

The first time \(or every time you add a new `.source` file and regenerate the native/uast files\) you'll need to manually check carefully the `.uast` files because they'll be used as reference on for the integration tests. Once you're happy with them, you can run the integration tests with:

```bash
make integration-test
```

This will generate the `.native` and `.uast` files from the driver in the current directory and compare them with the previously generated files, failing and printing a diff if it detects any changes.

It's advisable to create very small `.source` files just testing the annotation of one language feature, or even several if the feature is complex. This will make the tests more atomic and help immensely with your sanity when trying to debug failed integration tests.

### Officially adding your driver to Babelfish

The last step if to have your language added to the [Babelfish Github's Project](https://github.com/bblfsh) so more people can use/break/debug or improve it. For this you need to request a new repository to be created which you can make by opening a PR or talking with us in the [public Babelfish Slack channel](https://sourced-community.slack.com/join/shared_invite/MTkwNTM0ODEyODIzLTE0OTYxMzc5NTMtODRhMDYyNzAyYQ). Once the driver subproject has been added you'll make a PR with your code.

