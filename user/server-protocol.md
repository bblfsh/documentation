# Babelfish Protocol

The Babelfish client-server protocol uses [gRPC](http://www.grpc.io) for method
selection (with the inherent Protobuf3 used for message exchanges). You can read
[the server](https://github.com/bblfsh/sdk/blob/master/protocol/generated.proto)
and [SDK](https://github.com/bblfsh/sdk/blob/master/uast/generated.proto) `.proto`
files to see the format description of the messages and types involved, but we'll
provide here a simple definition in JSON-like format.

## ParseUASTRequest

Issued by the client to request that a source code file must be parsed to 
an UAST tree. The client must provide the code in the `content` field, the
programming language in the `language` field (which obviously must be one
of the languages currently supported and the `filename`
field with the name of the file containing the source code.

Example:

```json
{
    "filename": "mytest.py",
    "language": "python",
    "content": "print('hello world')"
}
```

## ParseUASTResponse

This is the reply produced by the server as response to the above
ParseUASTRequest. Example:

```json
{
    "status": 0,
    "errors": "",
    "uast": {...},
}
```

The "uast" field would contain the UAST root node as an
[`github.com.bblfsh.sdk.uast.Node`
type](https://github.com/bblfsh/sdk/blob/master/uast/generated.proto#L11) which as
you can see in the linked definition includes the internal type (the type used by
the native AST), a map with the properties, the UAST roles, the position of the
source construct that generated the node and a list of children. 
