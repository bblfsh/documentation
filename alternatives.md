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

Babelfish and Kythe share a goal of defining a common representation
for concepts from different programming languages and provide a way to query it.
Both provide a unified data format. And both are language-independent.

The main difference is that Babelfish preserves all AST nodes including
control flow and expressions, while Kythe does not store whole AST and only preserve derivative information specific to the use case of code navigation like
links to declarations, resolved references, etc.

Babelfish provides a unified IR across languages, which is defined
as a non-goal of Kythe.

At the same time, since Babelfish only processes one file at the time, it cannot
resolve symbolic references in the code, while the goal on Kythe project is to
provide this information with a very high accuracy.

Kythe requires:
 - instrumented language compilers and build systems (by modifying
the sources of compiler and build system),
 - a project to be built though analysis phase.

While Babelfish:
 - works on individual file-level,
 - uses existing native language parsers to get an AST.

 Both allow to develop a new language driver for Bblfsh in less time, as well as to scale the analysis to millions of projects.

## [Language Server Protocol](https://microsoft.github.io/language-server-protocol/)

From official [documentation](https://microsoft.github.io/language-server-protocol/overview):
> A Language Server is meant to provide the language-specific smarts and communicate with development tools over a protocol that enables inter-process communication.

> The idea behind the Language Server Protocol (LSP) is to standardize the protocol for how such servers and development tools communicate. This way, a single Language Server can be re-used in multiple development tools, which in turn can support multiple languages with minimal effort.

LSP defines a common protocol and RPC for queries like Go-To-Definition,
Usages, etc with a list of LSP-compatible language servers and clients is available at
[langserver.org](https://langserver.org/).
But it does not define a common representation of AST because
the goal of the project is to enable easy access to analysis that is done
by compilers and language SDKs.

Babelfish provides a common representation for ASTs allowing to use its output
for static analysis by other tools. Babelfish cannot resolve symbolic references,
thus it cannot be used to query the usages of a given identifier.

## [srclib](https://srclib.org/)

_The project is no longer maintained, and was replaced by LSP_.

According to [documentation](https://srclib.org/):
> srclib makes developer tools like code search and static analyzers better. It supports things like jump to definition, find usages, type inference, and documentation generation.

> srclib handles: package detection, global dependency resolution, type inference, querying the graph of definitions and references in code, versioning using different VCS systems, and semantic blaming.

 - requires running a build to analyze
 - extracts graph of references/definitions
 - using language-specific toolchains based on http://ternjs.net/, https://yardoc.org/, https://jedi.readthedocs.io/en/latest/, etc


## [ctags](https://ctags.io/)

From project [documentation](http://ctags.sourceforge.net/whatis.html):
> Ctags generates an index (or tag) file of language objects found in source files that allows these items to be quickly and easily located by a text editor or other utility. A tag signifies a language object for which an index entry is available (or, alternatively, the index entry created for that object).

Both Babelfish and Ctags provide positional information for identifiers,
classes, directives, etc.

Ctag is an application binary that only provides index in serialized format, but Bblfsh relies on [client libraries](https://doc.bblf.sh/using-babelfish/clients.html) for multiple languages for data access.

The main difference is that Ctags does not provide any form of AST, while
Babelfish preserves native language AST as well as language-independent UAST.

Also, Babelfish ecosystem allows for more complex queries based on XPath over AST, while with Ctags it's only fast identifier lookup.

## [ANTLR](https://github.com/antlr/antlr4)

From project [documentation](http://www.antlr.org/about.html):
> ANTLR (ANother Tool for Language Recognition) is a powerful parser generator for reading, processing, executing, or translating structured text or binary files. It's widely used to build languages, tools, and frameworks. From a grammar, ANTLR generates a parser that can build parse trees and also generates a listener interface (or visitor) that makes it easy to respond to the recognition of phrases of interest.

ANTLR is one of the best-known modern representatives from the family of tools known as `LL(*)` "parser generators", it's written in Java and supports most existing languages and data formats.

In general, there are several important limitations that come with parser generator approach:
 - parsers generated from formal language grammars produce parse trees, otherwise known as Concrete Syntax Trees, that are language-dependent and different from native language ASTs, and thus require additional processing to be used for analysis
 - there is a maintenance cost for supporting language grammars evolution, defined outside of the language ecosystem/stdlib

Babelfish provides access to native language AST though [client libraries](https://doc.bblf.sh/using-babelfish/clients.html) for multiple languages, as well as language-independent UAST for all [supported languages](https://doc.bblf.sh/languages.html) and is based on parser libraries included in stdlib (when available).

## [Tree-sitter](https://github.com/tree-sitter/tree-sitter)

From [project documentation](http://tree-sitter.github.io/tree-sitter/):
> Tree-sitter is a parser generator tool and an incremental parsing library. It can build a concrete syntax tree for a source file and efficiently update the syntax tree as the source file is edited. This makes it suitable for use in text-editing programs.

Although tree-sitter, as a parser generator, produces a Contact Syntax Tree the shape of the tree is designed to closely mirror the structure of the AST and thus can be used directly for code analysis. Other differences include: tree-sitter's syntax nodes currently just contain an ordered list of children, the syntax nodes don't have a name for each child like predicate, body, etc.

Project also focuses on performance and incremental parsing as both are instrumental
for use-case of text editors.

Babelfish provides a full native language AST as well as language-independent
UAST for all supported languages. It is suitable for static analysis since
it preserves all features of original AST.

## [srcML](https://www.srcml.org/)

From [project documentation](https://www.srcml.org/about.html):
> The srcML format is an XML representation for source code, where the markup tags identify elements of the abstract syntax for the language. The srcml program is a command line application for the conversion source code to srcML, an interface for the exploration, analysis, and manipulation of source code in this form, and the conversion of srcML back to source code. Supports parsing C/C++, C#, and Java.

srcML defines an [XML schema](https://www.srcml.org/doc/srcMLGrammar.html) to annotate
source code files with it's syntactic structure. Given a file with the source code, is
converted to XML document, following that schema. srcML also provides tools in separate
binaries written in C++ to [query](https://www.srcml.org/tutorials/xpath-query.html)
and analyze files in this format using XPath. An AST structure is language-dependent
and parsing itself is done using ANTLR.

Babelfish provides a native language AST with positional information, that
can be used to generate the same markup. Also, it provides a language-independent
UAST for all supported languages. It allows performing the same queries
and analysis for different programming languages.

## [SmaCC](https://github.com/SmaCCRefactoring/SmaCC)

From [project documentation](http://www.refactoryworkers.com/SmaCC.html):
> SmaCC (Smalltalk Compiler-Compiler) is a freely available parser generator for Smalltalk. It generates LR parsers and is a replacement for the T-Gen parser generator. SmaCC overcomes many of T-Gen's limitations that make it difficult to produce parsers. SmaCC can generate parsers for ambiguous grammars and grammars with overlapping tokens. Both of these are not possible using T-Gen. In addition to handling more grammars than T-Gen, SmaCC has a smaller runtime than T-Gen and is faster than T-Gen. The latest version of SmaCC has support for GLR parsing, generating abstract syntax trees (ASTs), and transforming code.

SmaCC is LR, LALR and GLR parser generator written in SmallTalk. It can produce ASTs
that are close to native ASTs of programming languages. It includes tools for
applying large-scale transformations and rewrites to the trees. But generated AST is still language-dependent.

Babelfish also generates a native AST and provides an SDK to perform tree
rewrites. But it also provides a language-independent UAST for all
supported languages.