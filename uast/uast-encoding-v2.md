# UAST v2 binary encoding

The UAST binary encoding uses [Protocol Buffers](https://developers.google.com/protocol-buffers/)
to encode individual messages, but requires additional logic to reconstruct the UAST.

The design document can be found [here](https://docs.google.com/document/d/1iDOuWLnTckBdSscUMWGl2hBqAAqRn8jdRabXZSdr2aw/edit?usp=sharing)
and the reference implementation in Go can be found in the [Babelfish SDK](https://github.com/bblfsh/sdk/blob/b44cc02214bf0d9c55dc3b830f423a6df6709ca9/uast/nodes/nodesproto/nodesproto.go).

On a high level, a binary UAST uses [UAST v2 data model](./representation-v2.md).
It extends this model by adding an ephemeral node ID field used to link UAST nodes.

The file consist of two sections:

* A [header](#header), containing a file format marker ("magic"), format version and other options.
* A flat [list of nodes](#nodes) linked by node IDs.

The list contains zero or more graph nodes addressed by ID. Binary format version `1` must
only store tree structures, but the format was made compatible with more general graph structures
for future extensions.

The ID of the root tree node is stored in the [Header](#header) message.

## Encoding

### Messages

Each protocol buffer message in the binary format is length-prefixed, similar to `bytes` values used in protocol buffers fields.

First, a message size is encoded as an [unsigned varint](https://developers.google.com/protocol-buffers/docs/encoding),
followed by `N` bytes of the message itself.

There are two messages used to encode the UAST:

- [GraphHeader](https://github.com/bblfsh/sdk/blob/b44cc02214bf0d9c55dc3b830f423a6df6709ca9/uast/nodes/nodesproto/nodes.proto#L17)
- [Node](https://github.com/bblfsh/sdk/blob/b44cc02214bf0d9c55dc3b830f423a6df6709ca9/uast/nodes/nodesproto/nodes.proto#L46)

See the sections below for the usage of those messages.

### Header

The binary UAST file begins with a `\x00bgr` (`0x00 0x62 0x67 0x72`) "magic" that can be
used to detect a format of the file.

It is followed by a version number encoded as an unsigned 32 bit integer in little-endian byte order.
The current format version is `1`.

Next, a length-prefixed `GraphHeader` message follows:
```proto
message GraphHeader {
    uint64 last_id  = 1;
    uint64 root     = 2;
    uint64 metadata = 3;
}
```

`last_id` is an optional field that sets a maximal node ID reserved by the node ID allocator.
If not set, `max(node.id)` is assumed (maximal node ID in the [nodes list](#nodes-list)).
Implementation may reserve more IDs than required by the UAST nodes by setting `last_id > max(node.id)`.

`root` is an ID of the tree root node. Multiple unnamed roots can be stored by referencing an
[Array](#array) node by its ID, and multiple named roots can be stored by referencing an
[Object](#object) node. If not set explicitly, implementations should search the graph for
unreferenced nodes of type [Array](#array) or [Object](#object) and treat them as an array of roots.

`metadata` is an optional ID of a file metadata root, similar to `root`. This tree may
contain arbitrary tree structure that describes the UAST file. If not set, no metadata is
assumed. Implementation may ignore this tree, except for a case when `root` is not explicitly.

### Nodes list

Following the [Header](#header) message is a list of zero or more length-prefixed
[Node](https://github.com/bblfsh/sdk/blob/b44cc02214bf0d9c55dc3b830f423a6df6709ca9/uast/nodes/nodesproto/nodes.proto#L46)
messages, consuming the remainder of the file. There is no explicit node count provided by the format.
If a corrupt node is encountered (e.g., length prefix bigger than remaining data, or node message invalid),
this must result in a decoding error.

### Node

Each Node is described by the following protocol buffer message:
```proto
message Node {
    uint64 id = 1;
    
    oneof value {
        string string = 2;
        int64  int    = 3;
        uint64 uint   = 4;
        double float  = 5;
        bool   bool   = 6;
    }
    
    repeated uint64 keys = 7;
    uint64 keys_from = 10;
    
    repeated uint64 values = 8;
    
    bool is_object = 9;
    
    uint64 values_offs = 11;
}
```

The `id` field is unique, monotonically increasing positive value. ID value is used in most message
fields to reference other nodes.

If `id` field is omitted, the `prevNode.ID + 1` is assumed. The first node with no `id`
is assumed to have an ID value of `1`.

*Note:* Nil is a special node distinct from an empty object or array, used to distinguish a pointer to
an absent node from an absent pointer to a node. They are represented by ID value `0`.

All other nodes can be of one of 3 kinds: [Values](#values), [Arrays](#array) and [Objects](#object).

#### Values

Values nodes must not have any other fields except `id` (optional) and `value` "oneof" field:

```proto
message Node {
    uint64 id = 1;
    
    oneof value {
        string string = 2;
        int64  int    = 3;
        uint64 uint   = 4;
        double float  = 5;
        bool   bool   = 6;
    }
}
```

Each "oneof" value corresponds to a value type defined in the [UAST v2 representation](./representation-v2.md#values).

Value nodes may be deduplicated when encoding the UAST, thus implementations should make a
copy of the value when decoding.

#### Array

Array nodes must only have an `id` (optional), `values` and `values_offs` fields set.
`is_object` field may also be set to `false` when encoding an empty array.

```proto
message Node {
    uint64 id = 1;
    
    repeated uint64 values = 8;
    uint64 values_offs = 11;
    
    bool is_object = 9;
}
```

`values` field contains a list of node IDs of the children nodes in the array. The list order
corresponds to an array element order. Nil elements are encoded as ID value `0`.

`values_offs` is an optional node ID offset that should be added to all `values` elements
before resolving that child ID. Used for compression.

Empty array are encoded by setting `id` field (optional) and optionally setting `is_object`
to `false`. Message with no fields set is a valid empty array. `values` may be also set to
an empty array and/or the `values_offs` may be set to `0`.

#### Object

Object nodes are represented by two arrays: one for object keys (`keys` and `keys_from`)
and the second one for values (`values` and `values_offs`).

```proto
message Node {
    uint64 id = 1;
    
    repeated uint64 keys = 7;
    uint64 keys_from = 10;
    
    repeated uint64 values = 8;
    
    bool is_object = 9;
    
    uint64 values_offs = 11;
}
```

Empty objects must have a `is_object` fields set. Other fields may be either set to zero value,
or must not be set at all.

Keys array can be either specified explicitly with `keys`, or copied from another object
by specifying its ID in `keys_from`. Both `keys` and `keys_from` must not be set.

Values are decoded the same way as for [Arrays](#array). The size of values array must be
the same as arrays of keys.

Objects are considered to be a unordered sets, thus keys and values may be reordered for efficiency.

### Tree reconstruction

The state needed to decode the file consist of a `current_id` integer (set to `1` initially)
used for automatic `id` field generation when it's not set and a map of node IDs to [Node](#node)
structures.

It must also keep a set of seen nodes (`seen`) when reconstructing the tree to prevent loops.
It must also keep a flag for each entry to this set to track if a given node is a tree branch,
as opposed to a direct acyclic graph (DAG) branch. Nodes in tree branches are referenced
only once, as opposed to DAGs where branches may overlap.

Decoding process is can be the following:

- The decoder reads each node message from the file in sequence.
- It assigns an `id` automatically from `current_id` and then incrementing the `current_id`,
  or sets `current_id` to `id + 1` if the `id` field was set in the message.
- It must ensure that `id` values are monotonically increasing, including both auto-assigned
  and IDs specified in the messages. Non monotonically increasing IDs or duplicate IDs must
  result in a decoding error.
- It must ensure that ID is unique by checking it against the node map.
- If `value` "oneof" is set, assume the node is a primitive value, add it to a map and read
  the next message.
- If `keys`, `keys_from` and `values` are either not set or zero, check the `is_object` field
  to distinguish between empty array (`is_object` is `false` or not set) and empty objects
  (`is_object` is `true`). Add an empty object or array to the map and read the next message.
- If `keys` or `keys_from` is set, consider the node as an [Object](#object). Both fields must not be set.
    - If `keys_from` is set, use it as an ID and lookup a node in the map. The node should
      be an [Object](#object) and it must precede the message with `keys_from`. Copy keys
      from that object and use them as if `keys` field was set.
    - The size of `values` array must be exactly the same as `keys` array. If `values_offs`
      is set, add it to all elements of `values` array. IDs should not be resolved yet.
    - Each key from `keys` array corresponds to a single node ID from `values`. Implementation
      may store those KV pairs in a more suitable data structure.
    - Implementation may store the node in a list to resolve it later.
- If `keys` and `keys_from` are not set, and `values` is set, consider the node as an [Array](#array).
    - Decode the values array the same way as in Object by using the `values` and `values_offs`.
      IDs should not be resolved yet.
    - Implementation may store the node in a list to resolve it later.
- When all nodes were read from the file, resolve all [Object](#object) and [Array](#array)
  elements.
    - Each element in `keys` array must not be zero and must correspond to a String value
      in the nodes map. There must not be any duplicate string values in a single Object.
    - The `values` array for both node kinds may contain zero IDs that must be resolved to
      nil nodes, or must contain a non-zero ID that corresponds to a node in the nodes map.
      Resolved values must be checked against a `seen` node set to prevent loops and DAGs.
      Value nodes must not be inserted to this set because they may be deduplicated and
      share the same ID.
    - Implementation may also track the list of unreferenced objects in case the `root`
      was not specified explicitly in the header.
- If `root` was specified, lookup its ID in the nodes map. It must correspond to an
  [Object](#object), an [Array](#array) or be zero.
- If `root` is zero, use the unreferenced nodes list to create a new [Array](#array) that
  will be a root. `metadata` ID should be excluded from this array. Implementation may also
  require a `root` to always be set and trow an error in case it's not. Note, that `root`
  may be zero in the case of empty UAST as well. In this case, no nodes except ones referenced
  by `metadata` will be present in the file.
- Implementation may optionally use `metadata` to reconstruct the file metadata tree.
  `metadata` tree must not be the same as `root`.