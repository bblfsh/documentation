# BIP3: Agglutinative Roles language

| Field | Value |
| --- | --- |
| BIP | 3 |
| Title | Agglutinative Roles language |
| Author | Alfredo Beaumont |
| Status | Accepted |
| Created | 2017-09-04 |
| Updated | 2017-09-12 |

## Abstract

Current Role set includes roles that contain multiple properties, instead of having a single Role per property.

We propose modifying current role set to include only roles that represent a single property and to consider the Role set an agglutinative language, annotating each node with a set of Roles to fully express node properties.

We show how this can make the Role vocabulary more scalable, flexible and expressive, and how it can make the work of code analysts easier.

## Rationale

In current Role set many roles contain multiple properties \(i. e.: `OpPreIncrement` for `++i`\).

This presents some issues:

* It doesn't scale well.

  For a set of N properties, in the worst case, 2^N roles would be needed.

* It makes code analysis hard.

  For example, to search for all the operators you need to search for the union of many roles.

* There's hard to provide partial information for rare properties,

  if some property is not common enough as to deserve its own role.

A more flexible and scalable way to define the role set would be to have one role per property and to allow to combine these roles/properties as needed. In the example above `OpPreIncrement` would become the union of 3 roles \(`Operator`, `Increment`, `Prefix`\).

This combination of roles is already done to some extent, since a preincrement operator would actually be annotated with 2 roles: `Expression`, `OpPreIncrement`. The current proposal just deepens this property separation into roles.

As a derivative advantage of this new role language and its scalability, new roles could be defined that better capture some additional property shared by different tokens that will make code analysis even easier.

Following the previous example, we could use some additional roles, like `Arithmetic`, `Unary`, etc. to express additional properties of the role. So the final annotation could be something like: `Operator`, `Arithmetic`, `Unary`

* `Expression` \(or `Statement`, depending on the language\)
* `Operator`
* `Arithmetic`
* `Increment`
* `Unary`
* `Prefix`

With this kind of classification a code analyst that may be interested in analysing the usage of, say, arithmetic operators, would have an easier way to filter the `UAST` to find the interesting nodes.

Additionally, this agglutination of roles makes unsupported node types degrade more gracefully. For example, a `**` \(pow\) operator, which currently lacks a specific role, would be just left as: `Expression`, `Incomplete`. With the new language, the node could still retain most of the information:

* `Expression`
* `Operator`
* `Arithmetic`
* `Binary`

So can extract that it's a binary arithmetic operator generically, restricting the need to directly access the token to only those analysis interested in this kind of tokens specifically.

## Specification

The set of Roles are changed by this proposal. It's limited to partition the multiple property roles currently defined, leaving the potential addition of new property roles \(`Arithmetic`, `Comparison`, `Loop`, ...\) to a future BIP.

The following roles are proposed for removal:

* `SimpleIdentifier`
* `QualifiedIdentifier`
* `BinaryExpression`
* `BinaryExpressionLeft`
* `BinaryExpressionRight`
* `BinaryExpressionOp`
* `OpBitwiseLeftShift`
* `OpBitwiseRightShift`
* `OpBitwiseUnsignedRightShift`
* `OpBitwiseOr`
* `OpBitwiseXor`
* `OpBitwiseAnd`
* `OpEqual`
* `OpNotEqual`
* `OpLessThan`
* `OpLessThanEqual`
* `OpGreaterThan`
* `OpGreaterThanEqual`
* `OpSame`
* `OpNotSame`
* `OpContains`
* `OpNotContains`
* `OpPreIncrement`
* `OpPostIncrement`
* `OpPreDecrement`
* `OpPostDecrement`
* `OpNegative`
* `OpPositive`
* `OpBitwiseComplement`
* `OpDereference`
* `OpTakeAddress`
* `OpBooleanAnd`
* `OpBooleanOr`
* `OpBooleanNot`
* `OpBooleanXor`
* `OpAdd`
* `OpSubstract`
* `OpMultiply`
* `OpDivide`
* `OpMod`
* `PackageDeclaration`
* `ImportDeclaration`
* `ImportPath`
* `ImportAlias`
* `FunctionDeclaration`
* `FunctionDeclarationBody`
* `FunctionDeclarationName`
* `FunctionDeclarationReceiver`
* `FunctionDeclarationArgument`
* `FunctionDeclarationArgumentName`
* `FunctionDeclarationArgumentDefaultValue`
* `FunctionDeclarationVarArgsList`
* `TypeDeclaration`
* `TypeDeclarationBody`
* `TypeDeclarationBases`
* `TypeDeclarationImplements`
* `VisibleFromInstance`
* `VisibleFromType`
* `VisibleFromSubtype`
* `VisibleFromPackage`
* `VisibleFromSubpackage`
* `VisibleFromModule`
* `VisibleFromFriend`
* `VisibleFromWorld`
* `IfCondition`
* `IfBody`
* `IfElse`
* `SwitchCase`
* `SwitchCaseCondition`
* `SwitchCaseBody`
* `SwitchDefault`
* `ForInit`
* `ForExpression`
* `ForUpdate`
* `ForBody`
* `ForEach`
* `WhileCondition`
* `WhileBody`
* `DoWhileCondition`
* `DoWhileBody`
* `TryBody`
* `TryCatch`
* `TryFinally`
* `CallReceiver`
* `CallCallee`
* `CallPositionalArgument`
* `CallNamedArgument`
* `CallNamedArgumentName`
* `CallNamedArgumentValue`
* `BooleanLiteral`
* `ByteLiteral`
* `ByteStringLiteral`
* `CharacterLiteral`
* `ListLiteral`
* `MapLiteral`
* `NullLiteral`
* `NumberLiteral`
* `RegexpLiteral`
* `SetLiteral`
* `StringLiteral`
* `TupleLiteral`
* `TypeLiteral`
* `OtherLiteral`
* `MapEntry`
* `MapKey`
* `MapValue`
* `PrimitiveType`
* `AssignmentVariable`
* `AssignmentValue`
* `AugmentedAssignment`
* `AugmentedAssignmentOperator`
* `AugmentedAssignmentVariable`
* `AugmentedAssignmentValue`

The following roles are proposed for inclusion:

* `Identifier`
* `Qualified`
* `Operator`
* `Binary`
* `Unary`
* `Left`
* `Right`
* `Bitwise`
* `Boolean`
* `Unsigned`
* `LeftShift`
* `RightShift`
* `Or`
* `Xor`
* `And`
* `Equal`
* `Not`
* `LessThan`
* `LessThanOrEqual`
* `GreaterThan`
* `GreaterThanOrEqual`
* `Identical`
* `Contains`
* `Increment`
* `Decrement`
* `Negative`
* `Positive`
* `Dereference`
* `TakeAddress`
* `Add`
* `Substract`
* `Multiply`
* `Divide`
* `Modulo`
* `Package`
* `Declaration`
* `Import`
* `Path`
* `Alias`
* `Function`
* `Body`
* `Receiver`
* `Argument`
* `Name`
* `Value`
* `ArgsList`
* `Base`
* `Implements`
* `Visibility`
* `Instance`
* `Subtype`
* `Package`
* `Subpackage`
* `Module`
* `Friend`
* `World`
* `Condition`
* `Then`
* `Else`
* `Case`
* `Default`
* `Initialization`
* `Update`
* `Iterator`
* `Catch`
* `Finally`
* `Positional`
* `Byte`
* `ByteString`
* `Character`
* `List`
* `Map`
* `Null`
* `Number`
* `Regexp`
* `Set`
* `String`
* `Tuple`
* `Other`
* `Entry`
* `Key`
* `Primitive`

## Alternatives

Two possible alternatives would be:

* Use current language, that is, discard this proposal.

  This alternative would show the problems and limitations already commented in the _Rationale_ section.

* Use some kind of Role categorization.

  With this proposal, the number of roles per node will increase,

  and not all properties may have the same importance,

  some of them may be more important than others.

  For example, in the preincrement operator example used above,

  `Operator` role may be more relevant than `Prefix`,

  we may consider the first role a `noun`

  and the second and `adjective`,

  which doesn't make sense by itself,

  but can only go with a `noun`.

  We consider that this may be a valid approach,

  but it can be left for consideration in the future,

  after more experimentation,

  since it'd make the proposal more complex while not having a clear gain

  and it's easy to extend in that direction in the future.

## Impact

Incompatible changes to the Role set are proposed, in order to do that, the following changes would be needed:

* Versioning should be added to the [SDK](https://github.com/bblfsh/sdk/),

  to allow existing server and drivers work with a previous version of the SDK.

* Roles should be updated in the SDK.
* Protobuf generated code should be updated for [server](https://github.com/bblfsh/server) and [Python](https://github.com/bblfsh/client-python) and [Go](https://github.com/bblfsh/client-go) clients
* Update [libuast](https://github.com/bblfsh/libuast)'s role generator and role set.
* Existing drivers should be modified to support the new roles.

For the last point, a proposal of how to map existing roles into the new set of roles follows:

* `SimpleIdentifier`: `Identifier`
* `QualifiedIdentifier`: `Identifier`, `Qualified`
* `BinaryExpression`: `Expression`, `Binary`
* `BinaryExpressionLeft`: `Expression`, `Binary`, `Left`
* `BinaryExpressionRight`: `Expression`, `Binary`, `Right,`
* `BinaryExpressionOp`: `Expresion`, `Binary`, `Operator`
* `Infix`: `Infix`
* `Postfix`: `Postfix`
* `OpBitwiseLeftShift`: `Operator`, `Binary`, `Bitwise`, `LeftShift`
* `OpBitwiseRightShift`: `Operator`, `Binary`, `Bitwise`, `RightShift`
* `OpBitwiseUnsignedRightShift`: `Operator`, `Binary`, `Bitwise`, `RightShift`, `Unsigned`
* `OpBitwiseOr`: `Operator`, `Binary`, `Bitwise`, `Or`
* `OpBitwiseXor`: `Operator`, `Binary`, `Bitwise`, `Xor`
* `OpBitwiseAnd`: `Operator`, `Binary`, `Bitwise`, `And`
* `Expression`: `Expression`
* `Statement`: `Statement`
* `OpEqual`: `Operator`, `Binary`, `Equal`
* `OpNotEqual`: `Operator`, `Binary`, `Equal`, `Not`
* `OpLessThan`: `Operator`, `Binary`, `LessThan`
* `OpLessThanEqual`: `Operator`, `Binary`, `LessThanOrEqual`
* `OpGreaterThan`: `Operator`, `Binary`, `GreaterThan`
* `OpGreaterThanEqual`: `Operator`, `Binary`, `GreaterThanOrEqual`
* `OpSame`: `Operator`, `Binary`, `Identical`
* `OpNotSame`: `Operator`, `Binary`, `Identical`, `Not`
* `OpContains`: `Operator`, `Binary`, `Contains`
* `OpNotContains`: `Operator`, `Binary`, `Contains`, `Not`
* `OpPreIncrement`: `Operator`, `Unary`, `Increment`, `Prefix`
* `OpPostIncrement`: `Operator`, `Unary`, `Increment`, `Postfix`
* `OpPreDecrement`: `Operator`, `Unary`, `Decrement`, `Prefix`
* `OpPostDecrement`: `Operator`, `Unary`, `Decrement`, `Postfix`
* `OpNegative`: `Operator`, `Unary`, `Negative`
* `OpPositive`: `Operator`, `Unary`, `Positive`
* `OpBitwiseComplement`: `Operator`, `Unary`, `Bitwise`, `Not`
* `OpDereference`: `Operator`, `Unary`, `Dereference`
* `OpTakeAddress`: `Operator`, `Unary`, `TakeAddress`
* `File`: `File`
* `OpBooleanAnd`: `Operator`, `Binary`, `Boolean`, `Or`
* `OpBooleanOr`: `Operator`, `Binary`, `Boolean`, `Or`
* `OpBooleanNot`: `Operator`, `Binary`, `Boolean`, `Not`
* `OpBooleanXor`: `Operator`, `Binary`, `Boolean`, `Xor`
* `OpAdd`: `Operator`, `Binary`, `Add`
* `OpSubstract`: `Operator`, `Binary`, `Substract`
* `OpMultiply`: `Operator`, `Binary`, `Multiply`
* `OpDivide`: `Operator`, `Binary`, `Divide`
* `OpMod`: `Operator`, `Binary`, `Modulo`
* `PackageDeclaration`: `Package`, `Declaration`
* `ImportDeclaration`: `Import`, `Declaration`
* `ImportPath`: `Import`, `Path`
* `ImportAlias`: `Import`, `Alias`
* `FunctionDeclaration`: `Function`, `Declaration`
* `FunctionDeclarationBody`: `Function`, `Declaration`, `Body`
* `FunctionDeclarationName`: `Function`, `Declaration`, `Identifier`
* `FunctionDeclarationReceiver`: `Function`, `Declaration`, `Receiver`
* `FunctionDeclarationArgument`: `Function`, `Declaration`, `Argument`
* `FunctionDeclarationArgumentName`: `Function`, `Declaration`, `Argument`, `Name`
* `FunctionDeclarationArgumentDefaultValue`: `Function`, `Declaration`, `Argument`, `Value`
* `FunctionDeclarationVarArgsList`: `Function`, `Declaration`, `Argument`, `ArgsList`
* `TypeDeclaration`: `Type`, `Declaration`
* `TypeDeclarationBody`: `Type`, `Declaration`, `Body`
* `TypeDeclarationBases`: `Type`, `Declaration`, `Base`
* `TypeDeclarationImplements`: `Type`, `Declaration`, `Implements`
* `VisibleFromInstance`: `Visibility`, `Instance`
* `VisibleFromType`: `Visibility`, `Type`
* `VisibleFromSubtype`: `Visibility`, `Subtype`
* `VisibleFromPackage`: `Visibility`, `Package`
* `VisibleFromSubpackage`: `Visibility`, `Subpackage`
* `VisibleFromModule`: `Visibility`, `Module`
* `VisibleFromFriend`: `Visibility`, `Friend`
* `VisibleFromWorld`: `Visibility`, `World`
* `If`: `If`
* `IfCondition`: `If`, `Condition`
* `IfBody`: `If`, `Body`, `Then`
* `IfElse`: `If`, `Body`, `Else`
* `Switch`: `Switch`
* `SwitchCase`: `Switch`, `Case`
* `SwitchCaseCondition`: `Swich`, `Case`, `Condition,`
* `SwitchCaseBody`: `Switch`, `Case`, `Body`
* `SwitchDefault`: `Switch`, `Case`, `Default`
* `For`: `For`
* `ForInit`: `For`, `Initialization`
* `ForExpression`: `For`, `Condition`
* `ForUpdate`: `For`, `Update`
* `ForBody`: `For`, `Body`
* `ForEach`: `For`, `Iterator`
* `While`: `While`
* `WhileCondition`: `While`, `Condition`
* `WhileBody`: `While`, `Body`
* `DoWhile`: `DoWhile`
* `DoWhileCondition`: `DoWhile`, `Condition`
* `DoWhileBody`: `DoWhile`, `Body`
* `Break`: `Break`
* `Continue`: `Continue`
* `Goto`: `Goto`
* `Block`: `Block`
* `BlockScope`: `BlockScope`
* `Return`: `Return`
* `Try`: `Try`
* `TryBody`: `Try`, `Body`
* `TryCatch`: `Try`, `Catch`
* `TryFinally`: `Try`, `Finally`
* `Throw`: `Throw`
* `Assert`: `Assert`
* `Call`: `Function`, `Call`
* `CallReceiver`: `Call`, `Receiver`
* `CallCallee`: `Call`, `Callee`
* `CallPositionalArgument`: `Call`, `Argument`, `Positional`
* `CallNamedArgument`: `Call`, `Argument`
* `CallNamedArgumentName`: `Call`, `Argument`, `Name`
* `CallNamedArgumentValue`: `Call`, `Argument`, `Value`
* `Noop`: `Noop`
* `BooleanLiteral`: `Literal`, `Boolean`
* `ByteLiteral`: `Literal`, `Byte`
* `ByteStringLiteral`: `Literal`, `ByteString`
* `CharacterLiteral`: `Literal`, `Character`
* `ListLiteral`: `Literal`, `List`
* `MapLiteral`: `Literal`, `Map`
* `NullLiteral`: `Literal`, `Null`
* `NumberLiteral`: `Literal`, `Number`
* `RegexpLiteral`: `Literal`, `Regexp`
* `SetLiteral`: `Literal`, `Set`
* `StringLiteral`: `Literal`, `String`
* `TupleLiteral`: `Literal`, `Tuple`
* `TypeLiteral`: `Literal`, `Type`
* `OtherLiteral`: `Literal`
* `MapEntry`: `Map`, `Entry`
* `MapKey`: `Map`, `Key`
* `MapValue`: `Map`, `Value`
* `Type`: `Type`
* `PrimitiveType`: `Type`, `Primitive`
* `Assignment`: `Operator`, `Binary`, `Assignment`
* `AssignmentVariable`: `Binary`, `Left`
* `AssignmentValue`: `Binary`, `Right`
* `AugmentedAssignment`: `Operator`, `Binary`, `Assignment`
* `AugmentedAssignmentOperator`: `Operator`, `Binary`, \[`OperatorRole`\]
* `AugmentedAssignmentVariable`: `Left`
* `AugmentedAssignmentValue`: `Right`
* `This`: `This`
* `Comment`: `Comment`
* `Documentation`: `Documentation`
* `Whitespace`: `Whitespace`
* `Incomplete`: `Incomplete`
* `Unannotated`: `Unannotated`

