# Nodes representation

## Overview

(U)AST in Babelfish is stored in a generic schema-less representation
similar to JSON, but that preserve all the primitive value types that
AST can express.

Babelfish's parsers will provide an AST in this form, which allow to process
them with the same transformation pipeline.

For the specific list of UAST types, see [Semantic UAST](./semantic-uast.md).

## Types

The representation defines four high-level node kinds:

* Null (Nil)

* Objects

* Arrays

* Values

Null nodes are used to represent an existing property with an unspecified value.

Objects and Arrays are composite node kinds, in the sense that they can
refer to other nodes in the UAST.

Values, on the other hand, are always leaf nodes.

This type system might remind the type system used in JSON and it is in
fact a more strict subset of it.

### Object

Object nodes are defined as an unordered set of key-value pairs. This
kind is used for nearly all UAST node types defined by the [Semantic UAST](./semantic-uast.md).

Keys can only be strings, similar to JSON encoding.

There can be only one key-value pair for a given key. Nodes that have
multiple children of the same type should store them in a single key-value
pair that refers to an Array node.

There is only one special key defined by this spec:

* `@type` - a type of the node. Used to enforce a specific UAST schema.

In general, all fields starting with `@` should be considered internal,
if not stated otherwise.

### Array

Arrays are ordered lists of other nodes. Similar to JSON arrays,
they can contain nodes with different types and even kinds.

### Values

Values are leaf nodes and can have one of the following type:

* String - UTF8 string value.

* Int - 64 bit integer.

* Uint - 64 bit unsigned integer.

* Float - 64 bit floating point value.

* Bool - boolean value.

The type system remind JSON, except the fact that the kind of a number is
preserved.
