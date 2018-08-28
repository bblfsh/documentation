# UAST v2

## Overview

A Universal Abstract Syntax Tree \(UAST\) is a normalized form of [Abstract Syntax Tree \(AST\)](https://en.wikipedia.org/wiki/Abstract_syntax_tree).
By definition, the UAST representation does not depend on the language
of the parser for source files.

UAST is defined by two separate specifications:

* [Node representation](./representation_v2.md) - defines a common transport
  format for the UAST.

* [Semantic UAST](./semantic_uast.md) - defines UAST node types that are
  independent of the programming language.
