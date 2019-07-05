# Babelfish Protocol

The Babelfish client-server protocol uses [gRPC](http://www.grpc.io) for method selection \(with the [Protocol Buffers](https://developers.google.com/protocol-buffers/) format used for message serialization\).
On the [next page](grpc-usage-example.md) we'll see a demo of how this comes together in practice with some code.

## Protocol v2

<!-- FIXME(dennwc): explain the v2 protocol, list relevant proto files and UAST decoding spec -->

## Protocol v1

You can read [the server](https://github.com/bblfsh/sdk/blob/94e3b212553e761677da180f321d9a7a60ebec5f/protocol/generated.proto#L11) and [SDK](https://github.com/bblfsh/sdk/blob/94e3b212553e761677da180f321d9a7a60ebec5f/uast/generated.proto) `.proto` files to see the format description of the messages and types involved, but we'll provide here a simple definition in JSON-like format.

### ParseRequest

Issued by the client to request that a source code file must be parsed to an UAST tree. The client must provide the code in the `content` field, the programming language in the `language` field \(which must be one of the [languages currently supported](../languages.md) or empty to enable auto-detection\) and the `filename` field with the name of the file containing the source code.

Example:

```json
{
    "filename": "mytest.py",
    "language": "python",
    "content": "print('hello world')"
}
```

**Note:** The Babelfish server orchestrates the language parser containers depending on the demand which could mean that the first request for a language could take some time while it retrieves and starts the parsers.

### ParseResponse

This is the reply produced by the server as response to the above ParseRequest. Example:

```json
{
    "status": 0,
    "errors": "",
    "uast": {...},
}
```

The `uast` field would contain the UAST root node as an [`gopkg.in/bblfsh/sdk.v2/uast` type](https://github.com/bblfsh/sdk/blob/94e3b212553e761677da180f321d9a7a60ebec5f/uast/generated.proto#L11) which as you can see in the linked definition includes the internal type \(the type used by the native AST\), a map with the properties, the UAST roles, the position of the source construct that generated the node and a list of children as we'll see in the next section.

The status contains the return code of the Request. If it's != 0 \(which will be mapped to "Ok" or "ok" in an enum in the clients\) no further processing should be done for that request since it failed to parse the code.

### Nodes

```json
{
    "internal_type": "someNativeType",
    "properties": {...},
    "children": [{...Node...}, {...Node...}, ...],
    "token": "symbolicOrLiteralName",
    "start_position": 10,
    "end_position": 20,
    "roles": ["role1", "role2", ...]
}
```
