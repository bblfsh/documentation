
# Getting Started

Right now, 2017Q2 the easiest way to get started is to use local driver development environment: manually clone and build a container for a language driver with an SDK, then run it through Docker in order to get source -> UAST.

I.e for a [Java driver](https://github.com/bblfsh/java-driver) :

## Pre-requests
 - Docker
 - Go, for SKD build
 - (macOS) `coreutils` and `gettext`, for SDK to work


## Build
```
# build SDK
go get -u github.com/bblfsh/sdk/...

# build a driver + container
git clone https://github.com/bblfsh/java-driver.git
cd java-driver
go get -v -t ./...
bblfsh-sdk prepare-build
make build
```

In future, this will be replaced by `docker fetch` for the pre-build language driver containers

## Run
Get latest commit hash and run

```
cat 'native/src/main/java/bblfsh/Request.java' | docker run -it bblfsh/java-driver:dev-<commit[:7]> parse-native
```

to get UAST for the given `.java` file.

For further details on using a driver please refer [Driver protocol](../driver/protocol.md#example)
