# Babelfish - Universal Code Parser

## Introduction

**Babelfish is a self-hosted server for source code parsing.** The Babelfish service can parse any file, in any supported language, extracting an [Abstract Syntax Tree \(AST\)][ast] from it and converting it into a [**Universal Abstract Syntax Tree \(UAST\)**](uast/uast-specification-v2.md). The UAST enables further analysis and transformations with either the included tools or your own tools by providing a standard open format. Jump to the [Getting Started](using-babelfish/getting-started.md) section to start using it!

### Motivation & Scope

Babelfish was created as a solution for large scale code analysis. To analyze the source code from millions of repositories, at each revision.

The current **scope is to enable parsing of single files in any popular programming language** and producing a [Universal Abstract Syntax Tree \(UAST\)](uast/uast-specification-v2.md).

This current scope is expected to expand in the near future to full project analysis, where the source code can be analyzed with its full context, and not just per-file.

For more information about how Babelfish compares to other similar systems,
see [this page](alternatives.md).

### Use Cases

Some of the use cases that we aim to support with UAST are:

* **Feature extraction for Machine Learning on Code:** For example, extracting a list of all tokens for every file, or a list of all function calls, etc.
* **Language-agnostic static analysis:** making it easy to write static analyzers in any language, analyzing any supported language
* **UAST diffs:** Understanding changes made to code with finer-grained granularity. Is this commit changing variable names? Is it adding a loop?
* **Uniform import extraction:** Extracting all imports from every language in a uniform way.
* **Statistical analysis of language features:** How many people use for-comprehension in Python.

### Current status

#### Universal AST coverage

Not all constructs are converted in a language-independent way yet.
As of Q1 2019, every language driver is expected to support Identifiers,
String literals, Imports, and Functions in the [Universal AST schema](uast/uast-specification-v2.md).

That is sufficient to examine the symbols exported by the package, but not control flow (yet).

Meanwhile, in order to identify constructs that are not covered by the UAST
schema yet - one can look for a `@role` field.
A [Role](https://godoc.org/gopkg.in/bblfsh/sdk.v2/uast/role#Role) is added
to every native AST node and it contains a language-independent annotations.
However, the tree structure for a role construct may be different between languages.

#### Protocol v2

Currently, Babelfish is in the process of transition to an updated data model, denoted [UAST v2](uast/uast-specification-v2.md),
which includes a new node representation and canonicalized ("semantic") UAST structures.

All the beta+ drivers support UAST v2 in their latest versions. UAST v2 support requires bblfshd ≥ 2.6.1.

Libuast has been updated to support the new node format, but some of the
the [clients](./using-babelfish/clients.md) may still work in v1 compatibility mode
to be able to execute XPath queries.

See [v2 transition options](./using-babelfish/advanced-usage.md#protocol-v2-transition) for details.

### Further Reading

This repository contains the project documentation, which you can also see properly rendered at [https://docs.sourced.tech/babelfish](https://docs.sourced.tech/babelfish).

[ast]: https://en.wikipedia.org/wiki/Abstract_syntax_tree

### License

This work is licensed under a Creative Commons — Attribution-ShareAlike 4.0 International License. See [LICENSE](./LICENSE.md).
