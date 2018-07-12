# BIP5: Shape-based AST transformations

| Field | Value |
| --- | --- |
| BIP | 5 |
| Title | Shape-based AST transformations |
| Author | Denys Smirnov |
| Status | Accepted |
| Created | 2018-02-23 |
| Updated | 2018-04-13 |
| Target version | 1.x |

## Abstract

Currently, an AST transformation code exists mostly in `ObjectToNode` class that applies some hand-written transformation over the tree, such as rewriting type fields, promoting nodes, etc. This approach is not scalable and does not allow to reuse the code of these transformations. Further, existing transformer implementations have repetitive code for walking the tree and applying transformations that is also not ideal.

We propose to define a transformation DSL for Go similar to current annotation DSL. Existing transformers will be expressed using this library.

It will be shown that this approach provides a solid base for planned high-level UAST transformations and will make the code more portable and easier to read. Also, having such declarative approach will eventually allow us to use indexes for the tree, that can be used to apply transformations faster for large code bases.

## Rationale

Currently we have this code as a part of `ObjectToNode` transformation:

```go
case c.OffsetKey == k:
    i, err := toUint32(o)
    if err != nil {
        return err
    }

    if n.StartPosition == nil {
        n.StartPosition = &Position{}
    }

    n.StartPosition.Offset = i
case c.EndOffsetKey == k:
    i, err := toUint32(o)
    if err != nil {
        return err
    }

    if n.EndPosition == nil {
        n.EndPosition = &Position{}
    }

    n.EndPosition.Offset = i
// ...
```

This code applies a similar transformation to different AST node keys.

It's obvious that transformations cannot be reused easily because they are using different struct fields. Thus, the approach described below will require [Schema-less internal representation for AST](bip5-shape-based-ast-transformations.md).

You may notice that the code in fact only needs to know the source and target fields to apply the transformation, plus a knowledge about the structure of input and output trees. We can a DSL that specifies an exact shape of both input and output trees with placeholders \(or variables\) in place of node values that need to be moved.

Here is an example of DSL for the source tree above:

```go
Part(Obj{
    OffsetKey: Int2Str(Var("off")),
})
```

It defines an operation that needs to find all nodes with a non-optional key `OffsetKey`, store it into variable `off` and convert it to an integer.

Note that conversion is written as int-to-string because this is easier to reason about the transformation in terms of writing nodes, as if a function call is not describing the structure, but constructing it: get variable `off`, convert to string, store it as the key `OffsetKey`.

You can see that a `Part` helper wraps the main transformation. It will save all other node keys that were not used in previous transformations.

But this is only the part of the transformation to read an input tree, we need to define the second one that will write an output tree:

```go
Part(Obj{
    "EndPosition": Obj{
        KeyType: String("Position"),
        "Offset": Var("off"),
    },
})
```

Or with some helpers:

```go
Part(Obj{
    "EndPosition": TypedObj("Position",
        "Offset": Var("off"),
    ),
})
```

It defines a shape of the tree in the same terms: the root object should have an `EndPosition` key, that has an object value of type `Position` and this new object should have a key `Offset` with our stored `off` variable.

Note that this time str-to-int conversion is omitted - it only happens on one side of the transformation, because the variable should store a "canonical" representation of that value, as defined by final tree shape.

Both tree shapes are provided as arguments to `Map` that will bind the transformations together creating a mapping:

```go
Map("test"
    Part(Obj{
        OffsetKey: Int2Str(Var("off")),
    }),
    // ->
    Part(Obj{
        "EndPosition": TypedObj("Position", Obj{
            "Offset": Var("off"),
        }),
    }),
)
```

You may also notice that we can swap shapes and automatically reverse the transformation. This is the benefit of using declarative DSL approach. This DSL could also be used without changes to build queries for AST. In fact, it already acts like query language to be able to find nodes with specific keys.

Another benefit of this approach is that each transformation step is very simple: check a specific key and push a value to stack, check that current value is an object \(or create it if we are constructing output tree\), convert value and store it as a variable. By combining these small steps we can express arbitrary complex transformations. Helpers that combine multiple transformations can be written to keep DSL high-level enough.

## Specification

Transformations will use schema-less representation, as described in [previous BIP](bip5-shape-based-ast-transformations.md).

All transformations will act on a shared state, represented by new `State` object. It will store all variables that were encountered during transform nodes execution and will use Go call stack to store references to current nodes that are being processed.

`Map` will accept two `Op` trees: one for a read pass \(check an input, verify that the tree shape is correct, store variables\) and second one for a write pass \(create output shape, use variables and constraints to populate fields\).

`Op` interface will describe two methods: one for applying forward pass and one for a reverse pass:

```go
type Sel interface{
    Check(st *State, n Node) (bool, error)
}

type Mod interface{
    Construct(st *State, n Node) (Node, error)
}

type Op interface{
    Sel
    Mod
}
```

Each `Op` can manipulate `State` to store any information that is necessary.

`Check` will examine the tree, save variables, and return if this node matches the shape or not. It may also return an error if shape matches, but there is a problem with converting values, checking enums, etc. `Check` should not modify the tree. Variable can be assigned only once - this ensures that code will not use same variable names in different transformation parts by mistake. If a variable is already assigned but is being checked by another transformation node, the value of a variable is compared. If the value is different, it's a runtime error.

`Construct` will use variables to modify the tree to match the target shape. It will return a new node that should be used instead of old one. `Construct` should not change any variables in the `State`.

`Op` should be strictly reversible: account for nil array vs empty arrays, nil, empty and missing fields, etc.

If steps provided by SDK cannot express the desired transformation, driver can define its own implementation of `Op` that follows the same rules. This will necessary for custom line number calculations, custom value conversion, etc.

### Proposed transformation steps

* `Is(val Value)` - checks if the current node is a primitive and is equal to a given value. Reversal changes the type of the node to primitive and assigns given value to the node.
* `Var(name string)` - stores current node as a value to a named variable in the shared state. Reversal replaces current node with the one from named variable. Variables are not primitives and can move subtrees.
* `AnyNode(apply Mod)` - matches any node and throws it away. Reversal will create a node with apply op.
* `And(ops ...Op)` - checks current node with all ops and fails if any of them fails. Reversal applies all modifications from ops to the current node. Typed ops should be at the beginning of the list to make sure that `Construct` creates a correct node type before applying specific changes to it.
* `Not(s Sel)` - only implements `Sel` interface, and succeeds when it's argument check fails. It should also check that no new variables were assigned inside sub op, since it makes no sense to extract variables from `Not` branch.
* `Check(sel []Sel, op Op)` - checks all sel arguments before executing op's check. Reversal only executes apply of op. Mostly used for `Not`.
* `Obj map[string]Op` - verifies that the current node is an object and checks fields with provided ops. Reversal changes node type to object and applies a provided operation to each field. This operation will populate a list of unprocessed keys for current object, so the transformation code can verify that transform was complete.
* `Arr(elems ...Op)` - checks if the current object is an array with a number of elements matching ops number, and applies ops to corresponding elements. Apply creates an array of the size that matches the number of ops and creates each element with the corresponding op.
* `Part(var string, obj Obj)` - stores all unprocessed fields of `obj` to a shared variable. A reversal step will create the same fields in the destination object. This step will also clear a list of unprocessed fields for the current node.
* `Lookup(var string, m map[Value]Value)` - uses a value of current node to find a replacement for it in the map and stores result as a variable. The reverse step will use a reverse map to lookup value stored in a variable and will assign it to the current node. Since reversal transformation needs to build a reverse map, the mapping should not be ambiguous in reverse direction.

### Proposed helpers

* `String(v string)` = `Is(v)` - shorthand for making a string value
* `AnyVal(def Value)` = `Any(Is(def))` - accept any value, create with default
* `One(op Op)` = `Arr(op)` - array with one element
* `HasType(typ string)` = `Has("@type", typ)` - a shorthand for checking type field
* `TypedObj(typ string, obj Obj)` = `Obj{KeyType: String(typ), obj...}` - a shorthand for an object with a specific type

### Examples

We will describe a usage of proposed DSL for existing transformations.

#### InternalTypeKey \(`@type`\)

```go
Map("test"
    Part("other", Obj{
        InternalTypeKey: Var("typ"),
    }),
    // ->
    Part("other", Obj{
        "@type": Var("typ"),
    }),
)
```

#### Token \(`@token`\)

Since the new transformation engine requires all fields to be described, a user needs to explicitly define each AST type and it's corresponding token.

Previously it was possible to specify two fields of one node that will become tokens, and it might become an error condition. With the new approach, this can be noticed earlier, because each type is defined separately.

Here is a small example for transforming Go comment node tokens:

```go
Map("test",
    Obj{
        "type": String("go:Comment"),
        "Text": Var("comment"),
    },
    // ->
    TypedObj("uast:Comment", Obj{
        "@token": Var("comment"),
    }),
)
```

#### OffsetKey

`OffsetKey` is used to save start position of AST node. Previously, the transformation was only able to attach positional information to node that had this key in it.

With a new approach, it's possible to move this information to any field in resulting subtree.

For example, we can use comment example and use its slash position as offset, stored in `@start` field of `Pos` object, inside `@pos` field of current object.

```go
Map("test",
    Obj{
        "type": String("go:Comment"),
        "Text": Var("comment"),
        "Slash": Var("off"),
    },
    // ->
    TypedObj("uast:Comment", Obj{
        "@token": Var("comment"),
        "@pos": TypedObj("uast:Pos", Obj{
            "@start": Obj{
                // store offset as a start position
                "@off": Var("off"),
            },
            "@end": Obj{
                // calculate as start+len(val) later
                "@off": AnyVal(-1),
            },
        }),
    }),
)
```

Note that there is no line/column information here. It can be computed later by applying custom transformation step.

The same approach can be used to describe `EndOffsetKey`, `LineKey`, etc by providing a helper with source field name and destination field name.

#### Annotation DSL

Although annotation DSL is mostly similar to new DSL, it still has some distinction that will require drivers rewrite:

* `Descendants` was used to traverse all nodes, but now `Map` will walk all nodes by default.
* `Descendants` was used to skip parts of AST structure to get to internal children. This will not be supported in new DSL since it makes it impossible to apply a reverse transformation, nor to verify separate transformation steps.
* `HasInternalRole(k), ...` has used instead of `Has(k, ...)` - it was listing children with specific field instead of listing children on a specific parent's field. This will not be supported in new DSL and should already be solved in previous BIP.

Mapping of old selectors to new DSL:

* `Self` -&gt; `And` or `Obj`.
* `Children` -&gt; Not supported directly, although `Children(HasInternalRole(k), ...)` can be mapped to `Obj{ k: ... }`.
* `Descendants` -&gt; Not supported. Was used mostly for iterating over all nodes, that will be done automatically now. Any other usages should specify what fields to traverse explicitly.
* `DescendantsOrSelf` -&gt; Not supported, see `Descendants`.
* `Roles(...)` -&gt; `Obj{ "@role": Arr(...) }`.

Mapping of old predicates to new DSL:

* `HasInternalType(t)` -&gt; `Obj{ "@type": t }`.
* `HasProperty(k, v)` -&gt; `Obj{ k: v }`.
* `HasInternalRole(t)` -&gt; Not supported directly, see `Children` selector.
* `HasChild(op)` -&gt; Not supported directly, see `Children` selector.
* `HasToken(v)` -&gt; `Obj{ "@token": v }`.
* `And(ops)` -&gt; `And(ops)`.
* `Not(op), ...` -&gt; `Check(Not(op), ...)`.
* `Or(ops)` -&gt; Was used to tag multiple types with a single role. Should be split into separate rules for each type to provide reversibility.

And here is an example of applying role annotations via new DSL:

```go
Map(
    Obj{
        "type": String("go:Comment"),
        "Text": Var("comment"),
        "Slash": Var("off"),
    },
    // ->
    TypedObj("uast:Comment", Obj{
        "@token": Var("comment"),
        "@role": Arr(
            CommentRole,
        )),
    }),
)
```

## Alternatives

* Use current approach and refactor the code for reusability. This will lead to splitting the code into small reusable parts, which will be mostly similar to proposed DSL. The biggest disadvantage of an old approach is that it will require to hand-craft all value assignments between old and new trees, the creation of nodes, etc. Also, an old approach cannot provide reversibility \(that is useful for verifying high-level UAST\), and cannot be used later as a form of a query language.
* Proposed approach requires all fields to be either mentioned as a constraint \(`Is`, `String`, etc\), or stored into a variable \(`Var`\). `Part` was proposed as a way to store all unused keys. It is possible to do this automatically and preserve everything that was not touched by transformations defined by the user. We think that this approach will only lead to more programming errors. In proposed solution user is forced to describe what exactly are the fields in the object \(assuming `Part` is not used\). Thus, it will act like a type-safe DSL, allowing to strictly check if user's expectations of AST are correct or not.
* It is also possible to omit reversibility requirement for transformation steps. Although it might sound like it will make the code simpler, in fact, transformations like `Obj` will just be split as a source transformation `HasKeys` and destination transformation `PutKeys`, since both are required to properly define a final transformation. Thus, it's easier to make a single transform with both forward and reverse steps. It's worth to note that sometimes it might be necessary to define a non-reversible transformation that may drop some data \(thus making it non-reversible\). Such transformations can just set an error state on the transformer, or panic in case they were called to make a reverse pass.

## Impact

The change requires SDK to switch to schema-less representation internally. Babelfish server should still return an old Node representation to maintain compatibility with current clients.

It's still possible to preserve limited compatibility for `ObjectToNode`, but it may be better to split it into separate transforms that drivers will use. And since we are targeting high-level UAST, all drivers will need to be changed to use new transformations anyway.

`Positioner` can be rewritten to new DSL, assuming that it will be able to access source code from transformation framework.

Annotation DSL will require only minor changes to match new DSL requirements, assuming previous BIP is implemented. Although the change is minimal, drivers will be forced to use DSL differently - it was common to omit intermediate fields while annotating objects, and new DSL requires all fields to be defined explicitly. Again, since high-level UAST will require a driver rewrite, annotations will be added on top of new AST-to-UAST transformations.

