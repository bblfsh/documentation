# Code to AST

For each language, we generate a language-specific AST. This AST is as close as possible to what we get from the standard or common implementation of the AST in each language.

Drivers MUST guarantee that given the same `(blob, driver, driver version)`, the generated AST is always the same.

Drivers SHOULD generate an AST without any information loss from the source code. That is, it should be possible to write a function that given a AST, generates the same source code that was used to generate that AST, byte by byte.

Most parsers available do not completely prevent information loss during the conversion. So we currently aim to preserve as much information as possible. We define different levels of information loss in terms of which kind of code generation would they allow.

* **Lossless:** Converting code to AST and then back to code would.

  `code == codegen(AST(code))`.

* **Formatting information loss:** Only superfluous formatting information is lost

  \(e.g. whitespace, indentation\). Code generated from the AST could be the same

  as the original code after passing a code formatter. `fmt(code) == codegen(AST(code))`.

* **Syntactic sugar information loss:** There is information loss about syntactic

  sugar. Code generated from the AST could be the same as the original code after

  desugaring it. `desugar(code) == codegen(AST(code))`.

* **Comment loss:** Comments are not present in the AST.

Drivers MUST NOT lose any data beyond the previous cases.

