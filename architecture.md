
# Architecture

## Overview

Our main building block is the **language driver**, each language driver provides
parsing services for one or more languages. It can be written in any language and
is packaged as a Docker container. These containers are executed by the babelfish
**server** in a very specific runtime.

```
 ________          ________                _________________________
|        |  grpc  |        | stdin/stdout |                         |
| Client | <----> | Server | <----------> |     Language Driver     |
|________|        |________|   protobuf   |_________________________|
                                          |                         |
                                          |         Container       |
                                          |_________________________|
```

## Language Drivers

A language driver is a Docker container that takes file parsing requests and
respond with the AST or the UAST.

The name of the container should be the language name. It may also contain the
language version if different versions require different containers. Containers
must be tagged with [Semantic Versioning 2.0.0](http://semver.org/), where
semver compatibility semantics are applied to the returned AST.

Our implementations of language drivers are composed of two pieces: an AST parser,
written in the source language and a AST to UAST converter written in Go. In the
case of AST parsers written in Go, AST parser and UAST converter might be
combined in a single executable.

The entry point of the container is the Go UAST converter, which wraps the
language-specific AST parser. This is how the Python driver looks like:

```
                 _________________________________________
                |                     |                   |
  stdin/stdout  |                 stdin/stdout            |
  <---------->  | UAST converter <-----------> AST parser |
    protobuf    |     (Go)           JSON        (Python) |
                |_____________________|___________________|
                |                                         |
                |                 Container               |
                |_________________________________________|
```

## Server

The server exposes a parsing interface for every language. It takes requests
which it passes to the appropriate language driver. It maintains warm instances
of language drivers, pools and handles any other orchestration needs.

### Driver Runtime

The server executes driver in a very constrained **driver runtime**. It does not
use Docker but its own lightweight container runtime.

Drivers must assume:

* No network interface.
* Only one core available.

Currently we're allowing drivers to write on disk, although this might be
disallowed in the future.
