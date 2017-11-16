# UAST Query Language

The clients support querying over the returned UAST to filter for specific
nodes using a client-specific `filter()` method.

The querying uses [xpath syntax](https://www.w3.org/TR/xpath/) and it can be
really useful to extract features from the code and thanks to the annotation
system implemented by babelfish, it can be done in a universal fashion.

Any of the [node](https://godoc.org/github.com/bblfsh/sdk/uast#Node) fields can
be used for querying, for example:

- All the numeric literals in Python: `//NumLiteral`
- All the numeric literals in ANY language: `//*[@roleNumber and @roleLiteral]`
- All the integer literals in Python: `//*[@NumType='int']`
- All the elements in the tree that have either start or end offsets: `//*[@startOffset or @endOffset]`
- All the simple identifiers: `//*[@roleIdentifier and not(@roleQualified)]`
- All the simple identifiers that don't have any positioning: `//*[@roleIdentifier and not(@roleQualified) and not(@startOffset) and not(@endOffset)]`
- All the arguments in function calls: `//*[@roleCall and @roleArgument]`
- All the numeric literals in binary arithmetic operators: `//*[@roleBinary and @roleOperator and @roleArithmetic]//*[@roleNumber and @roleLiteral]`
