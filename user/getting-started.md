# Getting Started

The first thing you need to use Babelfish is to setup and run the Babelfish
Server. Once the server is running, you can use the Babelfish Tools to get some
information from your code.

## Babelfish server

### Prerequisites

- Linux: [Docker](https://www.docker.com/community-edition) (optional)
- macOS: [Docker for Mac](https://www.docker.com/docker-mac)
- Windows: [Docker for Windows](https://www.docker.com/docker-windows)

### Running with Docker (recommended)

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

As you've probably already seen,
the logging level is set to `debug` by default,
which is a good default for current project status,
but it may be cumbersome.
If the output is too verbose,
it can be adjusted with the `log-level` parameter:

```bash
$ docker run --privileged -p 9432:9432 --name bblfsh bblfsh/server bblfsh server --log-level info
```

Now you should only see log entries of `info` level and above:

```
time="2017-06-01T09:12:22Z" level=info msg="starting gRPC server"
```

### Running standalone

> **[warning] Standalone server only runs on Linux!**
>
> Babelfish Server relies on Linux containers to run language drivers. Windows
> and macOS users are advised to [use Docker](#running-with-docker-recommended).

[Download the bblfsh binary](https://github.com/bblfsh/server/releases).

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

### Overriding driver images

In case you need that the Babelfish server to run different driver images than the default ones, you can configure what images it should use through the environment variable `BBLFSH_DRIVER_IMAGES`. The command `bblfsh server ` looks for this variable which must look like:

    BBLFSH_DRIVER_IMAGES="language=docker-registry:namespace/repository:tag;language2=docker-registry:namespace/repository:tag"

So to get a specific version of an image from
[Dockerhub](https://hub.docker.com/u/bblfsh/) the line would be:

```bash
$ BBLFSH_DRIVER_IMAGES="python=docker://bblfsh/python:v0.4.2" docker run \
  -e BBLFSH_DRIVER_IMAGES -v /var/run/docker.sock:/var/run/docker.sock \
  --privileged -p 9432:9432 --name bblfsh bblfsh/server
```

Instead, if you want the server to retrieve specific driver images from a local Docker
daemon (e.g. when testing a driver that you're developing) you could do something similar to:

```bash
$
BBLFSH_DRIVER_IMAGES="python=docker-daemon:bblfsh/python:dev-96b24d3;java=docker-daemon:bblfsh/java-driver:latest"
docker run -e BBLFSH_DRIVER_IMAGES -v
/var/run/docker.sock:/var/run/docker.sock --privileged -p 9432:9432 --name bblfsh bblfsh/server
time="2017-07-12T14:11:13Z" level=debug msg="binding to 0.0.0.0:9432"
time="2017-07-12T14:11:13Z" level=debug msg="initializing runtime at /tmp/bblfsh-runtime"
time="2017-07-12T14:11:13Z" level=debug msg="Overriding image for "python: docker-daemon:bblfsh/python:dev-96b24d3"
time="2017-07-12T14:11:13Z" level=debug msg="Overriding image for java: docker-daemon:bblfsh/java-driver:latest""
time="2017-07-12T14:11:13Z" level=debug msg="starting server"
time="2017-07-12T14:11:13Z" level=debug msg="registering gRPC service"
time="2017-07-12T14:11:13Z" level=info msg="starting gRPC server"
```

Notice how in this case we need to share the host Docker server's Unix socket with
the container (`-v /var/run/docker.sock:/var/run/docker.sock`) so it can access it
to retrieve the local images.

Or if you prefer running it in standalone mode:

```
$ BBLFSH_DRIVER_IMAGES="python=docker-daemon:bblfsh/python:dev-96b24d3;java=docker-daemon:bblfsh/java-driver:latest" bblfsh server
DEBU[0000] binding to 0.0.0.0:9432                      
DEBU[0000] initializing runtime at /tmp/bblfsh-runtime  
DEBU[0000] Overriding image for python: docker-daemon:bblfsh/python:dev-96b24d3
DEBU[0000] Overriding image for java: docker-daemon:bblfsh/java-driver:latest
DEBU[0000] starting server                              
DEBU[0000] registering gRPC service                     
INFO[0000] starting gRPC server
```


### Setting maximum message size

If a customized [gRPC](https://grpc.io) message size is needed, you can use the command flag option `--max-message-size`. By default  [gRPC](https://grpc.io) uses 4MB as the upper limit, but you can override it:

    docker run --privileged -p 9432:9432 --name bblfsh bblfsh/server --max-message-size=100

or running the Babelfish Server in local:

    sudo bblfsh server --max-message-size=100

The number given to the `--max-message-size` option represents the size in MB, and it defines the limit for both directions: send and receive.


## Babelfish Tools

Babelfish Tols provide some language analysis tools on top of Babelfish. You can
use them for various purposes:

- Check that the server is working properly
- Get some data from the source code
- See how a language analysis tool is implemented, as a basis for your own
  tools.


### Setup

Running Babelfish Tools standalone requires getting the `bblfsh-tools`
binary. Currently this requires a working setup of
[Go](https://golang.org/doc/install). You can get it with the following command:

```bash
$ go get -u github.com/bblfsh/tools/...
```

### Usage

Babelfish Tools provides a set of tools built on top of Babelfish, to
see which tools are supported, run:

`bblfsh-tools --help`

There's an special tool, the `dummy` tool, which should let you know if the
connection with the server succeeded:

`bblfsh-tools dummy path/to/source/code`

If the server is not in the default location, use the `address` parameter:

`bblfsh-tools dummy --address location:port path/to/source/code`

Once the connection with the server is working fine, you can use any other
available tool in a similar way.

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
