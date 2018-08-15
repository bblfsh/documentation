# BIP2: Explicit Annotated status and completeness information

| Field | Value |
| :--- | :--- |
| BIP | 2 |
| Title | Explicit Annotated status and completeness information |
| Author | Juanjo Álvarez |
| Status | Accepted |
| Created | 2017-06-01 |
| Updated | 2017-08-02 |

## Abstract

This BIP proposed two new `Incomplete` and `Unannotated` roles would mark if the node has been annotated in any way and the quality of the current annotation.

## Rationale

At the current stage of the Babelfish project, we strive for genericity in the list of UAST roles and not adding very language-specific features into it until we can have a wider view about how to integrate or compose these roles for covering as much ground as possible.

This is a worthy objective but has the problem of leaving the nodes related with these features unannotated or with annotations that don't fully express all the semantic information that the language source code or the native AST have.

This BIP proposes a mechanism to mark nodes with two possible completion status roles. This will make the UAST user have a wider view of the completeness of the information provided by it on specific nodes and also could enable new uses cases \(whose specific definition is outside of the scope of this BIP\) like helping debug drivers and provide a completeness summary, with the added benefit of providing a numerically observable measure of the completeness of every developed driver.

## Example

The current UAST specification doesn't have a role for Python's list comprehensions. The programmer would then add the `Incomplete` and `ForEach` roles thus annotating part of the semantic meaning of the role while at the same time marking the incompleteness of this information to the UAST user.

## Specification

### UAST Roles

This change will add two new roles the UAST definition:

* `Incomplete`: For nodes that don't have their full semanting meaning annotated, including both partially annotated nodes and nodes without any other annotation.
* `Unannotated`: For nodes that have received no other annotations. Automatically added by the SDK. Driver's developers should aim to have none of these.

After this proposal have been fully implemented it'll be implicitly understood that nodes that have roles set and no `Incomplete` role are expressing full semantic information.

### Driver and SDK annotation

The `Incomplete` roles will be set explicitly by the driver's programmer using the normal process in he driver's `annotation.go` file.

The `Unannotated` role will be added automatically by the SDK to nodes that didn't get any other role.

## Impact

* The new roles would have to be added to the SDK's `uast.go` file and the drivers rebuild against it.
* The `Incomplete` mark would have to be added by the driver developers when defining the rules in the `annotation.go` file. Existing drivers will need to be updated.
* The SDK annotation process will implement the automatic would have to implement setting the `Unannotated` marks. This will be implemented in the `protocol/native/objecttonoder.go` file.
* We'll add a new process to the build system that will check both the `driver/normalizer/??ast/??ast.go` file with the static list of nodes for existing but unannotated nodes and the integration test results and output a summary of incomplete an unannotated nodes.
* The documentation for the driver developer would have to be updated to include the purpose and mechanics of these new property/s or role/s. This is in the file `driver/annotations.md`.
* This BIP will be published with the documentation in the `proposals` section.

## Discarded alternatives

During the development of this BIP alternatives to the current implementation were discarded:

* Using just an `Incomplete` role or using three roles \(current `Incomplete` and `Unannotated` roles plus an `Annotated` one\). It was decided that the choosen roles provided information but `Annotated` didn't.
* Using a node property with the completeness mark instead of new roles. It was discarded because roles will be easier to integrate into existing client workflows.

## References

* [Specification of the Node struct with the properties already in use](https://doc.bblf.sh/uast/specification.html)

