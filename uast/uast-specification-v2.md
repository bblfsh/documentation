# UAST v2

## Overview

A Universal Abstract Syntax Tree \(UAST\) is a normalized form of [Abstract Syntax Tree \(AST\)](https://en.wikipedia.org/wiki/Abstract_syntax_tree).
By definition, the UAST representation does not depend on the language
of the parser for source files.

UAST is defined by two separate specifications:

* [Node representation](./representation-v2.md) - defines a common transport
  format for the UAST.

* [Semantic UAST](./semantic-uast.md) - defines UAST node types that are
  independent of the programming language.

## Visualization

Visualizing the structure of the UAST is very simple: YAML output of `bblfsh-cli` can be read by the `bblfhs-sdk ast2gv` command that outputs UAST representation using GraphVis.

PNG and SVG output is supported for convenience but requires a [graphviz](https://www.graphviz.org/) binary to be installed on your system.

```bash
go get gopkg.in/bblfsh/client-go.v3/cmd/bblfsh-cli/...
bblfsh-cli client.go > client.uast.yml

go get gopkg.in/bblfsh/sdk.v2/cmd/bblfsh-sdk/...
bblfsh-sdk ast2gv -o png client.uast.yml
open client.uast.yml.png
```
