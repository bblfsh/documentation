# Getting Started

## Using the online web client

The easiest way to get started with Babelfish is to try the online [playground](http://play.bblf.sh/) where you can write or paste your code and run the parser to see the generated UAST.

## Installing bblfshd locally

After playing with the web client, you will probably want to get Babelfish running locally. The first thing to do for that is to setup and run the [`bblfshd`](https://github.com/bblfsh/bblfshd) command. Once the server is running, you can connect to it using any of the available [clients](clients.md).

### Prerequisites

* Linux: [Docker](https://docs.docker.com/install/)
* macOS: [Docker for Mac](https://docs.docker.com/docker-for-mac/install/)
* Windows: [Docker for Windows](https://docs.docker.com/docker-for-windows/install/)

### Running with Docker \(recommended\)

The easiest way to run the _bblfshd_ daemon is using Docker. With respect to privileges of `bblfshd` daemon, there are two ways of executing it:
- As `rootless` container. You will need to get the file [`bblfshd-seccomp.json`](https://raw.githubusercontent.com/bblfsh/bblfshd/master/bblfshd-seccomp.json). Further requirements to run in rootless mode are available at [bblfshd#rootless.md](https://github.com/bblfsh/bblfshd/blob/master/rootless.md). Examples below use this mode.
- In privileged mode (we advise against it). Check out instructions [here](https://github.com/bblfsh/bblfshd/blob/master/README.md#privileged-mode).

With respect to storage, you can run it in one of these modes:
- _Embedded drivers mode_, meaning that recommended drivers will be included in the Docker image. Installing more drivers is similar to stateless mode. The Docker image will weight more of course, and there may be language drivers included in it you will not use.
- _Stateless mode_, meaning `bblfshd` is fethed without any driver included, and the drivers you install in it will be wiped out once you remove the container.
- _Stateful mode_, using a Docker volume to store part of the container internal filesystem and thus add persistence. `bblfhsd` is fetched without drivers, as in _stateless mode_, but all drivers you install will persist in the volume you use, and could be plugged in another `bblfshd` instance or even the same if you stop `bblfshd` and start it again.

#### Embedded drivers

For embedded drivers mode, run the following command:

```bash
$ docker run -d --name bblfshd \
  -p 9432:9432 \
  -v /proc:/newproc \
  --security-opt seccomp=./bblfshd-seccomp.json \
  bblfsh/bblfshd:latest-drivers
```

#### Stateless mode

For stateless mode, run the following command:

```bash
$ docker run -d --name bblfshd \
  -p 9432:9432 \
  -v /proc:/newproc \
  --security-opt seccomp=./bblfshd-seccomp.json \
  bblfsh/bblfshd:latest
```

#### Stateful mode

This is an alternative to stateless mode. Depending on your operating system, you can use a Docker volume or a normal local filesystem directory.

##### Docker volume \(Linux and macOS\)

First, create a volume with the command:

```bash
$ docker volume create bblfshd-cache
```

Then you can run the daemon with this command:

```bash
$ docker run -d --name bblfshd \
  -p 9432:9432 \
  -v /proc:/newproc \
  -v bblfshd-cache:/var/lib/bblfshd \
  --security-opt seccomp=./bblfshd-seccomp.json \
  bblfsh/bblfshd:latest
```

##### Volume mapped to a local directory \(Linux only\)

In this case, just specify the local directory in the `-v` parameter when running the daemon:

```bash
$ docker run -d --name bblfshd \
  -p 9432:9432 \
  -v /proc:/newproc \
  -v /var/lib/bblfshd:/var/lib/bblfshd bblfsh/bblfshd:latest \
  --security-opt seccomp=./bblfshd-seccomp.json \
  bblfsh/bblfshd:latest
```

### Notes on command line parameters to "docker run"

- The flag `-d` makes the container run as a daemon. You can omit it and you will see the log outputs.
- `--name bblfshd` gives a name to the container as `bblfshd`.
- Exposing the port \(`9432`\) with `-p 9432:9432` makes it easier to connect to the gRPC server from outside the container.
- `-v /proc:/newproc` is mandatory to run in rootless mode, due to a [bug in `libcontainer`](https://github.com/opencontainers/runc/issues/1658), the container technology used in bblfshd.
- `--security-opt seccomp=./bblfshd-seccomp.json` allows the default syscalls Docker permits plus `mount`, `unshare`, `pivot_root`, `keyctl`, `umount2` and `sethostname`, needed to spawn containers inside `bblfshd`. More on that matter can be consulted in [Docker docs](https://docs.docker.com/engine/security/seccomp/).

If you are behind an HTTP or HTTPS proxy server, for example in corporate settings, you will need to add the `HTTP_PROXY`, `HTTPS_PROXY`, and `NO_PROXY` environment variables in the docker run command to configure HTTP or HTTPS proxy behavior.

```bash
$ docker run -d --name bblfshd \
  -p 9432:9432 \
  -v /proc:/newproc \
  -v bblfshd-cache:/var/lib/bblfshd \
  --security-opt seccomp=./bblfshd-seccomp.json \
  -e HTTP_PROXY="http://proxy.example.com:80/" \
  bblfsh/bblfshd:latest
```

If your system uses SELinux \(like Fedora, Red Hat or CentOS among other Linux distributions\) you'll need to install a policy module. You can find instructions about it in the [bblfshd README](https://github.com/bblfsh/bblfshd#selinux).

### Testing the daemon

If everything worked, `docker logs bblfshd` should output something like this:

```text
[2019-05-01T17:58:29Z]  INFO bblfshd version: v2.12.1 (build: 2019-04-10T16:05:44+0000)
[2019-05-01T17:58:29Z]  INFO initializing runtime at /var/lib/bblfshd
[2019-05-01T17:58:29Z]  INFO server listening in 0.0.0.0:9432 (tcp)
[2019-05-01T17:58:29Z]  INFO control server listening in /var/run/bblfshctl.sock (unix)
```

#### Driver management

Now we need to install the driver images into the daemon. You can install the official images with the following command:

```bash
$ docker exec -it bblfshd bblfshctl driver install --recommended
```

If you want to install a driver for a single language, you can do so by specifying the driver URI to `bblfshctl`:

```bash
docker exec -it bblfshd bblfshctl driver install docker://bblfsh/python-driver:latest
```

You can check the installed versions executing:

```bash
$ docker exec -it bblfshd bblfshctl driver list
```

```text
+------------+------------------------------------------+---------+--------+----------+--------+-------------+----------------------+
|  LANGUAGE  |                  IMAGE                   | VERSION | STATUS | CREATED  |   OS   |     GO      |        NATIVE        |
+------------+------------------------------------------+---------+--------+----------+--------+-------------+----------------------+
| python     | docker://bblfsh/python-driver:latest     | v2.3.0  | beta   | 8 months | alpine | 1.10-alpine | python:3.6-alpine    |
| java       | docker://bblfsh/java-driver:latest       | v2.2.0  | beta   | 8 months | alpine | 1.10-alpine | openjdk:8-jre-alpine |
+------------+------------------------------------------+---------+--------+----------+--------+-------------+----------------------+
```

You can remove a driver (e.g. `python` one) with:

```bash
$ docker exec -it bblfshd bblfshctl driver remove python
```

or all drivers with:

```bash
$ docker exec -it bblfshd bblfshctl driver remove --all
```

##### Testing the drivers

To test the driver you can execute a parse request to the server with the `bblfshctl parse` command, and an example contained in the Docker image:

```bash
$ docker exec -it bblfshd bblfshctl parse /opt/bblfsh/etc/examples/python.py
```

Use this only for testing the installation; if you want to do any real parsing with files in your local filesystem you should use one of the [clients](clients.md) which would also allow you to run XPath queries over the results.

### Running standalone

A standalone distribution of `bblfshd` and `bblfshctl` can be found at the GitHub [release page](https://github.com/bblfsh/bblfshd/releases). This will contain a single binary that can be run anywhere but it depends on the `ostree` library as explained in the [readme](https://github.com/bblfsh/bblfshd#dependencies).

_bblfshd_ is only provided for Linux distributions, since it relies on Linux containers to run language drivers. And _bblfshctl_ can be found for Windows, macOS and Linux.

## Using bblfshctl

The _bblfshd_ daemon comes with a command line tool called _bblfshctl_, which can be used to monitor and manage the daemon.

If you are using the Docker image then the command line tool is provided with the `bblfsh/bblfshd` image and can be used with a `docker exec` command.

```bash
$ docker exec -it bblfshd bblfshctl --help
```

```text
Usage:
  bblfshctl [OPTIONS] <command>

Help Options:
  -h, --help  Show this help message

Available commands:
  driver     Manage drivers: install, remove and list
  instances  List the driver instances running on the daemon
  parse      Parse a file and prints the UAST or AST
  status     List all the pools of driver instances running
```

Note that the `--help` command works at any sublevel:

```bash
docker exec -it bblfshd bblfshctl driver --help
```

### Driver management

The _bblfshd's_ drivers can be installed, updated and removed with the `driver` command and its subcommands. Remember to prefix all these commands with `docker exec -it bblfshd` if you are running bblfshd from a Docker image.

Installing all official drivers:

```bash
$ bblfshctl driver install --recommended
```

Replacing an installed driver with a specific version:

```bash
$ bblfshctl driver install python bblfsh/python-driver:v2.3.0 --update
```

Listing all available drivers:

```bash
$ bblfshctl driver list
```

```text
+------------+------------------------------------------+---------+--------+----------+--------+-------------+----------------------+
|  LANGUAGE  |                  IMAGE                   | VERSION | STATUS | CREATED  |   OS   |     GO      |        NATIVE        |
+------------+------------------------------------------+---------+--------+----------+--------+-------------+----------------------+
| python     | docker://bblfsh/python-driver:latest     | v2.3.0  | beta   | 8 months | alpine | 1.10-alpine | python:3.6-alpine    |
| java       | docker://bblfsh/java-driver:latest       | v2.2.0  | beta   | 8 months | alpine | 1.10-alpine | openjdk:8-jre-alpine |
+------------+------------------------------------------+---------+--------+----------+--------+-------------+----------------------+
```
