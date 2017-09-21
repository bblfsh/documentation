# Babelfish clients

Over the network protocol and (soon) using
[`libuast`](https://github.com/bblfsh/libuast) there are some clients that allow
you to use Babelfish with a higher level API, abstracting some of the complexities
of [gRPC](https://grpc.io) communication, parsing, and the [base C `libuast`
API](https://github.com/bblfsh/libuast/tree/master/src).

## Existing clients

Currently there are working or planned clients in different stages of development
for these languages:

| Language | Status | Libuast | URL                                     |
| -------- | ------ | ------- | --------------------------------------- |
| Python   | Beta   | ✓       | https://github.com/bblfsh/client-python |
| Go       | Alpha  | ✓       | https://github.com/bblfsh/client-go     |
| Scala    | Alpha  | ✖       | https://github.com/bblfsh/client-scala  |

## Example

The client API's differ to adapt to their language specific idioms, the following
code shows a simple example with the Go client:

```go
package main

import (
	"fmt"
	"reflect"

	"gopkg.in/bblfsh/client-go.v1"
)

func main() {
	client, err := bblfsh.NewBblfshClient("localhost:9432")
	if err != nil {
		panic(err)
	}

	res, err := client.NewParseRequest().ReadFile("some_file.py").Do()
	if err != nil {
		panic(err)
	}
	if reflect.TypeOf(res.UAST).Name() != "Node" {
		fmt.Errorf("Node must be the root of a UAST")
	}
}
```

Currently we're integrating [`libuast`](https://github.com/bblfsh/libuast) into
the clients which will greatly increase the capabilities of the clients adding
[XPath](https://en.wikipedia.org/wiki/XPath)-like querying of the UAST tree.
