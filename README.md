# Babelfish - Universal Code Parser

## Introduction

**Babelfish is a self-hosted server for source code parsing.** The Babelfish service can parse any file, in any supported language, extracting an [Abstract Syntax Tree \(AST\)](https://en.wikipedia.org/wiki/Abstract_syntax_tree) from it and converting it into a [**Universal Abstract Syntax Tree \(UAST\)**](uast/uast-specification.md).

### Motivation

Babelfish was born as a solution for large scale code analysis. It was created to be able to analyze each revision of all of the world's public source code. 

### Analysis Scope

The current **scope is to parse single files in any programming language.**

This scope might expand in the future to full project analysis, where source code can be analyzed with its full context, and not just per-file. Once we get closer to our current scope, we will consider starting such effort.

### Use Cases

Some of the use cases that we want to support with AST are:

* **Extract features for machine learning on code \(MLonCode\) research.** 
  * e.g. List of all identifiers for each file, or a list of all function calls, etc.
* **AST-based diff'ing**
  * e.g. Understanding changes made to code with finer-grained granularity. Is this commit changing variable names? Is it adding a loop?
* **Import extraction**
  * Extracting all imports from every language in a uniform way.
* **Language feature analytics**
  * e.g. How many people use for-comprehension in Python?
* **Developer tooling** 
  * Improved linters, safety analysis, idiomatic usage, etc.

### Further reading

This repo contains the project documentation, which you can also see properly rendered at [https://doc.bblf.sh/](https://doc.bblf.sh/).

