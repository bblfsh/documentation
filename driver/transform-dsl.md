# Transformation DSL

Transformation DSL is a library that provides functions for transforming
AST between different representations and shapes (native, UAST, custom).

Transformations using this DSL usually consist of two steps: source
transform and target transform. This pair is called "mapping" in SDK.
Since both sides of transform are of the same type, target and destination
can be switched to create a reverse mapping (reverse transformation).

```golang
Map("test",

    // source tree
    Obj{
        // Check that the "type" field in the map have the "Ident" value
        "type": String("Ident"),

        // Store the value of the "name" property in the variable "x"
        "name": Var("x"),
    },

    // target tree (destination)
    Obj{
        // Set the "type" property to the string "go:Indent"
        "type": String("go:Ident"),

        // Retrieve the variable "x" saved above and set the "token" property to its value
        "token": Var("x"),
    },
)
```

Each part consists of multiple simple operations that are combined together
to describe a final transformation (tree shape).

## Operations

Operations may run in two modes: `Check` and `Construct`.

The first mode checks if a specific AST node matches the operation constraints.
Every AST node has an associated state that operates like a storage and
can pass values between the `Check` and `Construct`. Thus, if the `Check`
passes, an operation may store additional information to the state
associated with this AST node (more on this below).

Here is an example of an operation described in the `Check` context.

```golang
Obj{
    "type": String("Ident"),

    // store field value to a state variable "x"
    "name": Var("x"),
},
```

The `Construct` mode uses the state associated with AST node to reconstruct
original node.

These two modes are executed sequentially to transform an AST:

```golang
// create a new state for this AST node
s := NewState()

// verify that an AST tree matches all constraints of this transform
if ok, _ := op.Check(s, ast); ok {
    // construct a new AST using stored state (will match original AST)
    newAST, _ := op.Construct(s, nil)
}

s.Reset()

// check/apply next transform
```

### Examples

Let's consider a few examples. We will use the transformation defined above
over this source AST node expressed in json format:

```json
[
    {"type":"Ident", "name": null, "offset": 5}
]
```

Running a `Check` directly on this AST will fail, since the root is
an array, while an operation requires it to be an object (`Obj` operation,
which would correspond to a json Object or dictionary in the example above).
Operations always match only the root node without recursively traversing
to children. `Mapping` (the enclosing Map call in the example above)
on the other hand, runs operations for all nodes recursively.

Let's assume we run an operation for an element of that array.

```json
{"type":"Ident", "name": null, "offset": 5}
```

The first operation (the `Obj` with the enclosing braces in the example)
will succeed and proceed to run the `Check`s on the fields.

```golang
"type": String("Ident"),
```

This field check will succeed, since the node contains a "type" field and
it indeed contains an "Ident" string.

```golang
"name": Var("x"),
```

Since the "name" field exists in the object, its null value will be stored
as a state variable "x". Note that the DSL makes a distinction between null
fields (optional values) and fields that does not exist (optional fields).
Variables can store any node type, even if its value is null, thus it will
be stored in the node state:

```json
{
    "x": null
}
```

These are the only constraints that were defined for this transform, but
we still have an "offset" field left. At this point `Check` will fail
because the "offset" field was not handled by any transform. This usually
means that the AST shape was changed unexpectedly, or the developer forgot
to describe this field. Both cases indicate an error that must be addressed
(see the [Completeness](#completeness) section for more details).

But we can define our transform as partial, which will automatically
store all unhandled field to the node state, which is useful if we want
to change only specific fields:

```golang
Map("test",
    Part("other", Obj{ ..source.. }),
    // ->
    Part("other", Obj{ ..target.. }),
)
```

Note that `Part` operation should be paired inside `Map` because each of
them produces state variables for other to consume.

With `Part`, `Check` method will succeed and produce the following node state:

```json
{
    "x": null,
    "other": {"offset": 5}
}
```

Now, we will reconstruct an AST tree by using the `Construct` method and
the state for this node that we just received:

```
newAST, _ := op.Construct(s, nil)
```

In this case process runs in reverse - each operation uses state to
reconstruct AST nodes based on defined constraints.

First, an `Obj` operation will construct an empty AST node:

```json
{}
```

Then, it will start to populate fields:

```golang
"type": String("Ident"),
```

For this field, everything is strictly defined and no state is required
to reconstruct the field - it will simply assign a "type" to "Ident".

```json
{
    "type": "Ident"
}
```

For the "name" field, the state is required to reconstruct it:

```golang
"name": Var("x"),
```

At this point, the state variable "x" will be used to reconstruct the
field value (and assign null to it).

Remember that variables can store any node type: it might have been
a string, an array or another AST node.

```json
{
    "type": "Ident",
    "name": null
}
```

Since there is no more constraints defined, the `Part` transform will use
it's state stored in the "other" variable to populate all unhandled fields:

```json
{
    "type": "Ident",
    "name": null,
    "offset": 5
}
```

As you can see, an operation can be used to describe a tree shape and
reconstruct it back, assuming that all placeholder variables used in
an operation are defined in the node state.

### Mappings

Since mapping defines two operations (source and target), we can change
our code to use one operation for checking the tree and building
the state and a different operation for constructing it:

```golang
A := Part("other", Obj{
    "type": String("Ident"),
    "name": Var("x"),
})

B := Part("other", Obj{
    "type": String("go:Ident"),
    "token": Var("x"),
})

s := NewState()

// verify an AST with operation A, populate state
if ok, _ := A.Check(s, ast); ok {
    // construct a new AST using the state and operation B
    newAST, _ := B.Construct(s, nil)
}
```

As long as all variables defined in A exactly match all variables defined
in B, the `Check`-`Construct` routine will work in the same way as
described above.

Since the state only stores placeholder values, while the operation defines
the shape of the tree, any complex shape transformation can be defined:

```golang
Obj{
    "type": String("Binary"),
    "op": Var("op"),
    "vals": Arr( Var("left"), Var("right") ),
}

// state: {"op": ..., "left": ..., "right": ...}

Obj{
    "type": String("Binary"),
    "op": Obj{
        "type": String("Operation"),
        "token": Var("op")
    },
    "left": Var("left"),
    "right": Var("right"),
}
```

This example will transform this AST (matches the first transform):

```json
{
    "type": "Binary",
    "op": "+",
    "vals": [
        {"type":"Ident", "name": "v"},
        5
    ]
}
```

Into a new AST (defined by the second transform):

```json
{
    "type": "Binary",
    "op": {
        "type": "Operation",
        "token": "+"
    },
    "left": {"type":"Ident", "name": "v"},
    "right": 5
}
```


## Restrictions

The DSL also enforces some restrictions on transformations and operations.

### Completeness

All operations should either:

* Reject the tree if it does not match one of constraints.

* Describe the whole tree, and save unknown/placeholder nodes to the state.

This rule ensures that, if an operation matches the tree, it would be able
to replicate it from a state associated with the node since all fields
are covered by it.

This property allows to test each transformation by executing the
`Check`-`Construct` routine, or detect if a mapping defined by a driver's
developer ignores certain AST fields that are seen in the wild
(which would result in an error).

### Reversibility

Since all operations are required to implement both `Check` and `Construct`
they can be used interchangeably as either source and/or target transforms.

For a simple case, this means, that given a node T over which we apply
a transform A, we can produce a state S (with `Check`) and then reconstruct
the original node T (with `Construct`) by only using the state S.

```
S := A.Check(T); A.Construct(S) == T
```

In a more complete example, it's possible for the transform A to check
the tree T, which produces the state S and for transform B to use this
state to construct a new tree T'. At the same time the opposite is true:
transform B can be used to produce state S from T' for A to reconstruct
original tree T.

```
// forward transform
S  := A.Check(T)
T' := B.Construct(S)

// reverse transform
S' := B.Check(T')
S == S'
T == A.Construct(S')
```

## List of operations

### Values

#### Any value

If any value can match the transformation, it should be stored into
a named variable: `Var("x")`. Any name can be used for the variable as
long as it is unique in this transformation.

When writing mappings, a variable with the same name should exist in the
second transformation.

#### Equal

To check value for exact match, `Is(v)` operation is used.

The SDK defines helpers for common types: `String(v)`, `Int(v)`, etc.

Null values can be checked with `Is(nil)`.

#### Equal or nil

To check any node for nil or specific rule `Opt` is used:

```
Opt("is_nil", Is(42))
```

`Opt` needs a named variable to store it's state ("is_nil" in this case).

#### Value in

To check that a value is equal to any value in the list, two operations
should be executed:

1) Check that the value is in the list.

2) Store the value into a variable.

To express this, the `Check` helper can be used:

```golang
Check(
    // condition
    In(
        uast.String("A"),
        uast.String("B"),
    ),

    // action
    Var("x"),
)
```

#### Not equal, not in, etc

To check value is not equal `Check` can be combined with `Not`:

```golang
Check(
    // condition
    Not(String("A")),

    // store the value
    Var("x"),
)
```

Note that the value can still be of any type since we're using `Var` and
thus the type constraint of `String` inside the `Check` will not affect
the stored variable

In the same way, a value can be checked for a "not in" constraint:

```golang
Check(
    Not(In(
            uast.String("A"),
            uast.String("B"),
    )),
    // store the value
    Var("x"),
)
```

### Objects

#### Empty object

To represent an empty object, `Obj{}` can be used.

#### Object or nil

To check any node for nil or specific rule `Opt` is used:

```
Opt("is_nil", Obj{ ... })
```

`Opt` needs a named variable to store it's state ("is_nil" in this case).

#### All fields must exist

For object that always have a specific set of fields, the `Obj` operation
can be used:

```golang
Obj{
    "type": String("Ident"),
    "name": Var("name"),
}
```

Since this operation is represented by a map, the order of operations is
not preserved. Operations will be executed following an alphabetical order
from the field name. This might affect operations that requires a variable
to be defined, such as `Lookup` and `If`. To control the execution order,
use `Fields`.

#### Specific field options

To get more control about the ordering and other aspects of the
field's checks, use `Fields`:

```golang
Fields{
    {Name: "type", Op: String("If")},
    {Name: "then": Op: Var("then")},
    {Name: "else": Op: Var("then"), Optional: "else_exists"},
}
```

For each optional field, a variable should be provided to store the state.
Optional field means that they may not exist. For checking nil fields see
`Opt` operation.

#### Skip unknown fields

`Part` can be used to skip unknown fields in an object:

```golang
Part("other", Obj{
    "name": Var("x"),
})
```

`Part` will store all unhandled fields in the `other` variable. It usually
requires to have a `Part` on both sides of a mapping.

### Arrays

#### Fixed number of elements

`Arr` can be used to represent an array with a specific number of elements:

```golang
Arr(
    // element 0, fixed - must be an int with a value of 42
    Int(42),

    // element 1, any type - saved to a variable
    Var("x"),
)
```

A helper is available for arrays with exactly one element: `One(Int(42))`.

#### Any number of elements

To apply an operation to each array element, use `Each`:

```golang
Each("elems",
    Obj{
        "name": Var("x"),
    },
)
```

Note that `Each` requires a named variable to store theelements.
This also means that `Each` should always be on both sides of a mapping.

`Each` allows an array to be empty or nil.

The operation that is passed to `Each` will receive a clean state, thus
it will not be able to access any variables from the parent. It also
means that variable names inside this operation will not collide with
the variables defined by the parent transformation.

#### Any number of elements followed by a fixed "suffix"

To match arrays with a specific set of nodes at the end, use `Append`:

```golang
Append(
    Var("x"),
    Arr(Int(42)),
)
```

`Append` will verify that node is an array and will filter out "suffix"
nodes before passing it to primary operation (`Var("x")` in this case).

When used to costruct the tree, it will append "suffix" nodes to
the end of an array defined by sub-operation.

As in Go, `Append` will not affect variable "x".