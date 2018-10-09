# Semantic UAST

## Overview

The goal of the Semantic UAST is to provide a set of UAST node types with
a strictly defined semantic meaning that does not depend on the
programming language.

## Type system

Semantic UAST types are defined in the Babelfish SDK on top of the
[schema-less representation](./representation-v2.md).

The `@type` field in Object nodes is used to determine an exact type in
the Semantic UAST type system. Besides the Semantic UAST types, drivers
may emit language-dependant node types that were not yet covered by
Semantic UAST concepts.

### Namespaces

UAST node types can have a namespace similar to XML namespaces.

For example, Java AST defines a `Identifier` node type, while Go AST defines
a similar type called `Ident`, and the Semantic UAST has it's own
concept called `Identifier`.

To distinguish between these node types the `lang:` prefix is added to
each type, and `uast:` prefix is added for types defined by Semantic UAST.
The prefix without the `:` is called a **namespace**.

For our example, types listed above will be written in the following form
when adding namespaces:
`java:Identifier`, `go:Ident`, `uast:Identifier`.

### Common fields

As described in the [schema-less representation](./representation-v2.md)
spec, object fields starting with `@` are considered internal and may be
present on any object regardless of the type (schema).

This UAST specification defines few more special fields:

* `@pos` - stores the positional information related to this UAST node.
  See `Positions` type for more details.

* `@token` - a text representation of this node in the source file.
  This field is only available for compatibility reasons. If available,
  `@pos` should be used to get the source code corresponding to the UAST
  node.

* `@role` - stores an array with role codes. This field can be used to
  interpret native AST types that were not yet covered by Semantic UAST.

All other field are defined by the Semantic UAST schema.

### Types

#### Position

Represents a position in a source code file. Cannot have any fields except
ones defined below.

**@type:** `uast:Position`

Field | Type | Description
----- | ---- | -----------
`offset` | `Uint` | Position as an absolute byte offset (0-based index).
`line` | `Uint` | Line number (1-based index).
`col` | `Uint` | Column number (1-based index). The byte offset of the position relative to a line.

#### Positions

Object that stores all positional information for a node. This node kind

**@type:** `uast:Positions`

Field | Type | Description
----- | ---- | -----------
`start` | `uast:Position` | Start position of the node.
`end` | `uast:Position` | End position of the node.
`*` | `uast:Position` | Any number of custom positional fields.

Keys of this object can be arbitrary names for positional fields of the
UAST node. Only two fields are defined: `start` and `end` to allow users
to access source snippet related to the node.

As an example of a custom positional information, a ternary operator
`x ? y : z` node may store individual positions for `?` and `:` characters
as a separate `then` and `else` fields in `Positions` node.

#### Identifier

Identifier is a name for an entity. The name could be any valid UTF8 string.

**@type:** `uast:Identifier`

Field | Type | Description
----- | ---- | -----------
`Name` | `String` | An identifier name.

#### String

A UTF-8 string literal. `Format` parameter is a driver-specific string
format that was used for the literal in the source file.

**@type:** `uast:String`

Field | Type | Description
----- | ---- | -----------
`Value` | `String` | An unescaped and unquoted UTF-8 string value.
`Format` | `String` | Driver-specific format that was used for the literal in the source file.

#### QualifiedIdentifier

Qualified name consists of multiple identifiers organized in a hierarchy.
Identifiers are stored starting from the root level of hierarchy to the leaf.
The closest analogy is the filesystem path.

**@type:** `uast:QualifiedIdentifier`

Field | Type | Description
----- | ---- | -----------
`Names` | `[]uast:Identifier` | A path elements starting from the root of the hierarchy to the leaf.

#### Comment

Comments can span any number of lines. `Block` flag indicates that the
comment uses block syntax (`/* ... */` in Go) instead of line-comment
syntax (`//` in Go).

Comments might have a prefix and suffix for the whole comment, and each
comment line may also be prefixed with a `Tab` to express a following pattern:

```golang
/*
* This is a multiline
* block comment
*/
```

In this case the `Prefix` and `Suffix` will be set to `"\n"`, and
`Tab` would be set to `"* "`.

**@type:** `uast:Comment`

Field | Type | Description
----- | ---- | -----------
`Text` | `String` | An unescaped comment text (UTF-8).
`Prefix` | `String` | A prefix added to the first line of the comment.
`Suffix` | `String` | A suffix added to the last line of the comment.
`Tab` | `String` | A prefix added to each line of the comment.
`Block` | `Bool` | If the comment is a multi-line comment.

#### Block

Block groups multiple statements and enforces sequential execution of these statements.

Eventually, blocks will also include a reference to a scope if it defines one.

**@type:** `uast:Block`

Field | Type | Description
----- | ---- | -----------
`Statements` | `[]Node` | An ordered list of statements.

#### Alias

Aliases provide a way to assign a name to an entity or give it an
alternative name in a specific scope. An alias acts like an immutable
alias for an object. The only way to reassign the name used by an alias
in a specific scope is to shadow it in a new child scope.

Alias should contain a reference to the scope where a name should be defined.
But since scopes are not be covered by the current spec, an actual
definition of this relation will be specified in the future.

Examples of aliases are names for types, constants, functions,
local names for imports, local names for imported symbols, etc.

**@type:** `uast:Alias`

Field | Type | Description
----- | ---- | -----------
`Name` | `uast:Identifier` | A name that is assigned to an entity.
`Node` | `Node` | An entity that will be aliased by a new name.

#### Import

Imports are statements that can load external modules into a program or a library.

Import declaration can be described as a static statement in the sense
that an effect of it is not affected by code execution and is not affected
by the position of the node inside UAST.

An `Import` can either:

* Register all exported symbols in the target scope (`All == true`).
* Register specific symbols in the target scope (`len(Names) != 0`).
* Act as a side-effect import (both `All` and `Names` field are not set).

**@type:** `uast:Import`

Field | Type | Description
----- | ---- | -----------
`Path` | `uast:String | uast:Identifier | uast:QualifiedIdentifier | uast:Alias` | A name that is assigned to an entity.
`All` | `Bool` | Import all definitions from the modules into the scope.
`Names` | `[](uast:Alias | uast:Identifier)` | Import specific names from the module. Can refer to an `uast:Alias` to rename imported entities.

#### RuntimeImport

Runtime import has the same structure as an import declaration, but have
slightly different semantics. Runtime import may appear anywhere in
the code, thus it may be affected by code execution.

**@type:** `uast:RuntimeImport`

**Inherits:** `uast:Import`

#### RuntimeReImport

Runtime re-import has the same semantics as Runtime Import, but it will
re-execute an initialization code when importing the same package the second time.

**@type:** `uast:RuntimeReImport`

**Inherits:** `uast:RuntimeImport`
