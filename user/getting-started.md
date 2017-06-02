
# Getting Started

First thing you need to use Babelfish is running the Babelfish Server.

## Prerequisites

### Linux

- [Docker](https://www.docker.com/community-edition) (optional)

### macOS

- [Docker for Mac](https://www.docker.com/docker-mac)

### Windows

- [Docker for Windows](https://www.docker.com/docker-windows)

## Running with Docker (recommended)

The easiest way to run the server is using Docker. You can start it with the
following command:

```bash
$ docker run --privileged -p 9432:9432 --name bblfsh bblfsh/server
```

If everything worked, it should output something like this:

```
time="2017-06-01T09:12:22Z" level=debug msg="binding to 0.0.0.0:9432" 
time="2017-06-01T09:12:22Z" level=debug msg="initializing runtime at /tmp/bblfsh-runtime" 
time="2017-06-01T09:12:22Z" level=debug msg="starting server" 
time="2017-06-01T09:12:22Z" level=debug msg="registering gRPC service" 
time="2017-06-01T09:12:22Z" level=info msg="starting gRPC server" 
```

The only mandatory flag is [`--privileged`](https://docs.docker.com/engine/reference/run/#runtime-privilege-and-linux-capabilities).
Babelfish Server uses containers itself to run language drivers, which currently
requires it to run in privileged mode.

Exposing the Babelfish port (`9432`) with `-p 9432:9432:` makes it easier to
connect to it from outside the container.

Now you can test that it works by submitting a file for parsing:
 
```
$ echo "import foo" > sample.py
$ docker run -v $(pwd):/work --link bblfsh bblfsh/server bblfsh client --address=bblfsh:9432 /work/sample.py
```

First request might timeout, since the server has to fetch the required driver
before responding. If it does, just retry.

### Running standalone

> **[warning] Server only runs on Linux!**
>
> Babelfish Server relies on Linux containers to run language drivers. Windows
> and macOS users are advised to [use Docker](#running-with-docker-recommended).

Running Babelfish standalone requires getting the `bblfsh` binary. Currently
this requires a working setup of [Go](https://golang.org/). You can get it with
the following command:

```bash
$ go get -u github.com/bblfsh/server/...
```

This will install the `bblfsh` binary into `$GOPATH/bin`.

Run the server:

```bash
$ sudo bblfsh server
```

Note that running as root user is currently a requirement of the server.

The client can be run on any OS:

```
$ echo "import foo" > sample.py
$ bblfsh client sample.py
```

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

