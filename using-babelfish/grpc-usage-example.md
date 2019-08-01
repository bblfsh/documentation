# gRPC usage example

On the [previous section](babelfish-protocol.md) we checked the format of the messages exchanged using the gRPC protocol. Here, we're going to put that knowledge into practice with a simple example usage of the Babelfish server using that protocol.

This example is done in Go but you could use the gRPC interface using [any programming language that supports this protocol](http://www.grpc.io/about#osp).

## Getting and compiling the `.proto` file

The first step involving gRPC communication is getting the `.proto` file that defines the types involved in the communication. The [`.proto` format](https://developers.google.com/protocol-buffers/docs/proto) defines the structures and methods used for protocol buffer data. In the case of the Babelfish server protocol, you need to get [this `.proto` file](https://github.com/bblfsh/sdk/blob/94e3b212553e761677da180f321d9a7a60ebec5f/protocol/generated.proto) from the SDK. Then you have to compile it to a source file with the required structures and methods. How you generate this module is language-dependent but usually it involves installing gRPC using your language package management tool \(if it has a package for it\) and then use one of the provided tools.

For this example, we'll use the [GogoProtobuf implementation](https://github.com/gogo/protobuf) for the Go language. This provides the `protoc-gen-gogo` `.proto` compiler so we'll use it to generate the `.go` interface module:

```bash
$ protoc-gen-gogo --go_out=. generated.proto
```

This will create a generated.pb.go file that we could then directly import into out Go code. Since the SDK is also written in Go, you could skip this step and import the modules with the [already generated serializers](https://github.com/bblfsh/sdk/blob/94e3b212553e761677da180f321d9a7a60ebec5f/uast/generated.proto#L11) in the Babelfish's SDK.

## Making requests

Now we'll write a simple program that sends a request to get UAST of a simple "hello world" Python code that we'll provide.

```go
package main

import (
    "context"
    "fmt"
    "os"
    "time"

    "gopkg.in/bblfsh/sdk.v2/protocol"
    "gopkg.in/bblfsh/sdk.v2/uast"
    "google.golang.org/grpc"
)

func main() {
    // Connect to the running server
    conn, err := grpc.Dial("0.0.0.0:9432", grpc.WithTimeout(time.Second*2),
        grpc.WithInsecure())
    if err != nil {
        os.Exit(1)
    }
    client := protocol.NewProtocolServiceClient(conn)
    req := &protocol.ParseRequest{
        Filename: "hello.py",
        Content:  "print('hello world!')",
        Language: "python"}
}
```

Now that we've created a request, we need to send it \(previous code omitted\):

```go
    resp, err := client.Parse(context.TODO(), req)
```

## Reading and interpreting the response

The code in the previous section returned a `ParseResponse` object that will have the format of the [ParseResponse](babelfish-protocol.md#ParseResponse) as seen on the [server protocol](babelfish-protocol.md) page. You should check the `status` \(`Status` in the case of Go, since public members start with uppercase\); only a value of `protocol.Status.OK` will indicate success.

The most important member of the `ParseResponse` object is undoubtedly `uast` \(`UAST` in Go\). This will contain a `Node` object with the [structure detailed in the previous page](babelfish-protocol.md#Nodes). This first node returned would be the root node of the UAST, and you typically would iterate over the children nodes \(contained in the aptly named `children` field\) typically using [a visitor](https://en.wikipedia.org/wiki/Visitor_pattern) and reading the `token`s and `role`s in the tree to do your tool.

For demonstration purposes, we'll just write a simple function that iterates the tree in preorder and print the node tokens:

```go
func printTokens(n *uast.Node) {
    fmt.Println(n.Token)

    for _, c := range(n.Children) {
        printTokens(c)
    }
}
```

Now we only need to call this function on the root node in the `resp` variable:

```go
   // back to main
   printTokens(resp.UAST)
```

## Full source of the example

```go
package main

import (
    "context"
    "fmt"
    "os"
    "strings"
    "time"

    "gopkg.in/bblfsh/sdk.v2/protocol"
    "gopkg.in/bblfsh/sdk.v2/uast"
    "google.golang.org/grpc"
)

func main() {
    // Connect to the running server
    conn, err := grpc.Dial("0.0.0.0:9432", grpc.WithTimeout(time.Second*2),
        grpc.WithInsecure())
    if err != nil {
        os.Exit(1)
    }

    client := protocol.NewProtocolServiceClient(conn)
    req := &protocol.ParseRequest{
        Filename: "hello.py",
        Content:  "print('hello world!')",
        Language: "python"}

    resp, err := client.Parse(context.TODO(), req)
    if err != nil {
        os.Exit(1)
    }

    if resp.Status != protocol.Ok {
        fmt.Println("Parsing errors:", strings.Join(resp.Errors, ", "))
        os.Exit(1)
    }

    printTokens(resp.UAST)
}

func printTokens(n *uast.Node) {
    fmt.Println(n.Token)

    for _, c := range n.Children {
        printTokens(c)
    }
}
```

## Supported Languages

Protocol v2 and higher include a `SupportedLanguages` method that permits the client to
discover what languages the `bblfshd` currently supports. The [SupportedLanguagesResponse]( https://godoc.org/gopkg.in/bblfsh/sdk.v3/protocol#SupportedLanguagesResponse)
returns a list of driver manifests, giving the names and aliases of the languages each driver
understands, along with version, development status, and feature support tags
(see the [Manifest](https://godoc.org/gopkg.in/bblfsh/sdk.v3/driver/manifest#Manifest) type).

This method is supported by SDK versions ≥ 3.1.0.

Example (json):
```json
[
    {
        "name": "C++",
        "language": "cpp",
        "aliases": [
            "C++",
            "C",
            "CUDA",
            "OpenCL",
            "Metal"
        ],
        "version": {
            "version": "v1.4.0",
            "build": {
                "seconds": 1559053620
            }
        },
        "status": "DEV_BETA",
        "features": [
            "ast",
            "uast",
            "roles"
        ]
    },
    {
        "name": "Java",
        "language": "java",
        "version": {
            "version": "v2.7.2",
            "build": {
                "seconds": 1559055732
            }
        },
        "status": "DEV_BETA",
        "features": [
            "ast",
            "uast",
            "roles"
        ]
    },
    {
        "name": "JavaScript",
        "language": "javascript",
        "aliases": [
            "JS",
            "JSX"
        ],
        "version": {
            "version": "v2.9.0",
            "build": {
                "seconds": 1559054418
            }
        },
        "status": "DEV_BETA",
        "features": [
            "ast",
            "uast",
            "roles"
        ]
    }
]
```
