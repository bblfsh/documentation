# Babelfish Protocol

The Babelfish client-server protocol uses [gRPC](http://www.grpc.io) for method selection \(with the [Protocol Buffers](https://developers.google.com/protocol-buffers/) format used for message serialization\).
On the [next page](grpc-usage-example.md) we'll see a demo of how this comes together in practice with some code.

## Protocol v2

Protocol v2 introduced a few changes in data structures and in methods' signatures.
For instance `ProtocolService` was split into two services - [Driver](https://github.com/bblfsh/sdk/blob/v3.2.2/protocol/driver.proto#L55) which implements `Parse` method and [DriverHost](https://github.com/bblfsh/sdk/blob/v3.2.2/protocol/driver.proto#L103) which implements server related methods like `ServerVersion` and `SupportedLanguages`.
`ParseRequest` was extended by `Mode` field, so clients may decide what transformations will be used (native, preprocessed, annotated, semantic).
`ParseResponse` type was simplified. It contains the language, [binary encoded uast](../uast/uast-encoding-v2.md) and potential error message.

Example request:
```json
{
    "content": "unsigned long long fib(int n);\n\nint main() {\n    fib(12);\n    return 0;\n}\n\nunsigned long long fib(int n) {\n    return (n <= 1) ? 1ULL : fib(n-2) + fib(n-1);\n}",
    "filename": "test.c"
}
```
Example response:
```json
{
    "uast": "\000bgr\001\000\000\000\005\010\244\003\020\001\020\...",
    "language": "c"
}
```

Also `SupportedLanguagesResponse` was simplified. The structure wraps list of language driver manifests. Every manifest contains all needed information plus optional fields like aliases (so far only for [certain languages](../languages.md))

More examples you can find on the [next page](grpc-usage-example.md).

For more details you can look at [protocol definition](https://github.com/bblfsh/sdk/blob/v3.2.2/protocol/driver.proto)

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
