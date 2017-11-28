# Getting Started

## Using the online dashboard

The easiest way to get started with Babelfish is to try the online [dashboard](http://dashboard.bblf.sh/)
where you can write or paste your code and run the parser to see the generated UAST.

## Installing bblfshd locally

After playing with the dashboard, you will probably want to get Babelfish running locally.
The first thing to do for that is to setup and run the [`bblfshd`](https://github.com/bblfsh/bblfshd)
command. Once the server is running, you can connect to it using any of the available [clients](language-clients.md).

### Prerequisites

- Linux: [Docker](https://www.docker.com/community-edition) (optional)
- macOS: [Docker for Mac](https://www.docker.com/docker-mac)
- Windows: [Docker for Windows](https://www.docker.com/docker-windows)

### Running with Docker (recommended)

The easiest way to run the *bblfshd* is using Docker. You can start it with the
following command:

```bash
$ docker run -d --name bblfshd --privileged -p 9432:9432 bblfsh/bblfshd
```

This will run the image in a stateless mode, this meaning that any installed
drivers will be lost when you stop the container. To avoid this from happening,
add the `-v` parameter to keep the `/var/lib/bblfshd` directory in a volume:

```bash
$ docker run -d --name bblfshd --privileged -p 9432:9432 -v /var/lib/bblfshd:/var/lib/bblfshd bblfsh/bblfshd
```

On macOS, remove the parameter `-v /var/lib/bblfshd:/var/lib/bblfshd` since the
default case insensitive filesystem could conflict with the internal drivers'
case sensitive one. In that case remember that since this directory will then be
ephemeral, you'll need to reinstall the drivers after creating a new 
server container.

If everything worked, `docker logs bblfshd` should output something like this:

```
time="2017-10-10T08:59:20Z" level=info msg="bblfshd version: v2.0.0 (build: 2017-10-09T21:18:54+0000)"
time="2017-10-10T08:59:20Z" level=info msg="initializing runtime at /var/lib/bblfshd"
time="2017-10-10T08:59:20Z" level=info msg="server listening in 0.0.0.0:9432 (tcp)"
time="2017-10-10T08:59:20Z" level=info msg="control server listening in /var/run/bblfshctl.sock (unix)"

```

The only mandatory flag is [`--privileged`](https://docs.docker.com/engine/reference/run/#runtime-privilege-and-linux-capabilities).
*bblfshd* uses containers itself to run language drivers, which currently
requires it to run in privileged mode.

Exposing the port (`9432`) with `-p 9432:9432` makes it easier connect to the
gRPC server from outside the container.

Also the path `/var/lib/bblfshd` should be mounted in the volume in order to
have stateful bblfshd instances between reboots.

If you are behind an HTTP or HTTPS proxy server, for example in corporate settings, you will need to add the `HTTP_PROXY`, `HTTPS_PROXY`, and `NO_PROXY` environment variables in the docker run command to configure HTTP or HTTPS proxy behavior.

```bash
$ docker run -d --name bblfshd --privileged -p 9432:9432 -e 
HTTP_PROXY=http://proxy.example.com:80/" -v /var/lib/bblfshd:/var/lib/bblfshd bblfsh/bblfshd
```

#### Installing the drivers

Now we need to install the driver images into the daemon, you can install
the official images just running the command:

```sh
$ docker exec -it bblfshd bblfshctl driver install --all
```

You can check the installed versions executing:

```sh
$ docker exec -it bblfshd bblfshctl driver list
```

```sh
+----------+-------------------------------+---------+--------+---------+--------+-----+-------------+
| LANGUAGE |             IMAGE             | VERSION | STATUS | CREATED |   OS   | GO  |   NATIVE    |
+----------+-------------------------------+---------+--------+---------+--------+-----+-------------+
| python   | //bblfsh/python-driver:latest | v1.1.5  | beta   | 4 days  | alpine | 1.8 | 3.6.2       |
| java     | //bblfsh/java-driver:latest   | v1.1.0  | alpha  | 6 days  | alpine | 1.8 | 8.131.11-r2 |
+----------+-------------------------------+---------+--------+---------+--------+-----+-------------+
```

To test the driver you can executed a parse request to the server with the `bblfshctl parse` command,
and an example contained in the docker image:

```sh
$ docker exec -it bblfshd bblfshctl parse /opt/bblfsh/etc/examples/python.py
```


### Running standalone

A standalone distributions of `bblfshd` and `bblfshctl` can be found at the
GitHub [release page](https://github.com/bblfsh/bblfshd/releases).

*bblfshd* is only provided for Linux distributions, since it relies on Linux
containers to run language drivers. And *bblfshctl* can be found for Windows,
macOS and Linux.


## Using bblfshctl

The binary *bblfshd* is provided with a sister tool called *bblfshctl*, this
command is used to monitor and manage the daemon.

If you are using the docker image this binary is inside of the `bblfsh/bblfshd`
image and can be using with a `docker exec`.

```sh
$ docker exec -it bblfshd bblfshctl --help
```

```sh
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

### Driver management

The *bblfshd's* drivers can be installed, updated and remove with the `driver`
command and his subcommands.

Installing all the official driver:

```sh
$ bblfshctl driver install --all
```

Overriding a single driver and a specific version:

```sh
$ bblfshctl driver install python bblfsh/python-driver:v1.1.5 --update
```

Listing all the available drivers

```sh
$ bblfshctl driver list
```

```sh
+----------+-------------------------------+---------+--------+---------+--------+-----+-------------+
| LANGUAGE |             IMAGE             | VERSION | STATUS | CREATED |   OS   | GO  |   NATIVE    |
+----------+-------------------------------+---------+--------+---------+--------+-----+-------------+
| python   | //bblfsh/python-driver:latest | v1.1.5  | beta   | 4 days  | alpine | 1.8 | 3.6.2       |
| java     | //bblfsh/java-driver:latest   | v1.1.0  | alpha  | 6 days  | alpine | 1.8 | 8.131.11-r2 |
+----------+-------------------------------+---------+--------+---------+--------+-----+-------------+
```
