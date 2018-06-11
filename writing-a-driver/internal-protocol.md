# Internal Protocol

The **internal driver protocol** is used to communicate between the native AST parser and the UAST normalizer inside a driver.

## Transport

The internal driver protocol uses standard input and standard output as transport.

## Encoding

Inside the driver, internal communication between the UAST normalizer \(Go\) and the code to AST parser \(native\) uses JSON. Each message is encoded as a single-line JSON, ending with a new line.

## Processes

### Parse code

This is currently the default and only request supported by the internal driver protocol. It **parses a file and returns its AST**. A request contains the content of the file being analyzed as a string.

**Request** message has the following structure:

```text
{
    "content": <content>
}
```

**Response** structure is:

```text
{
    "status": <status> ("ok", "error", "fatal")
    "errors": [ <error message>, <error message>, ... ]
    "metadata": <metadata> (string dict)
    "ast": <AST> (object)
}
```

If the parsing is successful, `status` is `ok`. If the file could be parsed \(AST was generated\) but had parsing errors, `status` is `error`. If the file could not be parsed at all \(no AST\), `status` is `fatal`.

`errors` might contain any parsing errors found. If `status` is `ok`, then `errors` should be not set.

`metadata` is an optional string map containing arbitrary metadata, such as language version or dialect if it had to be detected before parsing.

Note that **binary files are not supported** by this process at the moment. If we want to add support for [Piet](http://www.dangermouse.net/esoteric/piet.html) in the future, we will add a binary content field.

* **TODO: address non-UTF-8 encoding problems.**

#### Example

```text
[request (pretty printed)]
{
    "content": "#!/bin/bash\nexec foo\n"
}
[response (pretty printed)]
{
    "status": "ok",
    "metadata": {
        "dialect": "bash"
    },
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

