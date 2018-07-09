# Babelfish - Universal Code Parser

## Introduction

**Babelfish is a self-hosted server for source code parsing.** The Babelfish service can parse any file, in any supported language, extracting an [Abstract Syntax Tree \(AST\)](https://en.wikipedia.org/wiki/Abstract_syntax_tree) from it and converting it into a [**Universal Abstract Syntax Tree \(UAST\)**](uast/uast-specification.md) which enables further analysis and transformation with either the included tools or your own tools by using an standard format.

### Motivation and Scope

Babelfish was born as a solution for massive code analysis. The goal: analyzing all source code from every repository in the world, for every revision ever.

Current **scope is parsing single files in any programming language** and producing a universal abstract syntax tree. This helps us to run it at scale.

This scope might expand in the future to full project analysis, where source code can be analyzed with its full context, and not just per-file. Once we get closer to our current scope, we may consider starting such effort.

### Use Cases

Some of the use cases that we want to support with AST are:

* **AST-based diff'ing.** Understanding changes made to code with finer-grained granularity. Is this commit changing variable names? Is it adding a loop?

* **Import extraction.** Extracting all imports from every language in a uniform way.

* **Extract representation for Data Science experiments.** For example, extracting a list of all tokens for every file, or a list of all function calls, etc.

* **Making statistics of language features.** How many people use for-comprehension in Python?

* **Detecting similar coding patterns across languages.**

* **Programmer-assisting tools** Improved linters, safety analysis, idiomatic usage, etc.

### Further Reading

This repo contains the project documentation, which you can also see properly rendered at [https://doc.bblf.sh/](https://doc.bblf.sh/).
