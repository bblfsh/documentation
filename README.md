<img src="https://avatars2.githubusercontent.com/u/25795418?v=3&s=200f" align="right" width="200px" height="200px" alt="Babelfish logo" />

# Babelfish - universal code parser

**Babelfish is a self-hosted server for source code parsing.** Babelfish can parse any file, in any language, extract an ASTs (Abstract Syntax Tree) from it, and convert it to a [**UAST**](./uast/specification.md) (Universal Abstract Syntax Tree).

## Motivation and Scope

Babelfish was born as a solution for massive code analysis. The goal: analyzing
all source code from every repository in the world, for every revision ever.

Current **scope is parsing single files in any programming language**
and producing a universal abstract syntax tree. This helps us to run it at scale.

This scope might expand in the future to full project analysis, where source code
can be analyzed with its full context, an not just per-file. Once we get closer to
our current scope, we may consider starting such effort.

## Use Cases

Some of the use cases that we want to support with AST are:

* **AST-based diff'ing.** Understanding changes made to code with finer-grained granularity. Is this commit changing variable names? Is it adding a loop?
* **Import extraction.** Extracting all imports from every language in a uniform way.
* **Extract representation for Data Science experiments.** For example, extracting a list of all tokens for every file, or a list of all function calls, etc.
* **Making statistics of language features.** How many people uses for-comprehension in Python?
* **Detecting similar coding patterns across languages.**
