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

From official [Kythe documentation](https://kythe.io/docs/kythe-overview.html#_goals_of_kythe):
> The best way to view Kythe is as a “hub” for connecting tools for various languages, clients and build systems. By defining language-agnostic protocols and data formats for representing, accessing and querying source code information as data, Kythe allows language analysis and indexing to be run as services.

> Some tools (e.g., static analyzers) already have expressive purpose-built internal representations for code. Kythe is not meant to be a universal replacement for such IRs — instead, our goal is to provide a way for such tools to capture “interesting subsets” of an analysis for sharing with other tools.

Babelfish and Kythe share the same goal of defining a common representation
for concepts from different programming languages and providing a way to query it.
Both provide a unified data format. And both are language-independent.

The main difference is that Babelfish preserves all AST nodes including
the control flow and expressions, while Kythe does not store the whole AST and only preserves derivative
information specific to the use case of code navigation, e.g.
links to declarations, resolved references, etc.

Babelfish provides a unified IR across languages, which is defined
as a non-goal of Kythe.

At the same time, since Babelfish only processes one file at a time, it cannot
resolve symbolic references in the code, while the goal of Kythe project is to
provide that information with very high accuracy.

Kythe requires:
 - instrumented language compilers and build systems (by modifying
the sources of the compiler and the build system),
 - a project to be built through the analysis phase.

Whereas Babelfish:
 - works on the individual file level,
 - uses existing native language parsers to get an AST.

The latter features allow to develop a new language driver faster, as well as to scale the analysis to millions of projects.

## [Language Server Protocol](https://microsoft.github.io/language-server-protocol/)

From the official [documentation](https://microsoft.github.io/language-server-protocol/overview):
> A Language Server is meant to provide the language-specific smarts and communicate with development tools over a protocol that enables inter-process communication.

> The idea behind the Language Server Protocol (LSP) is to standardize the protocol for how such servers and development tools communicate. This way, a single Language Server can be re-used in multiple development tools, which in turn can support multiple languages with minimal effort.

LSP defines a common protocol and RPC for queries like Go-To-Definition,
Usages, etc. The list of LSP-compatible language servers and clients is available at
[langserver.org](https://langserver.org/).
However, it does not define a common AST representation because
the goal of the project is to enable easy access to analysis that is performed
by compilers and language SDKs.

Babelfish provides a common AST representation which is friendly
for static analysis with other tools. Babelfish cannot resolve symbolic references,
thus it cannot be used to query for the usages of a given identifier.

## [srclib](https://srclib.org/)

_The project is no longer maintained, and was replaced by LSP_.

According to the [documentation](https://srclib.org/):
> srclib makes developer tools like code search and static analyzers better. It supports things like jump to definition, find usages, type inference, and documentation generation.

> srclib handles: package detection, global dependency resolution, type inference, querying the graph of definitions and references in code, versioning using different VCS systems, and semantic blaming.

srclib:
 - requires running a build to analyze.
 - extracts the graph of references/definitions.
 - uses language-specific toolchains based on http://ternjs.net/, https://yardoc.org/, https://jedi.readthedocs.io/en/latest/, etc.


## [ctags](https://ctags.io/)

From the project [documentation](http://ctags.sourceforge.net/whatis.html):
> Ctags generates an index (or tag) file of language objects found in source files that allows these items to be quickly and easily located by a text editor or other utility. A tag signifies a language object for which an index entry is available (or, alternatively, the index entry created for that object).

Both Babelfish and Ctags provide the positional information about identifiers,
classes, directives, etc.

Ctag is a binary application that only provides a serialized index, while Bblfsh relies on the [client libraries](https://doc.bblf.sh/using-babelfish/clients.html) for multiple languages to access the data.

The main difference is that Ctags does not provide any form of AST, while
Babelfish preserves the native language AST as well as the language-independent UAST.

Besides, Babelfish ecosystem supports more complex queries based on XPath over AST, while Ctags provides nothing but the fast identifier lookup.

## [ANTLR](https://github.com/antlr/antlr4)

From the project [documentation](http://www.antlr.org/about.html):
> ANTLR (ANother Tool for Language Recognition) is a powerful parser generator for reading, processing, executing, or translating structured text or binary files. It's widely used to build languages, tools, and frameworks. From a grammar, ANTLR generates a parser that can build parse trees and also generates a listener interface (or visitor) that makes it easy to respond to the recognition of phrases of interest.

ANTLR is one of the best-known modern representatives of the family of tools known as `LL(*)` - "parser generators". It is written in Java and supports the majority of the existing languages and data formats.

In general, there are several important limitations that come with the parser generator approach:
 - parsers generated from formal language grammars produce parse trees, otherwise known as Concrete Syntax Trees, that are language-dependent and different from the native language ASTs, and thus require additional processing to be used for analysis.
 - there is a certain maintenance cost of updating the grammar which is defined outside of the language ecosystem/stdlib.

Babelfish provides access to the native language AST though [client libraries](https://doc.bblf.sh/using-babelfish/clients.html) for multiple languages, as well as the language-independent UAST for all the [supported languages](https://doc.bblf.sh/languages.html) and is based on parser libraries included in stdlib when possible.

## [Tree-sitter](https://github.com/tree-sitter/tree-sitter)

From the [project documentation](http://tree-sitter.github.io/tree-sitter/):
> Tree-sitter is a parser generator tool and an incremental parsing library. It can build a concrete syntax tree for a source file and efficiently update the syntax tree as the source file is edited. This makes it suitable for use in text-editing programs.

Although tree-sitter, being a parser generator, produces a Concrete Syntax Tree, the shape of the tree is designed to closely mirror the structure of the AST and thus can be used directly for code analysis. Other differences include: tree-sitter's syntax nodes currently just contain an ordered list of children, the syntax nodes don't have a name for each child such as a predicate, a body, etc.

That project also focuses on performance and incremental parsing since both are essential
for text editors.

Babelfish provides the full native language AST as well as the language-independent
UAST for all the supported languages. It is suitable for static analysis since
it preserves all the features of the original AST.

## [srcML](https://www.srcml.org/)

From the [project documentation](https://www.srcml.org/about.html):
> The srcML format is an XML representation for source code, where the markup tags identify elements of the abstract syntax for the language. The srcml program is a command line application for the conversion source code to srcML, an interface for the exploration, analysis, and manipulation of source code in this form, and the conversion of srcML back to source code. Supports parsing C/C++, C#, and Java.

srcML defines an [XML schema](https://www.srcml.org/doc/srcMLGrammar.html) to annotate
source code files with their syntactic structure. Given a source code file, it is
converted to an XML document which follows that schema. srcML also provides tools in separate
binaries written in C++ to [query](https://www.srcml.org/tutorials/xpath-query.html)
and analyze files in that format using XPath. The AST structure is language-dependent
and parsing itself is performed using ANTLR.

Babelfish provides the native language AST with positional information, that
can be used to generate the same markup. Besides, it provides the language-independent
UAST for all the supported languages. It allows running the same queries
and analyzing different programming languages.

## [SmaCC](https://github.com/SmaCCRefactoring/SmaCC)

From the [project documentation](http://www.refactoryworkers.com/SmaCC.html):
> SmaCC (Smalltalk Compiler-Compiler) is a freely available parser generator for Smalltalk. It generates LR parsers and is a replacement for the T-Gen parser generator. SmaCC overcomes many of T-Gen's limitations that make it difficult to produce parsers. SmaCC can generate parsers for ambiguous grammars and grammars with overlapping tokens. Both of these are not possible using T-Gen. In addition to handling more grammars than T-Gen, SmaCC has a smaller runtime than T-Gen and is faster than T-Gen. The latest version of SmaCC has support for GLR parsing, generating abstract syntax trees (ASTs), and transforming code.

SmaCC is an LR, LALR and GLR parser generator written in SmallTalk. It can produce ASTs
that are close to native ASTs of programming languages. It includes tools for
applying large-scale transformations and rewrites to the trees. However, the generated AST is still language-dependent.

Babelfish also generates the native AST and provides an SDK to perform tree
rewrites. It additionally provides the language-independent UAST for all
the supported languages.
