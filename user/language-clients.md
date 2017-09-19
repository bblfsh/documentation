# Babelfish clients

Over the network protocol and (soon) using
[`libuast`](https://github.com/bblfsh/libuast) there are some clients that allow
you to use Babelfish with a higher level API, abstracting some of the complexities
of [gRPC](https://grpc.io) communication, parsing, and the [base C `libuast`
API](https://github.com/bblfsh/libuast/tree/master/src).

## Existing clients

Currently there are working or planned clients in different stages of development
for these languages:

| Language | Status  | Libuast | URL                                    |
| -------- | ------- | ------- | -------------------------------------- |
| Python   | Beta    | ✓       | https://github.com/bblfsh/client-scala |
| Go       | Alpha   | ✓       | https://github.com/bblfsh/client-scala |
| Scala    | Alpha   | ✖        | https://github.com/bblfsh/client-scala |
| Java     | Planned | ✖        | https://github.com/bblfsh/client-java  |

## Example

The client API's differ to adapt to their language specific idioms, the following
code shows a simple example with the Python client:

```python
from bblfsh import BblfshClient
# Protobuf definition file:
from github.com.bblfsh.sdk.uast.generated_pb2 import Node
from bblfsh.launcher import ensure_bblfsh_is_running

if __name__ == "__main__":
    # Create a Babelfish server instance if it's not already running:
    ensure_bblfsh_is_running()

    client = BblfshClient("0.0.0.0:9432")

    # Using language autodetection:
    uast = client.parse("some_file.py")
    # ...or manually specifying the language:
    uast2 = client.parse("some_file.py", language="Python")

    # Node is the root of the UAST tree:
    assert(isinstance(uast.uast, Node))
```

Currently we're integrating [`libuast`](https://github.com/bblfsh/libuast) into
the clients which will greatly increase the capabilities of the clients adding
[XPath](https://en.wikipedia.org/wiki/XPath)-like querying of the UAST tree.
