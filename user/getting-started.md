
# Getting Started

Right now, 2017Q2 the easiest way to get started is to use local driver development environment: manually clone and build a container for a language driver with an SDK, then run it through Docker in order to get source -> UAST.

i.e for a [Python driver](https://github.com/bblfsh/python-driver) :

## Prerequisites
 - Docker
 - Go, for SKD build
 - (macOS) `coreutils` and `gettext`, for SDK to work

## Build

**Build SDK**
```
go get -u github.com/bblfsh/sdk/...
```

**Build driver + container**
```
# Ignore the 'no buildable Go source files' message
go get -u github.com/bblfsh/python-driver

cd $GOPATH/src/github.com/bblfsh/python-driver
go get -v -t ./...
bblfsh-sdk prepare-build
make build
```

In future, this will be replaced by `docker fetch` for the pre-build language driver containers

## Run
```
# Get the hash of the docker container you just built
docker images | grep python-driver

# Replace the hash (the first 7 chars of the latest commit in the driver project)
docker run -it --rm -v $GOPATH/src/github.com/bblfsh/python-driver:/code bblfsh/python-driver:dev-<commit[:7]>/opt/driver/bin/driver parse-uast --format=prettyjson /code/tests/sources/comprehension.py

# Example with hash included that gets the UAST for the given `.py` file:
docker run -it --rm -v $GOPATH/src/github.com/bblfsh/python-driver:/code bblfsh/python-driver:dev-955fc5d /opt/driver/bin/driver parse-uast --format=prettyjson /code/tests/sources/comprehension.py

# To see further commands or help, append -h:
docker run -it --rm -v $GOPATH/src/github.com/bblfsh/python-driver:/code bblfsh/python-driver:dev-955fc5d /opt/driver/bin/driver -h
```

For further details on using a driver please refer [Driver protocol](../driver/protocol.md#example)

## FAQ

**I am getting GOPATH errors or the binary isn't found, what do I do?**

Bash shell: add this to your ~/.bashrc 
```
export GOPATH=$HOME/go
export PATH=$PATH:$GOROOT/bin:$GOPATH/bin
```

Fish shell: add this to your ~/.config/fish/config.fish 
```
set -gx GOPATH $HOME/go
set -U fish_user_paths $fish_user_paths $GOPATH/bin
```

