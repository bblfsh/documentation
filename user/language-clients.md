# Babelfish clients

There are some clients in different languages that provide a higher level API,
built on top of [gRPC](https://grpc.io) and [`libuast`](https://github.com/bblfsh/libuast).

These clients make it easier to both parse and analyze the resulting UAST,
abstracting from network communication and providing a query language to
filter UASTs.

## Existing clients

There are clients for the following languages:

| Language | Status | Libuast | URL                                     |
| -------- | ------ | ------- | --------------------------------------- |
| Python   | Beta   | ✓       | https://github.com/bblfsh/client-python |
| Go       | Beta   | ✓       | https://github.com/bblfsh/client-go     |
| Scala    | Beta   | ✓       | https://github.com/bblfsh/client-scala  |

## Example

The client API's differ to adapt to their language specific idioms, the following
code shows a simple example with the Go client that parsers a Python file
and applies a filter to return all the simple identifiers:

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

	query := "//*[@roleIdentifier and not(@roleQualified)]"
	nodes, _ := tools.Filter(res.UAST, query)
	for _, n := range nodes {
		fmt.Println(n)
	}
}
```

## Query language

[`libuast`](https://github.com/bblfsh/libuast) provides a [xpath](https://www.w3.org/TR/xpath/) based query language for UASTs.

Any of the [node](https://godoc.org/github.com/bblfsh/sdk/uast#Role) fields can be used for querying, which are mapped in the following way:

* `InternalType` is converted to the element name
* `Token`, if available, is converted to an attribute with `token` as keyword and the actual token as value
* Every `Role` is converted to an attribute concatenating a `role` prefix and the role name in CamelCase.
* Every `Property` is converted to an attribute with the property keyword as keyword and the property value as value
* `StartPosition`, if available, is mapped to three attributes:
  * A `startOffset` attribute, with the offset as value
  * A `startLine` attribute, with the line as value
  * A `startCol` attribute, with the column as value
* `EndPosition`, if available, is mapped to three attributes:
  * A `endOffset` attribute, with the offset as value
  * A `endLine` attribute, with the line as value
  * A `endCol` attribute, with the column as value

which are mapped in to XML in the following way:

```xml
<{{InternalType}}
    token='{{Token}}'
	{{for role in Roles}}
	role{{role}}
	{{for key, value in Properties}}
	{{key}}='{{value}}
	startOffset={{StartPosition.Offset}}
	startLine={{StartPosition.Line}}
	startCol={{StartPosition.Col}}
	endOffset={{EndPosition.Offset}}
	endLine={{EndPosition.Line}}
	endCol={{EndPosition.Col}}>
	{{Children}}
</{{InternalType}}>
```

This means that both language specific queries (`InternalType`, `Properties`) and language agnostic queries (`Roles`) can be done.
For example:

- Return all the numeric literals in Python: `//NumLiteral`
- Return all the numeric literals in ANY language: `//*[@roleNumber and @roleLiteral]`
- Return all the integer literals in Python: `//*[@NumType='int']`

The query language also allow some more complex queries:

- All the elements in the tree that have either start or end offsets: `//*[@startOffset or @endOffset]`
- All the simple identifiers: `//*[@roleIdentifier and not(@roleQualified)]`
- All the simple identifiers that don't have any positioning: `//*[@roleIdentifier and not(@roleQualified) and not(@startOffset) and not(@endOffset)]`
- All the arguments in function calls: `//*[@roleCall and @roleArgument]`
- All the numeric literals in binary arithmetic operators: `//*[@roleBinary and @roleOperator and @roleArithmetic]//*[@roleNumber and @roleLiteral]`
