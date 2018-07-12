# Advanced Usage

## Adding all drivers

In the [previous section](getting-started.md) the `bblfshctl` command to install the drivers was shown with the `--recommended` switch. This will install all the drivers in beta stage or better and annotated UAST support. But if you're interested in installing [all official language drivers](../languages.md), even the ones that have alpha status and only native AST support, you can use the `--all` switch instead:

```bash
$ docker exec -it bblfshd bblfshctl driver install --all
```

## Adding drivers from the local Docker daemon

You can add a driver to the server using one stored in the local Docker daemon instead of the official ones at Dockerhub. This is especially useful when developing new drivers as it allows you to test versions of your driver integrated into the server.

If you want to do this, the Unix socket that Docker uses to communicate with the instance must be mounted in the bblfshd image, which you can do with the `-v` parameter like in the next example:

```bash
$ docker run -d -v /var/run/docker.sock:/var/run/docker.sock --name bblfshd --privileged -p 9432:9432 -v /var/lib/bblfshd:/var/lib/bblfshd bblfsh/bblfshd
```

Now you can add drivers from the local Docker using `bblfshctl` like this:

```bash
docker exec -it bblfshd bblfshctl driver install python docker-daemon:bblfsh/python-driver:dev-123321-dirty
```

## Running a driver without the server

You can also directly run a driver's Docker image without a server. Like the bblfshd server, it will serve using the gRPC protocol on the 9432 port \(through you can easily change it using the `-p` option to `docker run`\).

Running a driver this way means that requests will be processed serially so no more than one request at a time will be served and of course no more than the driver's language can be parsed, through it can be convenient for some quick UAST extraction of specific files or quick tests since it skips the steps of running a server and managing its drivers.

To run a driver independently of the server run:

```bash
docker run --rm -p 9432:9432 bblfsh/python-driver
```

\(change `python-driver` for the driver that you want to run\).

Then you can send requests using any of the clients.

