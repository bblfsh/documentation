# UAST Query Language

The clients support querying over the returned UAST to filter for specific
nodes using a client-specific `filter()` method.

The querying uses [xpath syntax](https://www.w3.org/TR/xpath/) and it can be
really useful to extract features from the code and, thanks to the annotation
system implemented by babelfish, it can be done in a universal way.

Any of the [node](https://godoc.org/gopkg.in/bblfsh/sdk.v2/uast#Node) fields can
be used for querying in the following way:

- `InternalType` is converted to the element name
- `Token`, if available, is converted to an attribute with `token` as keyword and the actual token as value
- Every `Role` is converted to an attribute concatenating a `role` prefix and the role name in CamelCase.
- Every `Property` is converted to an attribute with the property keyword as keyword and the property value as value
- `StartPosition`, if available, is mapped to three attributes:
  - A `startOffset` attribute, with the offset as value
  - A `startLine` attribute, with the line as value
  - A `startCol` attribute, with the column as value
- `EndPosition`, if available, is mapped to three attributes:
  - A `endOffset` attribute, with the offset as value
  - A `endLine` attribute, with the line as value
  - A `endCol` attribute, with the column as value

Internally, these are mapped in the XML node in the following way:

```xml
<{{InternalType}}
    token='{{Token}}'
    {{for role in Roles}}
    role{{role}}
    {{for key, value in Properties}}
    {{key}}='{{value}}'
    startOffset={{StartPosition.Offset}}
    startLine={{StartPosition.Line}}
    startCol={{StartPosition.Col}}
    endOffset={{EndPosition.Offset}}
    endLine={{EndPosition.Line}}
    endCol={{EndPosition.Col}}>
    {{Children}}
</{{InternalType}}>
```

This means that both language specific queries (InternalType, Properties) and language agnostic queries (Roles) can be done. For example:

- All the numeric literals in Python: `//Num`
- All the numeric literals in ANY language: `//*[@roleNumber and @roleLiteral]`
- All the integer literals in Python: `//*[@NumType='int']`

The query language also allow some more complex queries:

- All the elements in the tree that have either start or end offsets: `//*[@startOffset or @endOffset]`
- All the simple identifiers: `//*[@roleIdentifier and not(@roleQualified)]`
- All the simple identifiers that don't have any positioning: `//*[@roleIdentifier and not(@roleQualified) and not(@startOffset) and not(@endOffset)]`
- All the arguments in function calls: `//*[@roleCall and @roleArgument]`
- All the numeric literals in binary arithmetic operators: `//*[@roleBinary and @roleOperator and @roleArithmetic]//*[@roleNumber and @roleLiteral]`

[XPath 1.0 functions](https://developer.mozilla.org/en-US/docs/Web/XPath/Functions) can also be used in the queries, but
if the query returns a type different than the default node list you must use
one of the specific typed functions: `filter_bool` (or `filterBool` in some
clients), `filter_number` or `filter_string`. If you use the wrong type the
error will tell you what is the type returned by the expression.
