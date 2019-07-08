# Clients

There are some clients in different languages that provide a higher level API, built on top of [gRPC](https://grpc.io) and [`libuast`](https://github.com/bblfsh/libuast).

These clients make it easier to both parse and analyze the resulting UAST, abstracting from network communication and providing a query language to filter UASTs and can be used both as a library or as command line programs.

## Existing clients

There are clients for the following languages:

| Language | Status | UAST.v2 / libuast.v3 | URL |
| :--- | :--- | :--- | :--- |
| Python |  Beta | ✓ | [bblfsh/client-python](https://github.com/bblfsh/python-client) |
| Go | Beta | ✓ | [bblfsh/go-client](https://github.com/bblfsh/go-client) |
| Scala |  | [WIP](https://github.com/bblfsh/scala-client/issues/83)  | [bblfsh/client-scala](https://github.com/bblfsh/scala-client) |

## Examples

The client API's differ to adapt to their language specific idioms, the following code snippets show several simple examples with the Go, Python and Scala clients that parse a file and apply a filter to return all the simple identifiers:

### Go example

As a command, using [bblfsh-cli](https://github.com/bblfsh/go-client#cli):

```bash
bblfsh-cli -q [XPath query] -m semantic [file.ext]
```

As a library:

```go
package main

import (
	"context"
	"fmt"
	"time"

	bblfsh "github.com/bblfsh/go-client/v4"
	"github.com/bblfsh/go-client/v4/tools"
	"github.com/bblfsh/sdk/v3/uast"
	"github.com/bblfsh/sdk/v3/uast/nodes"
)

func main() {
	client, err := bblfsh.NewClientContext(context.Background(), "localhost:9432")
	if err != nil {
		panic(err)
	}
    defer client.Close()

	res, _, err := client.NewParseRequest().ReadFile("some_file.py").UAST()
	if err != nil {
		panic(err)
	}

	it, err := tools.Filter(res, "//*[not(@token = '') and not(@role='Qualified')]")
	if err != nil {
		panic(err)
	}
	for it.Next() {
		// Print the internal type
		n := it.Node()
		fmt.Printf("Type: %q (%T)\n", uast.TypeOf(n), n)

		node, ok := n.(nodes.Object)
		if !ok {
			continue
		}

		// Print the positions
		pos := uast.PositionsOf(node)
		fmt.Println("StartPos:", pos.Start(), " EndPos:", pos.End())

		// Print the token
		fmt.Println("Token:", uast.TokenOf(node))
	}

	// Get the normalized identifiers
	it, err = tools.Filter(res, "//uast:Identifier")
	if err != nil {
		panic(err)
	}
	for it.Next() {
		node, ok := it.Node().(nodes.Object)
		if !ok {
			continue
		}
		fmt.Println(node["Name"])
	}
}
```

### Python example

As a command:

```bash
python3 -m bblfsh -q [XPath query] -f [file.ext]
```

As a library:

```python
import bblfsh

if __name__ == "__main__":
    client = bblfsh.BblfshClient("0.0.0.0:9432")
    ctx = client.parse("some_file.py")

    it = ctx.filter("//uast:Identifier")
    for n in it:
        print(n.get())
```

### Scala example

As a command:

```bash
java -jar bblfsh-client-assembly-1.0.1.jar -q [XPath query] -f file.py
```

As a library:

```scala
import org.bblfsh.client.BblfshClient._

import gopkg.in.bblfsh.sdk.v1.protocol.generated.ParseResponse
import gopkg.in.bblfsh.sdk.v1.uast.generated.Node

import scala.io.Source

class BblfshClientParseTest {
  val fileName = "src/test/resources/SampleJavaFile.java"
  val fileContent = Source.fromFile(fileName) .getLines.mkString

  val resp = client.parse(fileName, fileContent)

  if (resp.uast.isDefined) {
     rootNode = resp.uast.get
     val filtered = client.filter(rootNode, "//*[@role='Identifier' and not(@role='Qualified')]")
     filtered.foreach{ println }
  } else {
    // ... handle resp.uast.errors
  }
}
```

## Query language

When using one of the clients that support libuast you can query the UAST result nodes using an [xpath-like](https://www.w3.org/TR/xpath/) query language. Check the [UAST querying page in this documentation](uast-querying.md) for the details.

## Iterators

The client also allows you to instance an Iterator object and iterate over the tree on several predefined orders:

* [Pre-Order](https://en.wikipedia.org/wiki/Tree_traversal#Pre-order)
* [Post-Order](https://en.wikipedia.org/wiki/Tree_traversal#Post-order)
* [Level-Order / Breadth first](https://en.wikipedia.org/wiki/Tree_traversal#Breadth-first_search)
* Position-Order \(this will retrieve the nodes in the same order as their position in the source code\).

To check the exact way to use an iterator you must consult the readme of the specific client you're using, but they're generally easy to use as this Python example shows:

```python
import bblfsh
client = bblfsh.BblfshClient("0.0.0.0:9432")
root = client.parse("/path/to/myfile.py")

for node in bblfsh.iterator(root, bblfsh.TreeOrder.PRE_ORDER):
    #... do stuff with the node
```

