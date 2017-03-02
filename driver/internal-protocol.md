
# Internal Driver Protocol

The **internal driver protocol** is used to communicate between the native AST
parser and the UAST converter inside a driver.

## Transport

The internal driver protocol uses standard input and standard output as transport.
It follows the same rules as the [driver protocol](protocol.md#transport).

## Encoding

Inside the driver, internal communication between the UAST converter (Go) and
the AST parser (native) uses JSON. Each message is encoded as a single-line
JSON, ending with a new line.

## Processes

### Parse AST

The Parse AST process is the same as its external interface in the
[driver protocol](protocol.md#parse-ast).
