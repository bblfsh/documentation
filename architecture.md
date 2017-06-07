
# Architecture

## Overview

Our main building block is the **language driver**, each language driver provides
parsing services for one language. It can be written in any language and is
packaged as a standard Docker container. These containers are executed by the
babelfish **server** in a very specific runtime.

```mermaid
graph LR
    Client-- TCP<br/>gRPC -->Server
    Server-- stdin/stdout<br/>gRPC -->Driver
    subgraph Container
        Driver
    end
```

## Language Drivers

A language driver is a program, containerized with Docker, that takes file parsing
requests with source code and reply with an [UAST](./uast/specification.md)
representation of the provided source code or an error message.

Our implementations of language drivers are composed of two pieces: an AST parser,
that can be written in any language (usually the source language) and an AST to
UAST normalizer written in Go. In the case of AST parsers written in Go or other
languages producing linkable object files or shared library, AST parser and UAST
normalizer might be combined by dynamic or static linking without a separate
runnable component.

The entry point of the container is the Go UAST normalizer, which communicates
internally with the language-specific AST parser. This is how the Python driver
looks like:

```mermaid
graph LR
    Server-- stdin/stdout<br/>gRPC -->UAST
    subgraph Container
        UAST["UAST Normalizer<br/>(Go)"]-- stdin/stdout<br/>JSON -->AST["AST Parser<br/>(Python)"]
    end
```

## Server

The server is the higher level component of the architecture managing
client requests (done via [gRPC](http://www.grpc.io) using [a simple 
protocol](driver/protocol.md)) with the containerized
language drivers which it orchestate to keep warm instances, pools and handing
any other orchestation needs.

The server itself it's also designer to run inside a container allowing an [easy
deployment](user/getting-started.html#running-with-docker-recommended)
and operation.

### Driver Runtime

The server executes driver in a very constrained and lightweight **driver
runtime** based on
[libcontainer](https://github.com/opencontainers/runc/tree/master/libcontainer).
The runtime executes the driver contained in a standard Docker container in an
isolate namespace and cgroup avoiding Docker's complex filesystems or networking
features.

Drivers must assume:

* No network interface.
* Only one core available.

Currently we're allowing drivers to write on their container disk, although this
might be disallowed in the future.
