
# Driver Protocol

## Transport

Standard input and output is used as transport. This applies both to the
communication between **server and driver** as well as the internal communication
in the driver **UAST converter and AST parser**.

The server can terminate the process cleanly by closing the standard input pipe.
Closing the pipe while there is a pending response not read from standard output
is considered an error, and drivers should return a non-zero exit code. The
client must not perform half-closes (closing standard input if it expects to
continue reading from standard output).

All requests and responses have **sequential order**. Response to request N will
not start until response to request N-1 has finished.

The workflow of the protocol is as follows:

```mermaid
sequenceDiagram
    participant Server
    participant Driver
    loop MainLoop
        Server->>Driver: Request
        Driver-->>Server: Response
    end
```

## Encoding

Server-driver communication uses **protobuf 3** for message encoding. In the
driver side, this is implemented in the UAST converter in Go, using the
babelfish SDK.

* **TODO: Add proto**

## Processes

### Info

Info returns information about the driver. There is a general implementation in
the babelfish SDK that uses a manifest to provide this information.

* **TODO: Add proto**

**Request** message has the following structure:

```
{
    "action": "Info"
}
```

Language is optional, and only used for drivers that support multiple languages.
If language is not given, a driver default is assumed. If language is given and
not supported by the driver, an error is returned.

**Response** structure is:

```
{
    "driver" <driver name and version> (string)
    "supports": [
        {
            "language": <language> (string)
            "language_version": <version range> (version range)
            "metadata": <metadata> (string to string map)
        }
        ...
    ]
}
```

Metadata is a key-value map of metadata values such as compiler or runtime
versions.

* **TODO: add version range definition**

### Parse AST

This process **parses a file and returns its AST**. A request contains the content
of the file being analyzed as a string.

It must be implemented by the AST parser.

**Request** message has the following structure:

```
{
    "action": "ParseAST"
    "language": <language> (string, optional)
    "language_version": <language version> (string, optional)
    "content": <content> (string)
}
```

Language is optional, and only used for drivers that support multiple languages.
If language is not given, a driver default is assumed. If language is given and
not supported by the driver, an error is returned.

**Response** structure is:

```
{
    "status": <status> ("ok", "error", "fatal")
    "errors": [ <error message>, <error message>, ... ]
    "driver" <driver name and version> (string)
    "language": <language> (string)
    "language_version": <language version> (string)
    "ast": <AST> (object)
}
```

`driver` corresponds to the Docker container image and version. For example:
`parser-java:1.0`.

`language` should be specified in `CamelCase`.

If the parsing is successful, `status` is `ok`. If the file could be parsed
(AST was generated) but had parsing errors, `status` is `error`. If the file
could not be parsed at all (no AST), `status` is `fatal`.

`errors` might contain any parsing errors found. If `status` is `ok`, then
`errors` should be not set.

Note that **binary files are not supported** by this process at the moment. If we
want to add support for [Piet](http://www.dangermouse.net/esoteric/piet.html) in
the future, we will add a binary content field.

**Example:**

```
[request]
{
    "action": "ParseAST",
    "content": "#!/bin/bash\nexec foo\n"
}
[response]
{
    "driver": "bash:1.0",
    "ast": {
        "name": "script",
        "children": [
            { "name": "shebang", "bin": "/bin/bash" },
            { "name": "statement", "bin": "exec",
              "args": [ { "name": "string", "content": "foo" } ]
              }
        ]
    }
}
```

### Parse UAST

Equal to the *parse AST* process, but for UAST. It must be implemented by the
UAST converter.

**Request:**

```
{
    "action": "ParseUAST"
    "language": <language> (string, optional)
    "language_version": <language version> (string, optional)
    "content": <content> (string)
}
```

**Response** structure is:

```
{
    "status": <status> ("ok", "error", "fatal")
    "errors": [ <error message>, <error message>, ... ]
    "driver" <driver name and version> (string)
    "language": <language> (string)
    "language_version": <language version> (string)
    "uast": <UAST> (object)
}
```

* **TODO: Add proto**
