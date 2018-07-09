# Roles

Role is the main UAST annotation. It indicates that a node in an AST can be interpreted as acting with certain language-independent role.

| Role | Go | Java | JavaScript | PHP | Python | Ruby |
| --- | --- | --- | --- | --- | --- | --- |
| [Add](roles.md#add) | ✓ | ✓ | ✓ | ✓ | ✓ | ✓ |
| [Alias](roles.md#alias) | ✓ |  |  | ✓ | ✓ | ✓ |
| [And](roles.md#and) | ✓ | ✓ | ✓ | ✓ | ✓ | ✓ |
| [Annotation](roles.md#annotation) |  | ✓ |  |  | ✓ |  |
| [Anonymous](roles.md#anonymous) |  | ✓ |  | ✓ | ✓ | ✓ |
| [ArgsList](roles.md#argslist) | ✓ | ✓ | ✓ | ✓ | ✓ |  |
| [Argument](roles.md#argument) | ✓ | ✓ | ✓ | ✓ | ✓ | ✓ |
| [Arithmetic](roles.md#arithmetic) | ✓ | ✓ | ✓ |  | ✓ | ✓ |
| [Assert](roles.md#assert) |  | ✓ |  |  | ✓ |  |
| [Assignment](roles.md#assignment) | ✓ | ✓ | ✓ | ✓ | ✓ | ✓ |
| [Base](roles.md#base) |  | ✓ | ✓ | ✓ | ✓ |  |
| [Binary](roles.md#binary) | ✓ | ✓ | ✓ | ✓ | ✓ | ✓ |
| [Bitwise](roles.md#bitwise) | ✓ | ✓ | ✓ | ✓ | ✓ | ✓ |
| [Block](roles.md#block) | ✓ | ✓ | ✓ | ✓ | ✓ | ✓ |
| [Body](roles.md#body) | ✓ | ✓ | ✓ | ✓ | ✓ | ✓ |
| [Boolean](roles.md#boolean) | ✓ | ✓ | ✓ | ✓ | ✓ | ✓ |
| [Break](roles.md#break) | ✓ | ✓ | ✓ | ✓ | ✓ | ✓ |
| [Byte](roles.md#byte) |  |  |  |  |  |  |
| [ByteString](roles.md#bytestring) |  |  |  |  | ✓ |  |
| [Call](roles.md#call) | ✓ | ✓ | ✓ | ✓ | ✓ | ✓ |
| [Callee](roles.md#callee) | ✓ | ✓ | ✓ |  | ✓ |  |
| [Case](roles.md#case) | ✓ | ✓ | ✓ | ✓ |  | ✓ |
| [Catch](roles.md#catch) |  | ✓ | ✓ | ✓ | ✓ | ✓ |
| [Character](roles.md#character) | ✓ | ✓ |  |  |  |  |
| [Comment](roles.md#comment) | ✓ | ✓ | ✓ | ✓ | ✓ |  |
| [Condition](roles.md#condition) | ✓ | ✓ | ✓ | ✓ | ✓ | ✓ |
| [Contains](roles.md#contains) |  |  |  |  | ✓ |  |
| [Continue](roles.md#continue) | ✓ |  | ✓ | ✓ | ✓ | ✓ |
| [Declaration](roles.md#declaration) | ✓ | ✓ | ✓ | ✓ | ✓ | ✓ |
| [Decrement](roles.md#decrement) | ✓ | ✓ | ✓ | ✓ |  |  |
| [Default](roles.md#default) |  | ✓ |  | ✓ |  | ✓ |
| [Dereference](roles.md#dereference) | ✓ |  |  |  |  |  |
| [Divide](roles.md#divide) | ✓ | ✓ | ✓ | ✓ | ✓ | ✓ |
| [DoWhile](roles.md#dowhile) |  | ✓ | ✓ | ✓ |  |  |
| [Documentation](roles.md#documentation) |  | ✓ |  | ✓ |  |  |
| [Else](roles.md#else) | ✓ | ✓ | ✓ | ✓ | ✓ | ✓ |
| [Entry](roles.md#entry) | ✓ |  |  | ✓ |  |  |
| [Enumeration](roles.md#enumeration) |  | ✓ |  |  |  |  |
| [Equal](roles.md#equal) | ✓ | ✓ | ✓ | ✓ | ✓ | ✓ |
| [Expression](roles.md#expression) | ✓ | ✓ | ✓ | ✓ | ✓ | ✓ |
| [File](roles.md#file) | ✓ | ✓ | ✓ | ✓ | ✓ | ✓ |
| [Finally](roles.md#finally) |  | ✓ | ✓ | ✓ | ✓ | ✓ |
| [For](roles.md#for) | ✓ | ✓ | ✓ | ✓ | ✓ | ✓ |
| [Friend](roles.md#friend) |  |  |  |  |  |  |
| [Function](roles.md#function) | ✓ | ✓ | ✓ | ✓ | ✓ | ✓ |
| [Goto](roles.md#goto) | ✓ |  |  | ✓ |  |  |
| [GreaterThan](roles.md#greaterthan) | ✓ | ✓ | ✓ | ✓ | ✓ |  |
| [GreaterThanOrEqual](roles.md#greaterthanorequal) | ✓ | ✓ | ✓ | ✓ | ✓ | ✓ |
| [Identical](roles.md#identical) |  |  | ✓ | ✓ | ✓ | ✓ |
| [Identifier](roles.md#identifier) | ✓ | ✓ | ✓ | ✓ | ✓ | ✓ |
| [If](roles.md#if) | ✓ | ✓ | ✓ | ✓ | ✓ | ✓ |
| [Implements](roles.md#implements) |  |  |  | ✓ |  |  |
| [Import](roles.md#import) | ✓ | ✓ | ✓ | ✓ | ✓ |  |
| [Incomplete](roles.md#incomplete) | ✓ | ✓ | ✓ | ✓ | ✓ | ✓ |
| [Increment](roles.md#increment) | ✓ | ✓ | ✓ | ✓ |  |  |
| [Infix](roles.md#infix) |  |  |  |  |  |  |
| [Initialization](roles.md#initialization) | ✓ | ✓ | ✓ | ✓ |  |  |
| [Instance](roles.md#instance) |  | ✓ | ✓ |  |  | ✓ |
| [Invalid](roles.md#invalid) |  |  |  |  |  |  |
| [Iterator](roles.md#iterator) | ✓ | ✓ | ✓ | ✓ | ✓ | ✓ |
| [Key](roles.md#key) | ✓ |  | ✓ |  | ✓ |  |
| [Left](roles.md#left) | ✓ | ✓ | ✓ | ✓ | ✓ | ✓ |
| [LeftShift](roles.md#leftshift) | ✓ | ✓ | ✓ | ✓ | ✓ | ✓ |
| [LessThan](roles.md#lessthan) | ✓ | ✓ | ✓ | ✓ | ✓ |  |
| [LessThanOrEqual](roles.md#lessthanorequal) | ✓ | ✓ | ✓ | ✓ | ✓ | ✓ |
| [List](roles.md#list) | ✓ | ✓ | ✓ | ✓ | ✓ | ✓ |
| [Literal](roles.md#literal) | ✓ | ✓ | ✓ | ✓ | ✓ | ✓ |
| [Map](roles.md#map) | ✓ |  | ✓ | ✓ | ✓ | ✓ |
| [Module](roles.md#module) |  |  | ✓ | ✓ | ✓ | ✓ |
| [Modulo](roles.md#modulo) | ✓ | ✓ | ✓ | ✓ | ✓ | ✓ |
| [Multiply](roles.md#multiply) | ✓ | ✓ | ✓ | ✓ | ✓ | ✓ |
| [Name](roles.md#name) | ✓ | ✓ | ✓ | ✓ | ✓ | ✓ |
| [Negative](roles.md#negative) | ✓ | ✓ | ✓ |  | ✓ |  |
| [Noop](roles.md#noop) |  |  |  | ✓ | ✓ |  |
| [Not](roles.md#not) | ✓ | ✓ | ✓ | ✓ | ✓ | ✓ |
| [Null](roles.md#null) |  | ✓ | ✓ | ✓ | ✓ | ✓ |
| [Number](roles.md#number) | ✓ | ✓ | ✓ | ✓ | ✓ | ✓ |
| [Operator](roles.md#operator) | ✓ | ✓ | ✓ | ✓ | ✓ | ✓ |
| [Or](roles.md#or) | ✓ | ✓ | ✓ | ✓ | ✓ | ✓ |
| [Package](roles.md#package) |  | ✓ |  |  |  |  |
| [Pathname](roles.md#pathname) | ✓ | ✓ | ✓ | ✓ | ✓ |  |
| [Positional](roles.md#positional) | ✓ | ✓ |  |  | ✓ |  |
| [Positive](roles.md#positive) | ✓ | ✓ | ✓ |  | ✓ |  |
| [Postfix](roles.md#postfix) |  | ✓ | ✓ | ✓ |  |  |
| [Primitive](roles.md#primitive) | ✓ | ✓ |  |  | ✓ | ✓ |
| [Qualified](roles.md#qualified) |  | ✓ | ✓ | ✓ | ✓ | ✓ |
| [Receiver](roles.md#receiver) | ✓ | ✓ |  | ✓ | ✓ |  |
| [Regexp](roles.md#regexp) |  |  | ✓ |  |  | ✓ |
| [Relational](roles.md#relational) | ✓ | ✓ | ✓ | ✓ | ✓ | ✓ |
| [Return](roles.md#return) | ✓ | ✓ | ✓ | ✓ | ✓ |  |
| [Right](roles.md#right) | ✓ | ✓ | ✓ | ✓ | ✓ | ✓ |
| [RightShift](roles.md#rightshift) | ✓ | ✓ | ✓ | ✓ | ✓ | ✓ |
| [Role](roles.md#role) | ✓ |  |  |  |  |  |
| [Scope](roles.md#scope) | ✓ | ✓ | ✓ |  | ✓ |  |
| [Set](roles.md#set) |  |  |  |  | ✓ |  |
| [Statement](roles.md#statement) | ✓ | ✓ | ✓ | ✓ | ✓ | ✓ |
| [String](roles.md#string) | ✓ | ✓ | ✓ | ✓ | ✓ | ✓ |
| [Subpackage](roles.md#subpackage) |  |  |  |  |  |  |
| [Substract](roles.md#substract) | ✓ | ✓ | ✓ | ✓ | ✓ | ✓ |
| [Subtype](roles.md#subtype) |  | ✓ |  |  |  |  |
| [Switch](roles.md#switch) | ✓ | ✓ | ✓ | ✓ |  | ✓ |
| [TakeAddress](roles.md#takeaddress) | ✓ |  |  |  |  |  |
| [Then](roles.md#then) | ✓ | ✓ | ✓ | ✓ | ✓ | ✓ |
| [This](roles.md#this) |  | ✓ | ✓ |  |  | ✓ |
| [Throw](roles.md#throw) |  | ✓ | ✓ | ✓ | ✓ | ✓ |
| [Try](roles.md#try) |  | ✓ | ✓ | ✓ | ✓ | ✓ |
| [Tuple](roles.md#tuple) |  |  |  |  | ✓ | ✓ |
| [Type](roles.md#type) | ✓ | ✓ | ✓ | ✓ | ✓ | ✓ |
| [Unannotated](roles.md#unannotated) |  |  |  |  |  |  |
| [Unary](roles.md#unary) | ✓ | ✓ | ✓ | ✓ | ✓ |  |
| [Unsigned](roles.md#unsigned) |  | ✓ | ✓ |  |  |  |
| [Update](roles.md#update) | ✓ | ✓ | ✓ | ✓ | ✓ | ✓ |
| [Value](roles.md#value) | ✓ |  | ✓ | ✓ | ✓ |  |
| [Variable](roles.md#variable) | ✓ | ✓ | ✓ | ✓ |  |  |
| [Visibility](roles.md#visibility) |  | ✓ | ✓ | ✓ | ✓ | ✓ |
| [While](roles.md#while) |  | ✓ | ✓ | ✓ | ✓ | ✓ |
| [Whitespace](roles.md#whitespace) |  |  |  |  | ✓ |  |
| [World](roles.md#world) |  | ✓ |  | ✓ | ✓ | ✓ |
| [Xor](roles.md#xor) | ✓ | ✓ | ✓ | ✓ | ✓ | ✓ |

## Add

Add is an arithmetic operator \(i.e. `+`.\)

**Supported by**: [_Go_](https://github.com/bblfsh/go-driver/blob/master/driver/normalizer/annotation.go), [_Java_](https://github.com/bblfsh/java-driver/blob/master/driver/normalizer/annotation.go), [_Javascript_](https://github.com/bblfsh/javascript-driver/blob/master/driver/normalizer/annotation.go), [_Php_](https://github.com/bblfsh/php-driver/blob/master/driver/normalizer/annotation.go), [_Python_](https://github.com/bblfsh/python-driver/blob/master/driver/normalizer/annotation.go), [_Ruby_](https://github.com/bblfsh/ruby-driver/blob/master/driver/normalizer/annotation.go)

## Alias

Alias is an alternative name for some construct.

**Supported by**: [_Go_](https://github.com/bblfsh/go-driver/blob/master/driver/normalizer/annotation.go), [_Php_](https://github.com/bblfsh/php-driver/blob/master/driver/normalizer/annotation.go), [_Python_](https://github.com/bblfsh/python-driver/blob/master/driver/normalizer/annotation.go), [_Ruby_](https://github.com/bblfsh/ruby-driver/blob/master/driver/normalizer/annotation.go)

## And

And is an AND operation \(i.e. `&&`, `&`, `and`, etc.\)

**Supported by**: [_Go_](https://github.com/bblfsh/go-driver/blob/master/driver/normalizer/annotation.go), [_Java_](https://github.com/bblfsh/java-driver/blob/master/driver/normalizer/annotation.go), [_Javascript_](https://github.com/bblfsh/javascript-driver/blob/master/driver/normalizer/annotation.go), [_Php_](https://github.com/bblfsh/php-driver/blob/master/driver/normalizer/annotation.go), [_Python_](https://github.com/bblfsh/python-driver/blob/master/driver/normalizer/annotation.go), [_Ruby_](https://github.com/bblfsh/ruby-driver/blob/master/driver/normalizer/annotation.go)

## Annotation

Annotation is syntactic metadata

**Supported by**: [_Java_](https://github.com/bblfsh/java-driver/blob/master/driver/normalizer/annotation.go), [_Python_](https://github.com/bblfsh/python-driver/blob/master/driver/normalizer/annotation.go)

## Anonymous

Anonymous is an unbound construct

**Supported by**: [_Java_](https://github.com/bblfsh/java-driver/blob/master/driver/normalizer/annotation.go), [_Php_](https://github.com/bblfsh/php-driver/blob/master/driver/normalizer/annotation.go), [_Python_](https://github.com/bblfsh/python-driver/blob/master/driver/normalizer/annotation.go), [_Ruby_](https://github.com/bblfsh/ruby-driver/blob/master/driver/normalizer/annotation.go)

## ArgsList

ArgsList is variable number of arguments \(i.e. `...`, `Object...`, `*args`, etc.\)

**Supported by**: [_Go_](https://github.com/bblfsh/go-driver/blob/master/driver/normalizer/annotation.go), [_Java_](https://github.com/bblfsh/java-driver/blob/master/driver/normalizer/annotation.go), [_Javascript_](https://github.com/bblfsh/javascript-driver/blob/master/driver/normalizer/annotation.go), [_Php_](https://github.com/bblfsh/php-driver/blob/master/driver/normalizer/annotation.go), [_Python_](https://github.com/bblfsh/python-driver/blob/master/driver/normalizer/annotation.go)

## Argument

Argument is variable used as input/output in a function.

**Supported by**: [_Go_](https://github.com/bblfsh/go-driver/blob/master/driver/normalizer/annotation.go), [_Java_](https://github.com/bblfsh/java-driver/blob/master/driver/normalizer/annotation.go), [_Javascript_](https://github.com/bblfsh/javascript-driver/blob/master/driver/normalizer/annotation.go), [_Php_](https://github.com/bblfsh/php-driver/blob/master/driver/normalizer/annotation.go), [_Python_](https://github.com/bblfsh/python-driver/blob/master/driver/normalizer/annotation.go), [_Ruby_](https://github.com/bblfsh/ruby-driver/blob/master/driver/normalizer/annotation.go)

## Arithmetic

Arithmetic is a type of operation

**Supported by**: [_Go_](https://github.com/bblfsh/go-driver/blob/master/driver/normalizer/annotation.go), [_Java_](https://github.com/bblfsh/java-driver/blob/master/driver/normalizer/annotation.go), [_Javascript_](https://github.com/bblfsh/javascript-driver/blob/master/driver/normalizer/annotation.go), [_Python_](https://github.com/bblfsh/python-driver/blob/master/driver/normalizer/annotation.go), [_Ruby_](https://github.com/bblfsh/ruby-driver/blob/master/driver/normalizer/annotation.go)

## Assert

Assert checks if an expression is true and if it is not, it signals an error/exception, possibly stopping the execution.

**Supported by**: [_Java_](https://github.com/bblfsh/java-driver/blob/master/driver/normalizer/annotation.go), [_Python_](https://github.com/bblfsh/python-driver/blob/master/driver/normalizer/annotation.go)

## Assignment

Assignment is an assignment operator.

**Supported by**: [_Go_](https://github.com/bblfsh/go-driver/blob/master/driver/normalizer/annotation.go), [_Java_](https://github.com/bblfsh/java-driver/blob/master/driver/normalizer/annotation.go), [_Javascript_](https://github.com/bblfsh/javascript-driver/blob/master/driver/normalizer/annotation.go), [_Php_](https://github.com/bblfsh/php-driver/blob/master/driver/normalizer/annotation.go), [_Python_](https://github.com/bblfsh/python-driver/blob/master/driver/normalizer/annotation.go), [_Ruby_](https://github.com/bblfsh/ruby-driver/blob/master/driver/normalizer/annotation.go)

## Base

Base is the parent type of which another type inherits.

**Supported by**: [_Java_](https://github.com/bblfsh/java-driver/blob/master/driver/normalizer/annotation.go), [_Javascript_](https://github.com/bblfsh/javascript-driver/blob/master/driver/normalizer/annotation.go), [_Php_](https://github.com/bblfsh/php-driver/blob/master/driver/normalizer/annotation.go), [_Python_](https://github.com/bblfsh/python-driver/blob/master/driver/normalizer/annotation.go)

## Binary

Binary is any form of binary operator, in contrast with unary operators.

**Supported by**: [_Go_](https://github.com/bblfsh/go-driver/blob/master/driver/normalizer/annotation.go), [_Java_](https://github.com/bblfsh/java-driver/blob/master/driver/normalizer/annotation.go), [_Javascript_](https://github.com/bblfsh/javascript-driver/blob/master/driver/normalizer/annotation.go), [_Php_](https://github.com/bblfsh/php-driver/blob/master/driver/normalizer/annotation.go), [_Python_](https://github.com/bblfsh/python-driver/blob/master/driver/normalizer/annotation.go), [_Ruby_](https://github.com/bblfsh/ruby-driver/blob/master/driver/normalizer/annotation.go)

## Bitwise

Bitwise is any form of bitwise operation.

**Supported by**: [_Go_](https://github.com/bblfsh/go-driver/blob/master/driver/normalizer/annotation.go), [_Java_](https://github.com/bblfsh/java-driver/blob/master/driver/normalizer/annotation.go), [_Javascript_](https://github.com/bblfsh/javascript-driver/blob/master/driver/normalizer/annotation.go), [_Php_](https://github.com/bblfsh/php-driver/blob/master/driver/normalizer/annotation.go), [_Python_](https://github.com/bblfsh/python-driver/blob/master/driver/normalizer/annotation.go), [_Ruby_](https://github.com/bblfsh/ruby-driver/blob/master/driver/normalizer/annotation.go)

## Block

Block is a group of statements. If the source language has block scope, it should be annotated both with Block and BlockScope.

**Supported by**: [_Go_](https://github.com/bblfsh/go-driver/blob/master/driver/normalizer/annotation.go), [_Java_](https://github.com/bblfsh/java-driver/blob/master/driver/normalizer/annotation.go), [_Javascript_](https://github.com/bblfsh/javascript-driver/blob/master/driver/normalizer/annotation.go), [_Php_](https://github.com/bblfsh/php-driver/blob/master/driver/normalizer/annotation.go), [_Python_](https://github.com/bblfsh/python-driver/blob/master/driver/normalizer/annotation.go), [_Ruby_](https://github.com/bblfsh/ruby-driver/blob/master/driver/normalizer/annotation.go)

## Body

Body is a sequence of instructions in a block.

**Supported by**: [_Go_](https://github.com/bblfsh/go-driver/blob/master/driver/normalizer/annotation.go), [_Java_](https://github.com/bblfsh/java-driver/blob/master/driver/normalizer/annotation.go), [_Javascript_](https://github.com/bblfsh/javascript-driver/blob/master/driver/normalizer/annotation.go), [_Php_](https://github.com/bblfsh/php-driver/blob/master/driver/normalizer/annotation.go), [_Python_](https://github.com/bblfsh/python-driver/blob/master/driver/normalizer/annotation.go), [_Ruby_](https://github.com/bblfsh/ruby-driver/blob/master/driver/normalizer/annotation.go)

## Boolean

Boolean is any form of boolean operation.

**Supported by**: [_Go_](https://github.com/bblfsh/go-driver/blob/master/driver/normalizer/annotation.go), [_Java_](https://github.com/bblfsh/java-driver/blob/master/driver/normalizer/annotation.go), [_Javascript_](https://github.com/bblfsh/javascript-driver/blob/master/driver/normalizer/annotation.go), [_Php_](https://github.com/bblfsh/php-driver/blob/master/driver/normalizer/annotation.go), [_Python_](https://github.com/bblfsh/python-driver/blob/master/driver/normalizer/annotation.go), [_Ruby_](https://github.com/bblfsh/ruby-driver/blob/master/driver/normalizer/annotation.go)

## Break

Break is a construct for early exiting a block.

**Supported by**: [_Go_](https://github.com/bblfsh/go-driver/blob/master/driver/normalizer/annotation.go), [_Java_](https://github.com/bblfsh/java-driver/blob/master/driver/normalizer/annotation.go), [_Javascript_](https://github.com/bblfsh/javascript-driver/blob/master/driver/normalizer/annotation.go), [_Php_](https://github.com/bblfsh/php-driver/blob/master/driver/normalizer/annotation.go), [_Python_](https://github.com/bblfsh/python-driver/blob/master/driver/normalizer/annotation.go), [_Ruby_](https://github.com/bblfsh/ruby-driver/blob/master/driver/normalizer/annotation.go)

## Byte

Byte is a single-byte element.

**Supported by**:

## ByteString

ByteString is a raw byte string.

**Supported by**: [_Python_](https://github.com/bblfsh/python-driver/blob/master/driver/normalizer/annotation.go)

## Call

Call is any call, whether it is a function, procedure, method or macro. In its simplest form, a call will have a single child with a function name \(callee\). Arguments are marked with Argument and Positional or Name. In OO languages there is usually a Receiver too.

**Supported by**: [_Go_](https://github.com/bblfsh/go-driver/blob/master/driver/normalizer/annotation.go), [_Java_](https://github.com/bblfsh/java-driver/blob/master/driver/normalizer/annotation.go), [_Javascript_](https://github.com/bblfsh/javascript-driver/blob/master/driver/normalizer/annotation.go), [_Php_](https://github.com/bblfsh/php-driver/blob/master/driver/normalizer/annotation.go), [_Python_](https://github.com/bblfsh/python-driver/blob/master/driver/normalizer/annotation.go), [_Ruby_](https://github.com/bblfsh/ruby-driver/blob/master/driver/normalizer/annotation.go)

## Callee

Callee is the callable being called. It might be the name of a function or procedure, it might be a method, it might a simple name or qualified with a namespace.

**Supported by**: [_Go_](https://github.com/bblfsh/go-driver/blob/master/driver/normalizer/annotation.go), [_Java_](https://github.com/bblfsh/java-driver/blob/master/driver/normalizer/annotation.go), [_Javascript_](https://github.com/bblfsh/javascript-driver/blob/master/driver/normalizer/annotation.go), [_Python_](https://github.com/bblfsh/python-driver/blob/master/driver/normalizer/annotation.go)

## Case

Case is a clause whose expression is compared with the condition.

**Supported by**: [_Go_](https://github.com/bblfsh/go-driver/blob/master/driver/normalizer/annotation.go), [_Java_](https://github.com/bblfsh/java-driver/blob/master/driver/normalizer/annotation.go), [_Javascript_](https://github.com/bblfsh/javascript-driver/blob/master/driver/normalizer/annotation.go), [_Php_](https://github.com/bblfsh/php-driver/blob/master/driver/normalizer/annotation.go), [_Ruby_](https://github.com/bblfsh/ruby-driver/blob/master/driver/normalizer/annotation.go)

## Catch

Catch is a clause to capture exceptions.

**Supported by**: [_Java_](https://github.com/bblfsh/java-driver/blob/master/driver/normalizer/annotation.go), [_Javascript_](https://github.com/bblfsh/javascript-driver/blob/master/driver/normalizer/annotation.go), [_Php_](https://github.com/bblfsh/php-driver/blob/master/driver/normalizer/annotation.go), [_Python_](https://github.com/bblfsh/python-driver/blob/master/driver/normalizer/annotation.go), [_Ruby_](https://github.com/bblfsh/ruby-driver/blob/master/driver/normalizer/annotation.go)

## Character

Character is an encoded character.

**Supported by**: [_Go_](https://github.com/bblfsh/go-driver/blob/master/driver/normalizer/annotation.go), [_Java_](https://github.com/bblfsh/java-driver/blob/master/driver/normalizer/annotation.go)

## Comment

Comment is a code comment.

**Supported by**: [_Go_](https://github.com/bblfsh/go-driver/blob/master/driver/normalizer/annotation.go), [_Java_](https://github.com/bblfsh/java-driver/blob/master/driver/normalizer/annotation.go), [_Javascript_](https://github.com/bblfsh/javascript-driver/blob/master/driver/normalizer/annotation.go), [_Php_](https://github.com/bblfsh/php-driver/blob/master/driver/normalizer/annotation.go), [_Python_](https://github.com/bblfsh/python-driver/blob/master/driver/normalizer/annotation.go)

## Condition

Condition is a condition in an IfStatement or IfExpression.

**Supported by**: [_Go_](https://github.com/bblfsh/go-driver/blob/master/driver/normalizer/annotation.go), [_Java_](https://github.com/bblfsh/java-driver/blob/master/driver/normalizer/annotation.go), [_Javascript_](https://github.com/bblfsh/javascript-driver/blob/master/driver/normalizer/annotation.go), [_Php_](https://github.com/bblfsh/php-driver/blob/master/driver/normalizer/annotation.go), [_Python_](https://github.com/bblfsh/python-driver/blob/master/driver/normalizer/annotation.go), [_Ruby_](https://github.com/bblfsh/ruby-driver/blob/master/driver/normalizer/annotation.go)

## Contains

Contains is a membership predicate that checks if the lhs value is a member of the rhs container \(i.e. `in` in Python.\)

**Supported by**: [_Python_](https://github.com/bblfsh/python-driver/blob/master/driver/normalizer/annotation.go)

## Continue

Continue is a construct for continuation with the next iteration of a loop.

**Supported by**: [_Go_](https://github.com/bblfsh/go-driver/blob/master/driver/normalizer/annotation.go), [_Javascript_](https://github.com/bblfsh/javascript-driver/blob/master/driver/normalizer/annotation.go), [_Php_](https://github.com/bblfsh/php-driver/blob/master/driver/normalizer/annotation.go), [_Python_](https://github.com/bblfsh/python-driver/blob/master/driver/normalizer/annotation.go), [_Ruby_](https://github.com/bblfsh/ruby-driver/blob/master/driver/normalizer/annotation.go)

## Declaration

Declaration is a construct to specify properties of an identifier.

**Supported by**: [_Go_](https://github.com/bblfsh/go-driver/blob/master/driver/normalizer/annotation.go), [_Java_](https://github.com/bblfsh/java-driver/blob/master/driver/normalizer/annotation.go), [_Javascript_](https://github.com/bblfsh/javascript-driver/blob/master/driver/normalizer/annotation.go), [_Php_](https://github.com/bblfsh/php-driver/blob/master/driver/normalizer/annotation.go), [_Python_](https://github.com/bblfsh/python-driver/blob/master/driver/normalizer/annotation.go), [_Ruby_](https://github.com/bblfsh/ruby-driver/blob/master/driver/normalizer/annotation.go)

## Decrement

Decrement is an arithmetic operator that decrements a value \(i. e. `--i`.\)

**Supported by**: [_Go_](https://github.com/bblfsh/go-driver/blob/master/driver/normalizer/annotation.go), [_Java_](https://github.com/bblfsh/java-driver/blob/master/driver/normalizer/annotation.go), [_Javascript_](https://github.com/bblfsh/javascript-driver/blob/master/driver/normalizer/annotation.go), [_Php_](https://github.com/bblfsh/php-driver/blob/master/driver/normalizer/annotation.go)

## Default

Default is a clause that is called when no other clause is matches.

**Supported by**: [_Java_](https://github.com/bblfsh/java-driver/blob/master/driver/normalizer/annotation.go), [_Php_](https://github.com/bblfsh/php-driver/blob/master/driver/normalizer/annotation.go), [_Ruby_](https://github.com/bblfsh/ruby-driver/blob/master/driver/normalizer/annotation.go)

## Dereference

Dereference is an operation that gets the actual value of a pointer or reference \(i.e. `*x`.\)

**Supported by**: [_Go_](https://github.com/bblfsh/go-driver/blob/master/driver/normalizer/annotation.go)

## Divide

Divide is an arithmetic operator \(i.e. `/`.\)

**Supported by**: [_Go_](https://github.com/bblfsh/go-driver/blob/master/driver/normalizer/annotation.go), [_Java_](https://github.com/bblfsh/java-driver/blob/master/driver/normalizer/annotation.go), [_Javascript_](https://github.com/bblfsh/javascript-driver/blob/master/driver/normalizer/annotation.go), [_Php_](https://github.com/bblfsh/php-driver/blob/master/driver/normalizer/annotation.go), [_Python_](https://github.com/bblfsh/python-driver/blob/master/driver/normalizer/annotation.go), [_Ruby_](https://github.com/bblfsh/ruby-driver/blob/master/driver/normalizer/annotation.go)

## DoWhile

DoWhile is a loop construct with a body and a condition.

**Supported by**: [_Java_](https://github.com/bblfsh/java-driver/blob/master/driver/normalizer/annotation.go), [_Javascript_](https://github.com/bblfsh/javascript-driver/blob/master/driver/normalizer/annotation.go), [_Php_](https://github.com/bblfsh/php-driver/blob/master/driver/normalizer/annotation.go)

## Documentation

Documentation is a node that represents documentation of another node, such as function or package. Documentation is usually in the form of a string in certain position \(e.g. Python docstring\) or comment \(e.g. Javadoc, godoc\).

**Supported by**: [_Java_](https://github.com/bblfsh/java-driver/blob/master/driver/normalizer/annotation.go), [_Php_](https://github.com/bblfsh/php-driver/blob/master/driver/normalizer/annotation.go)

## Else

Else is the clause executed when the Condition is false.

**Supported by**: [_Go_](https://github.com/bblfsh/go-driver/blob/master/driver/normalizer/annotation.go), [_Java_](https://github.com/bblfsh/java-driver/blob/master/driver/normalizer/annotation.go), [_Javascript_](https://github.com/bblfsh/javascript-driver/blob/master/driver/normalizer/annotation.go), [_Php_](https://github.com/bblfsh/php-driver/blob/master/driver/normalizer/annotation.go), [_Python_](https://github.com/bblfsh/python-driver/blob/master/driver/normalizer/annotation.go), [_Ruby_](https://github.com/bblfsh/ruby-driver/blob/master/driver/normalizer/annotation.go)

## Entry

Entry is a collection element.

**Supported by**: [_Go_](https://github.com/bblfsh/go-driver/blob/master/driver/normalizer/annotation.go), [_Php_](https://github.com/bblfsh/php-driver/blob/master/driver/normalizer/annotation.go)

## Enumeration

Enumeration is a distinct type that represents a set of named constants

**Supported by**: [_Java_](https://github.com/bblfsh/java-driver/blob/master/driver/normalizer/annotation.go)

## Equal

Equal is an equality predicate \(i.e. `=`, `==`, etc.\)

**Supported by**: [_Go_](https://github.com/bblfsh/go-driver/blob/master/driver/normalizer/annotation.go), [_Java_](https://github.com/bblfsh/java-driver/blob/master/driver/normalizer/annotation.go), [_Javascript_](https://github.com/bblfsh/javascript-driver/blob/master/driver/normalizer/annotation.go), [_Php_](https://github.com/bblfsh/php-driver/blob/master/driver/normalizer/annotation.go), [_Python_](https://github.com/bblfsh/python-driver/blob/master/driver/normalizer/annotation.go), [_Ruby_](https://github.com/bblfsh/ruby-driver/blob/master/driver/normalizer/annotation.go)

## Expression

Expression is a construct computed to produce some value.

**Supported by**: [_Go_](https://github.com/bblfsh/go-driver/blob/master/driver/normalizer/annotation.go), [_Java_](https://github.com/bblfsh/java-driver/blob/master/driver/normalizer/annotation.go), [_Javascript_](https://github.com/bblfsh/javascript-driver/blob/master/driver/normalizer/annotation.go), [_Php_](https://github.com/bblfsh/php-driver/blob/master/driver/normalizer/annotation.go), [_Python_](https://github.com/bblfsh/python-driver/blob/master/driver/normalizer/annotation.go), [_Ruby_](https://github.com/bblfsh/ruby-driver/blob/master/driver/normalizer/annotation.go)

## File

File is the root node of a single file AST.

**Supported by**: [_Go_](https://github.com/bblfsh/go-driver/blob/master/driver/normalizer/annotation.go), [_Java_](https://github.com/bblfsh/java-driver/blob/master/driver/normalizer/annotation.go), [_Javascript_](https://github.com/bblfsh/javascript-driver/blob/master/driver/normalizer/annotation.go), [_Php_](https://github.com/bblfsh/php-driver/blob/master/driver/normalizer/annotation.go), [_Python_](https://github.com/bblfsh/python-driver/blob/master/driver/normalizer/annotation.go), [_Ruby_](https://github.com/bblfsh/ruby-driver/blob/master/driver/normalizer/annotation.go)

## Finally

Finally is a clause for a block executed after a block with exception handling.

**Supported by**: [_Java_](https://github.com/bblfsh/java-driver/blob/master/driver/normalizer/annotation.go), [_Javascript_](https://github.com/bblfsh/javascript-driver/blob/master/driver/normalizer/annotation.go), [_Php_](https://github.com/bblfsh/php-driver/blob/master/driver/normalizer/annotation.go), [_Python_](https://github.com/bblfsh/python-driver/blob/master/driver/normalizer/annotation.go), [_Ruby_](https://github.com/bblfsh/ruby-driver/blob/master/driver/normalizer/annotation.go)

## For

For is a loop with an initialization, a condition, an update and a body.

**Supported by**: [_Go_](https://github.com/bblfsh/go-driver/blob/master/driver/normalizer/annotation.go), [_Java_](https://github.com/bblfsh/java-driver/blob/master/driver/normalizer/annotation.go), [_Javascript_](https://github.com/bblfsh/javascript-driver/blob/master/driver/normalizer/annotation.go), [_Php_](https://github.com/bblfsh/php-driver/blob/master/driver/normalizer/annotation.go), [_Python_](https://github.com/bblfsh/python-driver/blob/master/driver/normalizer/annotation.go), [_Ruby_](https://github.com/bblfsh/ruby-driver/blob/master/driver/normalizer/annotation.go)

## Friend

Friend is an access granter for some private resources.

**Supported by**:

## Function

Function is a sequence of instructions packaged as a unit.

**Supported by**: [_Go_](https://github.com/bblfsh/go-driver/blob/master/driver/normalizer/annotation.go), [_Java_](https://github.com/bblfsh/java-driver/blob/master/driver/normalizer/annotation.go), [_Javascript_](https://github.com/bblfsh/javascript-driver/blob/master/driver/normalizer/annotation.go), [_Php_](https://github.com/bblfsh/php-driver/blob/master/driver/normalizer/annotation.go), [_Python_](https://github.com/bblfsh/python-driver/blob/master/driver/normalizer/annotation.go), [_Ruby_](https://github.com/bblfsh/ruby-driver/blob/master/driver/normalizer/annotation.go)

## Goto

Goto is an unconditional transfer of control statement.

**Supported by**: [_Go_](https://github.com/bblfsh/go-driver/blob/master/driver/normalizer/annotation.go), [_Php_](https://github.com/bblfsh/php-driver/blob/master/driver/normalizer/annotation.go)

## GreaterThan

GreaterThan is a comparison predicate that checks if the lhs value is greater than the rhs value \(i. e. `>`.\)

**Supported by**: [_Go_](https://github.com/bblfsh/go-driver/blob/master/driver/normalizer/annotation.go), [_Java_](https://github.com/bblfsh/java-driver/blob/master/driver/normalizer/annotation.go), [_Javascript_](https://github.com/bblfsh/javascript-driver/blob/master/driver/normalizer/annotation.go), [_Php_](https://github.com/bblfsh/php-driver/blob/master/driver/normalizer/annotation.go), [_Python_](https://github.com/bblfsh/python-driver/blob/master/driver/normalizer/annotation.go)

## GreaterThanOrEqual

GreaterThanOrEqual is a comparison predicate that checks if the lhs value is greater than or equal to the rhs value \(i.e. 1&gt;=\`.\)

**Supported by**: [_Go_](https://github.com/bblfsh/go-driver/blob/master/driver/normalizer/annotation.go), [_Java_](https://github.com/bblfsh/java-driver/blob/master/driver/normalizer/annotation.go), [_Javascript_](https://github.com/bblfsh/javascript-driver/blob/master/driver/normalizer/annotation.go), [_Php_](https://github.com/bblfsh/php-driver/blob/master/driver/normalizer/annotation.go), [_Python_](https://github.com/bblfsh/python-driver/blob/master/driver/normalizer/annotation.go), [_Ruby_](https://github.com/bblfsh/ruby-driver/blob/master/driver/normalizer/annotation.go)

## Identical

Identical is an identity predicate \(i. e. `===`, `is`, etc.\)

**Supported by**: [_Javascript_](https://github.com/bblfsh/javascript-driver/blob/master/driver/normalizer/annotation.go), [_Php_](https://github.com/bblfsh/php-driver/blob/master/driver/normalizer/annotation.go), [_Python_](https://github.com/bblfsh/python-driver/blob/master/driver/normalizer/annotation.go), [_Ruby_](https://github.com/bblfsh/ruby-driver/blob/master/driver/normalizer/annotation.go)

## Identifier

Identifier is any form of identifier, used for variable names, functions, packages, etc.

**Supported by**: [_Go_](https://github.com/bblfsh/go-driver/blob/master/driver/normalizer/annotation.go), [_Java_](https://github.com/bblfsh/java-driver/blob/master/driver/normalizer/annotation.go), [_Javascript_](https://github.com/bblfsh/javascript-driver/blob/master/driver/normalizer/annotation.go), [_Php_](https://github.com/bblfsh/php-driver/blob/master/driver/normalizer/annotation.go), [_Python_](https://github.com/bblfsh/python-driver/blob/master/driver/normalizer/annotation.go), [_Ruby_](https://github.com/bblfsh/ruby-driver/blob/master/driver/normalizer/annotation.go)

## If

If is used for if-then\[-else\] statements or expressions. An if-then tree will look like:

```text
If, Statement {
    **[non-If nodes] {
        If, Condition {
            [...]
                 }
    }
    **[non-If* nodes] {
        If, Then {
            [...]
        }
    }
    **[non-If* nodes] {
        If, Else {
            [...]
        }
    }
}
```

The Else node is optional. The order of Condition, Then and Else is not defined.

**Supported by**: [_Go_](https://github.com/bblfsh/go-driver/blob/master/driver/normalizer/annotation.go), [_Java_](https://github.com/bblfsh/java-driver/blob/master/driver/normalizer/annotation.go), [_Javascript_](https://github.com/bblfsh/javascript-driver/blob/master/driver/normalizer/annotation.go), [_Php_](https://github.com/bblfsh/php-driver/blob/master/driver/normalizer/annotation.go), [_Python_](https://github.com/bblfsh/python-driver/blob/master/driver/normalizer/annotation.go), [_Ruby_](https://github.com/bblfsh/ruby-driver/blob/master/driver/normalizer/annotation.go)

## Implements

Implements is the type \(usually an interface\) that another type implements.

**Supported by**: [_Php_](https://github.com/bblfsh/php-driver/blob/master/driver/normalizer/annotation.go)

## Import

Import indicates an import level property.

**Supported by**: [_Go_](https://github.com/bblfsh/go-driver/blob/master/driver/normalizer/annotation.go), [_Java_](https://github.com/bblfsh/java-driver/blob/master/driver/normalizer/annotation.go), [_Javascript_](https://github.com/bblfsh/javascript-driver/blob/master/driver/normalizer/annotation.go), [_Php_](https://github.com/bblfsh/php-driver/blob/master/driver/normalizer/annotation.go), [_Python_](https://github.com/bblfsh/python-driver/blob/master/driver/normalizer/annotation.go)

## Incomplete

Incomplete express that the semantic meaning of the node roles doesn't express the full semantic information. Added in BIP-002.

**Supported by**: [_Go_](https://github.com/bblfsh/go-driver/blob/master/driver/normalizer/annotation.go), [_Java_](https://github.com/bblfsh/java-driver/blob/master/driver/normalizer/annotation.go), [_Javascript_](https://github.com/bblfsh/javascript-driver/blob/master/driver/normalizer/annotation.go), [_Php_](https://github.com/bblfsh/php-driver/blob/master/driver/normalizer/annotation.go), [_Python_](https://github.com/bblfsh/python-driver/blob/master/driver/normalizer/annotation.go), [_Ruby_](https://github.com/bblfsh/ruby-driver/blob/master/driver/normalizer/annotation.go)

## Increment

Increment is an arithmetic operator that increments a value \(i. e. `++i`.\)

**Supported by**: [_Go_](https://github.com/bblfsh/go-driver/blob/master/driver/normalizer/annotation.go), [_Java_](https://github.com/bblfsh/java-driver/blob/master/driver/normalizer/annotation.go), [_Javascript_](https://github.com/bblfsh/javascript-driver/blob/master/driver/normalizer/annotation.go), [_Php_](https://github.com/bblfsh/php-driver/blob/master/driver/normalizer/annotation.go)

## Infix

Infix should mark the nodes which are parents of expression nodes using infix notation, e.g.: a+b. Nodes without Infix or Postfix mark are considered in prefix order by default.

**Supported by**:

## Initialization

Initialization is the assignment of an initial value to a variable \(i.e. a for loop variable initialization.\)

**Supported by**: [_Go_](https://github.com/bblfsh/go-driver/blob/master/driver/normalizer/annotation.go), [_Java_](https://github.com/bblfsh/java-driver/blob/master/driver/normalizer/annotation.go), [_Javascript_](https://github.com/bblfsh/javascript-driver/blob/master/driver/normalizer/annotation.go), [_Php_](https://github.com/bblfsh/php-driver/blob/master/driver/normalizer/annotation.go)

## Instance

Instance is a concrete occurrence of an object.

**Supported by**: [_Java_](https://github.com/bblfsh/java-driver/blob/master/driver/normalizer/annotation.go), [_Javascript_](https://github.com/bblfsh/javascript-driver/blob/master/driver/normalizer/annotation.go), [_Ruby_](https://github.com/bblfsh/ruby-driver/blob/master/driver/normalizer/annotation.go)

## Invalid

Invalid Role is assigned as a zero value since protobuf enum definition must start at 0.

**Supported by**:

## Iterator

Iterator is the element that iterates over something.

**Supported by**: [_Go_](https://github.com/bblfsh/go-driver/blob/master/driver/normalizer/annotation.go), [_Java_](https://github.com/bblfsh/java-driver/blob/master/driver/normalizer/annotation.go), [_Javascript_](https://github.com/bblfsh/javascript-driver/blob/master/driver/normalizer/annotation.go), [_Php_](https://github.com/bblfsh/php-driver/blob/master/driver/normalizer/annotation.go), [_Python_](https://github.com/bblfsh/python-driver/blob/master/driver/normalizer/annotation.go), [_Ruby_](https://github.com/bblfsh/ruby-driver/blob/master/driver/normalizer/annotation.go)

## Key

Key is the index value of a map.

**Supported by**: [_Go_](https://github.com/bblfsh/go-driver/blob/master/driver/normalizer/annotation.go), [_Javascript_](https://github.com/bblfsh/javascript-driver/blob/master/driver/normalizer/annotation.go), [_Python_](https://github.com/bblfsh/python-driver/blob/master/driver/normalizer/annotation.go)

## Left

Left is a left hand side in a binary expression.

**Supported by**: [_Go_](https://github.com/bblfsh/go-driver/blob/master/driver/normalizer/annotation.go), [_Java_](https://github.com/bblfsh/java-driver/blob/master/driver/normalizer/annotation.go), [_Javascript_](https://github.com/bblfsh/javascript-driver/blob/master/driver/normalizer/annotation.go), [_Php_](https://github.com/bblfsh/php-driver/blob/master/driver/normalizer/annotation.go), [_Python_](https://github.com/bblfsh/python-driver/blob/master/driver/normalizer/annotation.go), [_Ruby_](https://github.com/bblfsh/ruby-driver/blob/master/driver/normalizer/annotation.go)

## LeftShift

LeftShift is a left shift operation \(i.e. `<<`, `rol`, etc.\)

**Supported by**: [_Go_](https://github.com/bblfsh/go-driver/blob/master/driver/normalizer/annotation.go), [_Java_](https://github.com/bblfsh/java-driver/blob/master/driver/normalizer/annotation.go), [_Javascript_](https://github.com/bblfsh/javascript-driver/blob/master/driver/normalizer/annotation.go), [_Php_](https://github.com/bblfsh/php-driver/blob/master/driver/normalizer/annotation.go), [_Python_](https://github.com/bblfsh/python-driver/blob/master/driver/normalizer/annotation.go), [_Ruby_](https://github.com/bblfsh/ruby-driver/blob/master/driver/normalizer/annotation.go)

## LessThan

LessThan is a comparison predicate that checks if the lhs value is smaller than the rhs value \(i. e. `<`.\)

**Supported by**: [_Go_](https://github.com/bblfsh/go-driver/blob/master/driver/normalizer/annotation.go), [_Java_](https://github.com/bblfsh/java-driver/blob/master/driver/normalizer/annotation.go), [_Javascript_](https://github.com/bblfsh/javascript-driver/blob/master/driver/normalizer/annotation.go), [_Php_](https://github.com/bblfsh/php-driver/blob/master/driver/normalizer/annotation.go), [_Python_](https://github.com/bblfsh/python-driver/blob/master/driver/normalizer/annotation.go)

## LessThanOrEqual

LessThanOrEqual is a comparison predicate that checks if the lhs value is smaller or equal to the rhs value \(i.e. `<=`.\)

**Supported by**: [_Go_](https://github.com/bblfsh/go-driver/blob/master/driver/normalizer/annotation.go), [_Java_](https://github.com/bblfsh/java-driver/blob/master/driver/normalizer/annotation.go), [_Javascript_](https://github.com/bblfsh/javascript-driver/blob/master/driver/normalizer/annotation.go), [_Php_](https://github.com/bblfsh/php-driver/blob/master/driver/normalizer/annotation.go), [_Python_](https://github.com/bblfsh/python-driver/blob/master/driver/normalizer/annotation.go), [_Ruby_](https://github.com/bblfsh/ruby-driver/blob/master/driver/normalizer/annotation.go)

## List

List is a sequence.

**Supported by**: [_Go_](https://github.com/bblfsh/go-driver/blob/master/driver/normalizer/annotation.go), [_Java_](https://github.com/bblfsh/java-driver/blob/master/driver/normalizer/annotation.go), [_Javascript_](https://github.com/bblfsh/javascript-driver/blob/master/driver/normalizer/annotation.go), [_Php_](https://github.com/bblfsh/php-driver/blob/master/driver/normalizer/annotation.go), [_Python_](https://github.com/bblfsh/python-driver/blob/master/driver/normalizer/annotation.go), [_Ruby_](https://github.com/bblfsh/ruby-driver/blob/master/driver/normalizer/annotation.go)

## Literal

Literal is a literal value.

**Supported by**: [_Go_](https://github.com/bblfsh/go-driver/blob/master/driver/normalizer/annotation.go), [_Java_](https://github.com/bblfsh/java-driver/blob/master/driver/normalizer/annotation.go), [_Javascript_](https://github.com/bblfsh/javascript-driver/blob/master/driver/normalizer/annotation.go), [_Php_](https://github.com/bblfsh/php-driver/blob/master/driver/normalizer/annotation.go), [_Python_](https://github.com/bblfsh/python-driver/blob/master/driver/normalizer/annotation.go), [_Ruby_](https://github.com/bblfsh/ruby-driver/blob/master/driver/normalizer/annotation.go)

## Map

Map is a collection of key, value pairs.

**Supported by**: [_Go_](https://github.com/bblfsh/go-driver/blob/master/driver/normalizer/annotation.go), [_Javascript_](https://github.com/bblfsh/javascript-driver/blob/master/driver/normalizer/annotation.go), [_Php_](https://github.com/bblfsh/php-driver/blob/master/driver/normalizer/annotation.go), [_Python_](https://github.com/bblfsh/python-driver/blob/master/driver/normalizer/annotation.go), [_Ruby_](https://github.com/bblfsh/ruby-driver/blob/master/driver/normalizer/annotation.go)

## Module

Module is a set of functionality grouped.

**Supported by**: [_Javascript_](https://github.com/bblfsh/javascript-driver/blob/master/driver/normalizer/annotation.go), [_Php_](https://github.com/bblfsh/php-driver/blob/master/driver/normalizer/annotation.go), [_Python_](https://github.com/bblfsh/python-driver/blob/master/driver/normalizer/annotation.go), [_Ruby_](https://github.com/bblfsh/ruby-driver/blob/master/driver/normalizer/annotation.go)

## Modulo

Modulo is an arithmetic operator \(i.e. `%`, `mod`, etc.\)

**Supported by**: [_Go_](https://github.com/bblfsh/go-driver/blob/master/driver/normalizer/annotation.go), [_Java_](https://github.com/bblfsh/java-driver/blob/master/driver/normalizer/annotation.go), [_Javascript_](https://github.com/bblfsh/javascript-driver/blob/master/driver/normalizer/annotation.go), [_Php_](https://github.com/bblfsh/php-driver/blob/master/driver/normalizer/annotation.go), [_Python_](https://github.com/bblfsh/python-driver/blob/master/driver/normalizer/annotation.go), [_Ruby_](https://github.com/bblfsh/ruby-driver/blob/master/driver/normalizer/annotation.go)

## Multiply

Multiply is an arithmetic operator \(i.e. `*`.\)

**Supported by**: [_Go_](https://github.com/bblfsh/go-driver/blob/master/driver/normalizer/annotation.go), [_Java_](https://github.com/bblfsh/java-driver/blob/master/driver/normalizer/annotation.go), [_Javascript_](https://github.com/bblfsh/javascript-driver/blob/master/driver/normalizer/annotation.go), [_Php_](https://github.com/bblfsh/php-driver/blob/master/driver/normalizer/annotation.go), [_Python_](https://github.com/bblfsh/python-driver/blob/master/driver/normalizer/annotation.go), [_Ruby_](https://github.com/bblfsh/ruby-driver/blob/master/driver/normalizer/annotation.go)

## Name

Name is an identifier used to reference a value.

**Supported by**: [_Go_](https://github.com/bblfsh/go-driver/blob/master/driver/normalizer/annotation.go), [_Java_](https://github.com/bblfsh/java-driver/blob/master/driver/normalizer/annotation.go), [_Javascript_](https://github.com/bblfsh/javascript-driver/blob/master/driver/normalizer/annotation.go), [_Php_](https://github.com/bblfsh/php-driver/blob/master/driver/normalizer/annotation.go), [_Python_](https://github.com/bblfsh/python-driver/blob/master/driver/normalizer/annotation.go), [_Ruby_](https://github.com/bblfsh/ruby-driver/blob/master/driver/normalizer/annotation.go)

## Negative

Negative is an arithmetic operator that negates a value \(i.e. `-x`.\)

**Supported by**: [_Go_](https://github.com/bblfsh/go-driver/blob/master/driver/normalizer/annotation.go), [_Java_](https://github.com/bblfsh/java-driver/blob/master/driver/normalizer/annotation.go), [_Javascript_](https://github.com/bblfsh/javascript-driver/blob/master/driver/normalizer/annotation.go), [_Python_](https://github.com/bblfsh/python-driver/blob/master/driver/normalizer/annotation.go)

## Noop

Noop is a construct that does nothing.

**Supported by**: [_Php_](https://github.com/bblfsh/php-driver/blob/master/driver/normalizer/annotation.go), [_Python_](https://github.com/bblfsh/python-driver/blob/master/driver/normalizer/annotation.go)

## Not

Not is a negation operation. It may be used to annotate a complement of an operator.

**Supported by**: [_Go_](https://github.com/bblfsh/go-driver/blob/master/driver/normalizer/annotation.go), [_Java_](https://github.com/bblfsh/java-driver/blob/master/driver/normalizer/annotation.go), [_Javascript_](https://github.com/bblfsh/javascript-driver/blob/master/driver/normalizer/annotation.go), [_Php_](https://github.com/bblfsh/php-driver/blob/master/driver/normalizer/annotation.go), [_Python_](https://github.com/bblfsh/python-driver/blob/master/driver/normalizer/annotation.go), [_Ruby_](https://github.com/bblfsh/ruby-driver/blob/master/driver/normalizer/annotation.go)

## Null

Null is an empty value.

**Supported by**: [_Java_](https://github.com/bblfsh/java-driver/blob/master/driver/normalizer/annotation.go), [_Javascript_](https://github.com/bblfsh/javascript-driver/blob/master/driver/normalizer/annotation.go), [_Php_](https://github.com/bblfsh/php-driver/blob/master/driver/normalizer/annotation.go), [_Python_](https://github.com/bblfsh/python-driver/blob/master/driver/normalizer/annotation.go), [_Ruby_](https://github.com/bblfsh/ruby-driver/blob/master/driver/normalizer/annotation.go)

## Number

Number is a numeric value. This applies to any numeric value whether it is integer or float, any base, scientific notation or not, etc.

**Supported by**: [_Go_](https://github.com/bblfsh/go-driver/blob/master/driver/normalizer/annotation.go), [_Java_](https://github.com/bblfsh/java-driver/blob/master/driver/normalizer/annotation.go), [_Javascript_](https://github.com/bblfsh/javascript-driver/blob/master/driver/normalizer/annotation.go), [_Php_](https://github.com/bblfsh/php-driver/blob/master/driver/normalizer/annotation.go), [_Python_](https://github.com/bblfsh/python-driver/blob/master/driver/normalizer/annotation.go), [_Ruby_](https://github.com/bblfsh/ruby-driver/blob/master/driver/normalizer/annotation.go)

## Operator

Operator is any form of operator.

**Supported by**: [_Go_](https://github.com/bblfsh/go-driver/blob/master/driver/normalizer/annotation.go), [_Java_](https://github.com/bblfsh/java-driver/blob/master/driver/normalizer/annotation.go), [_Javascript_](https://github.com/bblfsh/javascript-driver/blob/master/driver/normalizer/annotation.go), [_Php_](https://github.com/bblfsh/php-driver/blob/master/driver/normalizer/annotation.go), [_Python_](https://github.com/bblfsh/python-driver/blob/master/driver/normalizer/annotation.go), [_Ruby_](https://github.com/bblfsh/ruby-driver/blob/master/driver/normalizer/annotation.go)

## Or

Or is an OR operation \(i.e. `||`, `or`, `|`, etc.\)

**Supported by**: [_Go_](https://github.com/bblfsh/go-driver/blob/master/driver/normalizer/annotation.go), [_Java_](https://github.com/bblfsh/java-driver/blob/master/driver/normalizer/annotation.go), [_Javascript_](https://github.com/bblfsh/javascript-driver/blob/master/driver/normalizer/annotation.go), [_Php_](https://github.com/bblfsh/php-driver/blob/master/driver/normalizer/annotation.go), [_Python_](https://github.com/bblfsh/python-driver/blob/master/driver/normalizer/annotation.go), [_Ruby_](https://github.com/bblfsh/ruby-driver/blob/master/driver/normalizer/annotation.go)

## Package

Package indicates that a package level property.

**Supported by**: [_Java_](https://github.com/bblfsh/java-driver/blob/master/driver/normalizer/annotation.go)

## Pathname

Pathname is a qualified name of some construct.

**Supported by**: [_Go_](https://github.com/bblfsh/go-driver/blob/master/driver/normalizer/annotation.go), [_Java_](https://github.com/bblfsh/java-driver/blob/master/driver/normalizer/annotation.go), [_Javascript_](https://github.com/bblfsh/javascript-driver/blob/master/driver/normalizer/annotation.go), [_Php_](https://github.com/bblfsh/php-driver/blob/master/driver/normalizer/annotation.go), [_Python_](https://github.com/bblfsh/python-driver/blob/master/driver/normalizer/annotation.go)

## Positional

Positional is an element which position has meaning \(i.e. a positional argument in a call\).

**Supported by**: [_Go_](https://github.com/bblfsh/go-driver/blob/master/driver/normalizer/annotation.go), [_Java_](https://github.com/bblfsh/java-driver/blob/master/driver/normalizer/annotation.go), [_Python_](https://github.com/bblfsh/python-driver/blob/master/driver/normalizer/annotation.go)

## Positive

Positive is an arithmetic operator that makes a value positive. It's usually redundant \(i.e. `+x`.\)

**Supported by**: [_Go_](https://github.com/bblfsh/go-driver/blob/master/driver/normalizer/annotation.go), [_Java_](https://github.com/bblfsh/java-driver/blob/master/driver/normalizer/annotation.go), [_Javascript_](https://github.com/bblfsh/javascript-driver/blob/master/driver/normalizer/annotation.go), [_Python_](https://github.com/bblfsh/python-driver/blob/master/driver/normalizer/annotation.go)

## Postfix

Postfix should mark the nodes which are parents of nodes using postfix notation, e.g.: ab+. Nodes without Infix or Postfix mark are considered in prefix order by default.

**Supported by**: [_Java_](https://github.com/bblfsh/java-driver/blob/master/driver/normalizer/annotation.go), [_Javascript_](https://github.com/bblfsh/javascript-driver/blob/master/driver/normalizer/annotation.go), [_Php_](https://github.com/bblfsh/php-driver/blob/master/driver/normalizer/annotation.go)

## Primitive

Primitive is a language builtin.

**Supported by**: [_Go_](https://github.com/bblfsh/go-driver/blob/master/driver/normalizer/annotation.go), [_Java_](https://github.com/bblfsh/java-driver/blob/master/driver/normalizer/annotation.go), [_Python_](https://github.com/bblfsh/python-driver/blob/master/driver/normalizer/annotation.go), [_Ruby_](https://github.com/bblfsh/ruby-driver/blob/master/driver/normalizer/annotation.go)

## Qualified

Qualified is a kind of property identifiers may have, when it's composed of multiple simple identifiers.

**Supported by**: [_Java_](https://github.com/bblfsh/java-driver/blob/master/driver/normalizer/annotation.go), [_Javascript_](https://github.com/bblfsh/javascript-driver/blob/master/driver/normalizer/annotation.go), [_Php_](https://github.com/bblfsh/php-driver/blob/master/driver/normalizer/annotation.go), [_Python_](https://github.com/bblfsh/python-driver/blob/master/driver/normalizer/annotation.go), [_Ruby_](https://github.com/bblfsh/ruby-driver/blob/master/driver/normalizer/annotation.go)

## Receiver

Receiver is the target of a construct \(message, function, etc.\)

**Supported by**: [_Go_](https://github.com/bblfsh/go-driver/blob/master/driver/normalizer/annotation.go), [_Java_](https://github.com/bblfsh/java-driver/blob/master/driver/normalizer/annotation.go), [_Php_](https://github.com/bblfsh/php-driver/blob/master/driver/normalizer/annotation.go), [_Python_](https://github.com/bblfsh/python-driver/blob/master/driver/normalizer/annotation.go)

## Regexp

Regexp is a regular expression.

**Supported by**: [_Javascript_](https://github.com/bblfsh/javascript-driver/blob/master/driver/normalizer/annotation.go), [_Ruby_](https://github.com/bblfsh/ruby-driver/blob/master/driver/normalizer/annotation.go)

## Relational

Relational is a type of operation

**Supported by**: [_Go_](https://github.com/bblfsh/go-driver/blob/master/driver/normalizer/annotation.go), [_Java_](https://github.com/bblfsh/java-driver/blob/master/driver/normalizer/annotation.go), [_Javascript_](https://github.com/bblfsh/javascript-driver/blob/master/driver/normalizer/annotation.go), [_Php_](https://github.com/bblfsh/php-driver/blob/master/driver/normalizer/annotation.go), [_Python_](https://github.com/bblfsh/python-driver/blob/master/driver/normalizer/annotation.go), [_Ruby_](https://github.com/bblfsh/ruby-driver/blob/master/driver/normalizer/annotation.go)

## Return

Return is a return statement. It might have a child expression or not as with naked returns in Go or return in void methods in Java.

**Supported by**: [_Go_](https://github.com/bblfsh/go-driver/blob/master/driver/normalizer/annotation.go), [_Java_](https://github.com/bblfsh/java-driver/blob/master/driver/normalizer/annotation.go), [_Javascript_](https://github.com/bblfsh/javascript-driver/blob/master/driver/normalizer/annotation.go), [_Php_](https://github.com/bblfsh/php-driver/blob/master/driver/normalizer/annotation.go), [_Python_](https://github.com/bblfsh/python-driver/blob/master/driver/normalizer/annotation.go)

## Right

Right is a right hand side if a binary expression.

**Supported by**: [_Go_](https://github.com/bblfsh/go-driver/blob/master/driver/normalizer/annotation.go), [_Java_](https://github.com/bblfsh/java-driver/blob/master/driver/normalizer/annotation.go), [_Javascript_](https://github.com/bblfsh/javascript-driver/blob/master/driver/normalizer/annotation.go), [_Php_](https://github.com/bblfsh/php-driver/blob/master/driver/normalizer/annotation.go), [_Python_](https://github.com/bblfsh/python-driver/blob/master/driver/normalizer/annotation.go), [_Ruby_](https://github.com/bblfsh/ruby-driver/blob/master/driver/normalizer/annotation.go)

## RightShift

RightShift is a right shift operation \(i.e. `>>`, `ror`, etc.\)

**Supported by**: [_Go_](https://github.com/bblfsh/go-driver/blob/master/driver/normalizer/annotation.go), [_Java_](https://github.com/bblfsh/java-driver/blob/master/driver/normalizer/annotation.go), [_Javascript_](https://github.com/bblfsh/javascript-driver/blob/master/driver/normalizer/annotation.go), [_Php_](https://github.com/bblfsh/php-driver/blob/master/driver/normalizer/annotation.go), [_Python_](https://github.com/bblfsh/python-driver/blob/master/driver/normalizer/annotation.go), [_Ruby_](https://github.com/bblfsh/ruby-driver/blob/master/driver/normalizer/annotation.go)

## Role

**Supported by**: [_Go_](https://github.com/bblfsh/go-driver/blob/master/driver/normalizer/annotation.go)

## Scope

Scope is a range in which a variable can be referred.

**Supported by**: [_Go_](https://github.com/bblfsh/go-driver/blob/master/driver/normalizer/annotation.go), [_Java_](https://github.com/bblfsh/java-driver/blob/master/driver/normalizer/annotation.go), [_Javascript_](https://github.com/bblfsh/javascript-driver/blob/master/driver/normalizer/annotation.go), [_Python_](https://github.com/bblfsh/python-driver/blob/master/driver/normalizer/annotation.go)

## Set

Set is a collection of values.

**Supported by**: [_Python_](https://github.com/bblfsh/python-driver/blob/master/driver/normalizer/annotation.go)

## Statement

Statement is some action to be carried out.

**Supported by**: [_Go_](https://github.com/bblfsh/go-driver/blob/master/driver/normalizer/annotation.go), [_Java_](https://github.com/bblfsh/java-driver/blob/master/driver/normalizer/annotation.go), [_Javascript_](https://github.com/bblfsh/javascript-driver/blob/master/driver/normalizer/annotation.go), [_Php_](https://github.com/bblfsh/php-driver/blob/master/driver/normalizer/annotation.go), [_Python_](https://github.com/bblfsh/python-driver/blob/master/driver/normalizer/annotation.go), [_Ruby_](https://github.com/bblfsh/ruby-driver/blob/master/driver/normalizer/annotation.go)

## String

String is a sequence of characters.

**Supported by**: [_Go_](https://github.com/bblfsh/go-driver/blob/master/driver/normalizer/annotation.go), [_Java_](https://github.com/bblfsh/java-driver/blob/master/driver/normalizer/annotation.go), [_Javascript_](https://github.com/bblfsh/javascript-driver/blob/master/driver/normalizer/annotation.go), [_Php_](https://github.com/bblfsh/php-driver/blob/master/driver/normalizer/annotation.go), [_Python_](https://github.com/bblfsh/python-driver/blob/master/driver/normalizer/annotation.go), [_Ruby_](https://github.com/bblfsh/ruby-driver/blob/master/driver/normalizer/annotation.go)

## Subpackage

Subpackage is a package that is below another package in the hierarchy.

**Supported by**:

## Substract

Substract in an arithmetic operator \(i.e. `-`.\)

**Supported by**: [_Go_](https://github.com/bblfsh/go-driver/blob/master/driver/normalizer/annotation.go), [_Java_](https://github.com/bblfsh/java-driver/blob/master/driver/normalizer/annotation.go), [_Javascript_](https://github.com/bblfsh/javascript-driver/blob/master/driver/normalizer/annotation.go), [_Php_](https://github.com/bblfsh/php-driver/blob/master/driver/normalizer/annotation.go), [_Python_](https://github.com/bblfsh/python-driver/blob/master/driver/normalizer/annotation.go), [_Ruby_](https://github.com/bblfsh/ruby-driver/blob/master/driver/normalizer/annotation.go)

## Subtype

Subtype is a type that can be used to substitute another type.

**Supported by**: [_Java_](https://github.com/bblfsh/java-driver/blob/master/driver/normalizer/annotation.go)

## Switch

Switch is used to represent a broad of switch flavors. An expression is evaluated and then compared to the values returned by different case expressions, executing a body associated to the first case that matches. Similar constructions that go beyond expression comparison \(such as pattern matching in Scala's match\) should not be annotated with Switch.

**Supported by**: [_Go_](https://github.com/bblfsh/go-driver/blob/master/driver/normalizer/annotation.go), [_Java_](https://github.com/bblfsh/java-driver/blob/master/driver/normalizer/annotation.go), [_Javascript_](https://github.com/bblfsh/javascript-driver/blob/master/driver/normalizer/annotation.go), [_Php_](https://github.com/bblfsh/php-driver/blob/master/driver/normalizer/annotation.go), [_Ruby_](https://github.com/bblfsh/ruby-driver/blob/master/driver/normalizer/annotation.go)

## TakeAddress

TakeAddress is an operation that gets the memory address of a value \(i. e. `&x`.\)

**Supported by**: [_Go_](https://github.com/bblfsh/go-driver/blob/master/driver/normalizer/annotation.go)

## Then

Then is the clause executed when the Condition is true.

**Supported by**: [_Go_](https://github.com/bblfsh/go-driver/blob/master/driver/normalizer/annotation.go), [_Java_](https://github.com/bblfsh/java-driver/blob/master/driver/normalizer/annotation.go), [_Javascript_](https://github.com/bblfsh/javascript-driver/blob/master/driver/normalizer/annotation.go), [_Php_](https://github.com/bblfsh/php-driver/blob/master/driver/normalizer/annotation.go), [_Python_](https://github.com/bblfsh/python-driver/blob/master/driver/normalizer/annotation.go), [_Ruby_](https://github.com/bblfsh/ruby-driver/blob/master/driver/normalizer/annotation.go)

## This

This represents the self-reference of an object instance in one of its methods. This corresponds to the `this` keyword \(e.g. Java, C++, PHP\), `self` \(e.g. Smalltalk, Perl, Swift\) and `Me` \(e.g. Visual Basic\).

**Supported by**: [_Java_](https://github.com/bblfsh/java-driver/blob/master/driver/normalizer/annotation.go), [_Javascript_](https://github.com/bblfsh/javascript-driver/blob/master/driver/normalizer/annotation.go), [_Ruby_](https://github.com/bblfsh/ruby-driver/blob/master/driver/normalizer/annotation.go)

## Throw

Throw is a statement that creates an exception.

**Supported by**: [_Java_](https://github.com/bblfsh/java-driver/blob/master/driver/normalizer/annotation.go), [_Javascript_](https://github.com/bblfsh/javascript-driver/blob/master/driver/normalizer/annotation.go), [_Php_](https://github.com/bblfsh/php-driver/blob/master/driver/normalizer/annotation.go), [_Python_](https://github.com/bblfsh/python-driver/blob/master/driver/normalizer/annotation.go), [_Ruby_](https://github.com/bblfsh/ruby-driver/blob/master/driver/normalizer/annotation.go)

## Try

Try is a statement for exception handling.

**Supported by**: [_Java_](https://github.com/bblfsh/java-driver/blob/master/driver/normalizer/annotation.go), [_Javascript_](https://github.com/bblfsh/javascript-driver/blob/master/driver/normalizer/annotation.go), [_Php_](https://github.com/bblfsh/php-driver/blob/master/driver/normalizer/annotation.go), [_Python_](https://github.com/bblfsh/python-driver/blob/master/driver/normalizer/annotation.go), [_Ruby_](https://github.com/bblfsh/ruby-driver/blob/master/driver/normalizer/annotation.go)

## Tuple

Tuple is an finite ordered sequence of elements.

**Supported by**: [_Python_](https://github.com/bblfsh/python-driver/blob/master/driver/normalizer/annotation.go), [_Ruby_](https://github.com/bblfsh/ruby-driver/blob/master/driver/normalizer/annotation.go)

## Type

Type is a classification of data.

**Supported by**: [_Go_](https://github.com/bblfsh/go-driver/blob/master/driver/normalizer/annotation.go), [_Java_](https://github.com/bblfsh/java-driver/blob/master/driver/normalizer/annotation.go), [_Javascript_](https://github.com/bblfsh/javascript-driver/blob/master/driver/normalizer/annotation.go), [_Php_](https://github.com/bblfsh/php-driver/blob/master/driver/normalizer/annotation.go), [_Python_](https://github.com/bblfsh/python-driver/blob/master/driver/normalizer/annotation.go), [_Ruby_](https://github.com/bblfsh/ruby-driver/blob/master/driver/normalizer/annotation.go)

## Unannotated

Unannotated will be automatically added by the SDK for nodes that did not receive any annotations with the current version of the driver's `annotations.go` file. Added in BIP-002.

**Supported by**:

## Unary

Unary is any form of unary operator, in contrast with binary operators.

**Supported by**: [_Go_](https://github.com/bblfsh/go-driver/blob/master/driver/normalizer/annotation.go), [_Java_](https://github.com/bblfsh/java-driver/blob/master/driver/normalizer/annotation.go), [_Javascript_](https://github.com/bblfsh/javascript-driver/blob/master/driver/normalizer/annotation.go), [_Php_](https://github.com/bblfsh/php-driver/blob/master/driver/normalizer/annotation.go), [_Python_](https://github.com/bblfsh/python-driver/blob/master/driver/normalizer/annotation.go)

## Unsigned

Unsigned is an form of unsigned operation.

**Supported by**: [_Java_](https://github.com/bblfsh/java-driver/blob/master/driver/normalizer/annotation.go), [_Javascript_](https://github.com/bblfsh/javascript-driver/blob/master/driver/normalizer/annotation.go)

## Update

Update is the assignment of a new value to a variable \(i.e. a for loop variable update.\)

**Supported by**: [_Go_](https://github.com/bblfsh/go-driver/blob/master/driver/normalizer/annotation.go), [_Java_](https://github.com/bblfsh/java-driver/blob/master/driver/normalizer/annotation.go), [_Javascript_](https://github.com/bblfsh/javascript-driver/blob/master/driver/normalizer/annotation.go), [_Php_](https://github.com/bblfsh/php-driver/blob/master/driver/normalizer/annotation.go), [_Python_](https://github.com/bblfsh/python-driver/blob/master/driver/normalizer/annotation.go), [_Ruby_](https://github.com/bblfsh/ruby-driver/blob/master/driver/normalizer/annotation.go)

## Value

Value is an expression that cannot be evaluated any further.

**Supported by**: [_Go_](https://github.com/bblfsh/go-driver/blob/master/driver/normalizer/annotation.go), [_Javascript_](https://github.com/bblfsh/javascript-driver/blob/master/driver/normalizer/annotation.go), [_Php_](https://github.com/bblfsh/php-driver/blob/master/driver/normalizer/annotation.go), [_Python_](https://github.com/bblfsh/python-driver/blob/master/driver/normalizer/annotation.go)

## Variable

Variable is a symbolic name associated with a value

**Supported by**: [_Go_](https://github.com/bblfsh/go-driver/blob/master/driver/normalizer/annotation.go), [_Java_](https://github.com/bblfsh/java-driver/blob/master/driver/normalizer/annotation.go), [_Javascript_](https://github.com/bblfsh/javascript-driver/blob/master/driver/normalizer/annotation.go), [_Php_](https://github.com/bblfsh/php-driver/blob/master/driver/normalizer/annotation.go)

## Visibility

Visibility is an access granter role, usually together with an specifier role

**Supported by**: [_Java_](https://github.com/bblfsh/java-driver/blob/master/driver/normalizer/annotation.go), [_Javascript_](https://github.com/bblfsh/javascript-driver/blob/master/driver/normalizer/annotation.go), [_Php_](https://github.com/bblfsh/php-driver/blob/master/driver/normalizer/annotation.go), [_Python_](https://github.com/bblfsh/python-driver/blob/master/driver/normalizer/annotation.go), [_Ruby_](https://github.com/bblfsh/ruby-driver/blob/master/driver/normalizer/annotation.go)

## While

While is a loop construct with a condition and a body.

**Supported by**: [_Java_](https://github.com/bblfsh/java-driver/blob/master/driver/normalizer/annotation.go), [_Javascript_](https://github.com/bblfsh/javascript-driver/blob/master/driver/normalizer/annotation.go), [_Php_](https://github.com/bblfsh/php-driver/blob/master/driver/normalizer/annotation.go), [_Python_](https://github.com/bblfsh/python-driver/blob/master/driver/normalizer/annotation.go), [_Ruby_](https://github.com/bblfsh/ruby-driver/blob/master/driver/normalizer/annotation.go)

## Whitespace

Whitespace.

**Supported by**: [_Python_](https://github.com/bblfsh/python-driver/blob/master/driver/normalizer/annotation.go)

## World

World is a set of every component.

**Supported by**: [_Java_](https://github.com/bblfsh/java-driver/blob/master/driver/normalizer/annotation.go), [_Php_](https://github.com/bblfsh/php-driver/blob/master/driver/normalizer/annotation.go), [_Python_](https://github.com/bblfsh/python-driver/blob/master/driver/normalizer/annotation.go), [_Ruby_](https://github.com/bblfsh/ruby-driver/blob/master/driver/normalizer/annotation.go)

## Xor

Xor is an exclusive OR operation \(i.e. `~`, `^`, etc.\)

**Supported by**: [_Go_](https://github.com/bblfsh/go-driver/blob/master/driver/normalizer/annotation.go), [_Java_](https://github.com/bblfsh/java-driver/blob/master/driver/normalizer/annotation.go), [_Javascript_](https://github.com/bblfsh/javascript-driver/blob/master/driver/normalizer/annotation.go), [_Php_](https://github.com/bblfsh/php-driver/blob/master/driver/normalizer/annotation.go), [_Python_](https://github.com/bblfsh/python-driver/blob/master/driver/normalizer/annotation.go), [_Ruby_](https://github.com/bblfsh/ruby-driver/blob/master/driver/normalizer/annotation.go)

