# Schema-less internal representation for AST

| Field | Value |
| --- | --- |
| BIP | ??1 |
| Title | Schema-less internal representation for AST |
| Author | Denys Smirnov |
| Status | Draft |
| Created | 2018-02-23 |
| Updated | 2018-02-23 |
| Target version | 1.15 ?? |

## Abstract

Current transformation code for AST/UAST nodes is based on Go structures
that makes it harder to manipulate. Transformation code usually cannot be
reused since it acts on specific fields of the Go struct (either children
or special-cased fields like type, token, position, etc).

We propose to change internal representation to a structure with
only few basic types: object, list, primitive (string, int, etc).
Additional code will be used to map this structure to current Node
structure to preserve backward compatibility.

We show that this representation will make transformation code easier to
write and extend. Also, generic transformation can be reused for any kind
of fields and AST structures.

## Rationale

Currently UAST node is defined in a following way:

```golang
type Node struct {
	InternalType  string
	Properties    map[string]string
	Children      []*Node
	Token         string
	StartPosition *Position
	EndPosition   *Position
	Roles         []Role
}
```

This structure has some issues:

* There are multiple special-cased properties: `InternalType`, `Token`,
  `StartPosition`, `EndPosition`. They cannot be manipulated with generics
  transformations defined by SDK, because they needs to be accessed as
  specific fields.

* List of child nodes (`Children`) is flat, without grouping by property
  name (predicate, or internalRole in terms of SDK). It makes it harder
  to process all nodes of a certain predicate as a whole (list of function
  arguments, for example).

* There is a set of primitive properties defined as a map (`Properties`).
  Again, generic transformations that are defined on child nodes cannot
  be used for these properties.

SDK works around these issues by defining a complex `ObjectToNode` class
that performs typical manipulations on special-cased fields. To be able
to represent any arbitrary transformation it defines a callback for the
driver to change AST nodes in arbitrary way, offloading work from SDK to
driver developers.

Currently manipulation pipeline looks like this:

```
| Native | --------> |   Go   | --------> | Transformations | ------>
| Driver |   JSON    | Driver |   Node    |   code in SDK   |  Node
```

Since native driver protocol uses arbitrary JSON objects as a representation
for an AST, this schema-less approach can be used in SDK itself to manipulate
the tree and conversion to Node struct can be postponed to the end of pipeline:

```
| Native | --------> |   Go   | --------> | Transformations | ------>
| Driver |   JSON    | Driver |  map[x]y  |   code in SDK   |  Node
```

This approach will allow to manipulate all properties of an AST objects
in a uniform way.

As an example of uniform transformations, we will consider following
transformations in `ObjectToNode`: `InternalType` and `TokenKeys`.

First transformation moves a values of a node type key of a native AST
from `Properties` map to `InternalType` field.

Second transformation moves any specified keys from `Properties` map
to `Token` field.

It's obvious that these transformations are exactly the same, but they
have to be implemented differently because they access different fields
in the structure. On the other hand, if these target fields would be
stored uniformly, it will be possible to reuse the code for this transformation.

Second example is a family of `Offset`/`Line`/`Column` transformations
that act on node position in the source file. A typical code for these
transformations is to find a key in `Properties` map, initialize a
`Position` structure and set a specific field for it. Again, since
fields of `Position` object should be accessed directly, there is no way
to share the code without making it too verbose (via reflection).

## Specification

Internal representation of AST tree is changed from `Node` tree to
a tree of `InternalNode` objects:

```golang
// InternalNode is an internal representation of AST node used in transformations.
// It can be one of the following kinds: Object, List, Primitive.
type InternalNode interface{}

type Object map[string]InternalNode

type List []InternalNode

type Primitive struct{
    Value interface{} // string or int
}
```

Legacy `Node` structure is mapped to `Object` in the following way:

* `InternalType` is stored under `@type` key as a string value.

* `Properties` are saved directly into a parent object map.

* `Children` is grouped by their predicate and store as single objects
  in the parent map, or as lists of objects.

* `Token` is stored under `@token` key.

* `StartPosition` is stored as `@start` key, and
  `EndPosition` is stored as `@end`.

* `Roles` are stored as a list on a `@role` key.

Transformation code should to be rewritten to act on new types.
Code in `ObjectToNode` needs to be reviewed for potential duplicates.

Annotation code changes could be omitted, since conversion to `Node`
structure might happen before annotations, but we want to make
an additional effort to port this code to use a new representation.

The reasoning behind is that future BIPs will introduce more generic
approach to transformation and annotation of an AST, thus both of these
types of operations needs to act on a single data model.

## Alternatives

*   Use current approach. Transformations that are required for real-world
    ASTs require moving nodes at the depth of 3-5 levels, which will force
    developer to manipulate and check fields of child nodes directly.
    Code for these check usually cannot be reused since it captures state
    of higher level nodes, or accesses different kind of special fields.

*   Use reflection to manipulate structures and abstract away actual
    transformations. This approach is close to ideal in a sense that
    it preserve strict Go structures, while allowing arbitrary manipulations
    to be made. In practice, however, we find this approach harder to support
    and it provides less flexibility in the end. It's harder to support
    because reflection code is too verbose and requires to many type
    conversion and safety checks to be made by developer. It's less flexible
    because reflection forces strict matching of types even while tree
    manipulation is in progress. Approach that we propose avoids this by
    postponing translation to Go structs to the end of tree manipulation.

## Impact

Most drivers will not be affected by this change, since mapping to legacy
`Node` structure will be preserved, and all SDK helpers will have the
same SDK as before. Drivers that utilize low-level SDK features will be
affected.

Protocols will not be affected, since the change only affects internal
representation of Node in SDK in transformation pipeline.

We expect no changes to generated ASTs. Any changes to the structure
will be introduced as a separate proposal.

Proposed changes might have minor effect on low-level SDK API (defining
custom transformations).

Changes that needs to be done to SDK:

* Rewrite transformation functions (`ObjectToNode`, `Positioner`, etc).
* Update internal annotation code and predicates that it uses.
  API or annotation DSL should not change.
* Update driver that used low-level SDK features like custom transformation
  functions.
