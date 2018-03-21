# Babelfish vs Other Software

Table of Contents
=================

* [Kythe](#kythe)
* [Language Server Protocol](#language-server-protocol)
* [srclib](#srclib)
* [ctags](#ctags)
* [ANTLR](#antlr)
* [Tree\-sitter](#tree-sitter)
* [srcML](#srcml)
* [SmaCC](#smacc)

<!-- TODO: https://github.com/oracle/opengrok/wiki/Comparison-with-Similar-Tools -->

## [Kythe](http://kythe.io/)

> The best way to view Kythe is as a “hub” for connecting tools for various languages, clients and build systems. By defining language-agnostic protocols and data formats for representing, accessing and querying source code information as data, Kythe allows language analysis and indexing to be run as services.

> Some tools (e.g., static analyzers) already have expressive purpose-built internal representations for code. Kythe is not meant to be a universal replacement for such IRs — instead, our goal is to provide a way for such tools to capture “interesting subsets” of an analysis for sharing with other tools.

Babelfish and Kythe share a goal of defining a common representation
for concepts from different programming languages and provide a way to query it.
Both provide a unified data format. And both are language-independent.

The main difference is that Babelfish preserves all AST nodes including
control flow and expressions, while Kythe focuses on class hierarchy,
dependencies, etc.

Babelfish provides a unified IR across languages, which is defined
as a non-goal of Kythe.

At the same time since Babelfish only processes one file at the time, it cannot
resolve symbolic references in the code, while the goal on Kythe project is to
provide this information with a very high accuracy.

Kythe requires to instrument language compilers and build systems (by modifying
the sources of compiler and build system), while Babelfish uses native language
parsers to get an AST, allowing to develop a new language driver in less time.

Kythe requires a project to be built to be processed because of integration with
compilers and build systems, while Babelfish can process only individual files.

## [Language Server Protocol](https://microsoft.github.io/language-server-protocol/)

> A Language Server is meant to provide the language-specific smarts and communicate with development tools over a protocol that enables inter-process communication.

> The idea behind the Language Server Protocol (LSP) is to standardize the protocol for how such servers and development tools communicate. This way, a single Language Server can be re-used in multiple development tools, which in turn can support multiple languages with minimal effort.

LSP defines a common protocol and RPC for queries like Go-To-Definition,
Usages, etc. But it does not define a common representation of AST because
the goal of the project is to enable easy access to analysis that is done
by compilers and language SDKs.

Babelfish provides a common representation for ASTs allowing to use its output
for static analysis by other tools. Babelfish cannot resolve symbolic references,
thus it cannot be used to query usages.

## [srclib](https://srclib.org/)

> srclib makes developer tools like code search and static analyzers better. It supports things like jump to definition, find usages, type inference, and documentation generation.

> srclib handles: package detection, global dependency resolution, type inference, querying the graph of definitions and references in code, versioning using different VCS systems, and semantic blaming.

<!-- TODO -->

## [ctags](http://ctags.sourceforge.net/)

> Ctags generates an index (or tag) file of language objects found in source files that allows these items to be quickly and easily located by a text editor or other utility. A tag signifies a language object for which an index entry is available (or, alternatively, the index entry created for that object).

Both Babelfish and Ctags provide positional information for identifiers,
classes, directives, etc.

The main difference is that Ctags does not provide any form of AST, while
Babelfish provides native language AST as well as language-independent UAST.

Also, Babelfish ecosystem allows making more complex queries over AST.

## [ANTLR](https://github.com/antlr/antlr4)

> ANTLR (ANother Tool for Language Recognition) is a powerful parser generator for reading, processing, executing, or translating structured text or binary files. It's widely used to build languages, tools, and frameworks. From a grammar, ANTLR generates a parser that can build parse trees and also generates a listener interface (or visitor) that makes it easy to respond to the recognition of phrases of interest.

ANTLR is a language generator with support for most existing languages
and data formats. But parsers usually produce parse trees that are usually
different from native language AST, thus requiring additional processing
to be used for analysis. The structure of parse tree is also language-dependent.

Babelfish provides a correct native language AST as well as
language-independent UAST for all supported languages.

## [Tree-sitter](https://github.com/tree-sitter/tree-sitter)

> Tree-sitter is a C library for incremental parsing, intended to be used via bindings to higher-level languages. It can be used to build a concrete syntax tree for a program and efficiently update the syntax tree as the program is edited. This makes it suitable for use in text-editing programs.

Tree-sitter provides a simplified AST to be able to execute a limited set
of queries. It usually cannot provide enough details suitable for static
analysis. Instead, it focuses on performance and implements real-time
tree diffing.

Babelfish provides a full native language AST as well as language-independent
UAST for all supported languages. It is suitable for static analysis since
it preserves all features of original AST.

## [srcML](https://www.srcml.org/)

> The srcML format is an XML representation for source code, where the markup tags identify elements of the abstract syntax for the language. The srcml program is a command line application for the conversion source code to srcML, an interface for the exploration, analysis, and manipulation of source code in this form, and the conversion of srcML back to source code. The current parsing technologies support C/C++, C#, and Java.

srcML defines an XML schema to annotate source code files with AST structure.
It also provides tools to query and analyze files in this format. An AST
structure is language-dependent.

Babelfish provides a native language AST with positional information, that
can be used to generate the same markup. Also, it provides a language-independent
UAST for all supported languages. It allows performing the same queries
and analysis for different programming languages.

## [SmaCC](http://www.refactoryworkers.com/SmaCC.html)

> SmaCC (Smalltalk Compiler-Compiler) is a freely available parser generator for Smalltalk. It generates LR parsers and is a replacement for the T-Gen parser generator. SmaCC overcomes many of T-Gen's limitations that make it difficult to produce parsers. SmaCC can generate parsers for ambiguous grammars and grammars with overlapping tokens. Both of these are not possible using T-Gen. In addition to handling more grammars than T-Gen, SmaCC has a smaller runtime than T-Gen and is faster than T-Gen. The latest version of SmaCC has support for GLR parsing, generating abstract syntax trees (ASTs), and transforming code.

SmaCC provides tools to build parser generators that can produce ASTs
that are close to native ASTs of programming languages. It also allows
applying transformations and rewrites to the tree. But generated AST is
still language-dependent.

Babelfish also generates a native AST and provides an SDK to perform tree
rewrites. But it also provides a language-independent UAST for all
supported languages.