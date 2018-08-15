# Babelfish - Universal Code Parser

## Introduction

**Babelfish is a self-hosted server for source code parsing.** The Babelfish service can parse any file, in any supported language, extracting an [Abstract Syntax Tree \(AST\)](https://en.wikipedia.org/wiki/Abstract_syntax_tree) from it and converting it into a [**Universal Abstract Syntax Tree \(UAST\)**](uast/uast-specification.md). The UAST enables further analysis and transformations with either the included tools or your own tools by providing a standard open format.

### Motivation & Scope

Babelfish was created as a solution for large scale code analysis. To analyze the source code from millions of repositories, at each revision.

The current **scope is to enable parsing of single files in any popular programming language** and producing a [Universal Abstract Syntax Tree \(UAST\)](uast/uast-specification.md). 

This current scope is expected to expand in the near future to full project analysis, where the source code can be analyzed with its full context, and not just per-file. 

### Use Cases

Some of the use cases that we aim to support with UAST are:

* **Feature extraction for Machine Learning on Code:** For example, extracting a list of all tokens for every file, or a list of all function calls, etc.
* **Language-agnostic static analysis:** making it easy to write static analyzers in any language, analyzing any supported language
* **UAST diffs:** Understanding changes made to code with finer-grained granularity. Is this commit changing variable names? Is it adding a loop?
* **Uniform import extraction:** Extracting all imports from every language in a uniform way.
* **Statistical analysis of language features:** How many people use for-comprehension in Python.

### Further Reading

This repository contains the project documentation, which you can also see properly rendered at [https://docs.sourced.tech/babelfish](https://docs.sourced.tech/babelfish).

