# UAST Querying

The clients support querying over the returned UAST to filter for specific nodes using a client-specific `filter()` method.

The querying uses [xpath syntax](https://www.w3.org/TR/xpath/) and it can be really useful to extract features from the code and, thanks to the annotation system implemented by babelfish, it can be done in a universal way.

Any of the [Node](https://godoc.org/gopkg.in/bblfsh/sdk.v2/uast#Node) fields can be used for querying in the following way:

* `@type` is converted to the element name
* `@token`, if available, is converted to an attribute with `@token` as keyword and the actual token as value
* Every `Role` is converted to an attribute with `@role` as keyword and the role name as value.
* Every `Property` is converted to an attribute with the property keyword as keyword and the property value as value
* `@pos`, if available, is mapped to a node with a `start` and `end` properties, each containing a node of `Position` type which in turn will have the `offset`, `line` and `col` properties.

Internally, these are mapped in the XML node in the following way:

```markup
<{{@type}}
    @token='{{Token}}'
    {{for role in Roles}}
    @role{{role}}
    {{for key, value in Properties}}
    {{key}}='{{value}}'
    @pos{{start-offset, start-line, start-col, end-offset, end-line, end-col}}
    {{Children}}
</{{@type}}>
```

This means that both language specific queries \(InternalType, Properties\) and language agnostic queries \(Roles\) can be done.

## Example Queries

* All the numeric literals in Python: `//Num`
* All the numeric literals in ANY language: `//*[@role='Number' and @role='Literal']`
* All the integer literals in Python: `//*[@NumType='int']`

The query language also allows some more complex queries:

* All the simple identifiers: `//*[@role='Identifier' and not(@role='Qualified')]` or, if you're getting the semantic UAST (which is the default): `//Identifier[not(@role='Qualified')]`.
* All the simple identifiers that don't have any positioning: `//*[@role='Identifier' and not(@role='Qualified') and not(@startOffset) and not(@endOffset)]`
* All the arguments in function calls: `//*[@role='Call' and @role='Argument']`
* The parent node of a call argument: `//*@[role='Call' and @role='Argument']/parent::*` (you can also specify the type of the parent node instead of an `*` or use `ancestor::type` to get the first ancestor of a specified type).
* All the numeric literals in binary arithmetic operators: `//*[@role='Binary' and @role='Operator' and @role='Arithmetic']//*[@role='Number' and @role='Literal']`

[XPath 1.0 functions](https://developer.mozilla.org/en-US/docs/Web/XPath/Functions) can also be used in the queries, but if the query returns a type different than the default node list you must use one of the specific typed functions: `filter_bool` \(or `filterBool` in some clients\), `filter_number` or `filter_string`. If you use the wrong type the error will tell you what is the type returned by the expression.

