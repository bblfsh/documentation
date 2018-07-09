# BIP1: Purpose and Guidelines

| Field | Value |
| --- | --- |
| BIP | 1 |
| Title | BIP Purpose and Guidelines |
| Author | Alfredo Beaumont |
| Status | Accepted |
| Created | 2017-05-26 |
| Updated | 2018-07-08 |

## Abstract

BIP stands for Babelfish Improvement Proposal. A BIP is a design document which describes a change proposal to Babelfish. It's informally defined. The scope of a BIP includes:

* Changes to [UAST](https://doc.bblf.sh/uast/uast-specification.html) components regarding code categorization through Roles or related entities.
* Changes to Babelfish protocols \(i. e. the [server protocol](https://doc.bblf.sh/using-babelfish/babelfish-protocol.html), driver protocol, [internal driver protocol](https://doc.bblf.sh/writing-a-driver/internal-protocol.html), etc.

Its scope may become broader in the future.

The BIP should provide a concise technical specification of, and rational for, the proposed change.

We intend BIPs to be the primary mechanism for proposing new features or changes to existing ones in the UAST components.

## Rationale

Currently changes to UAST are mostly language driven, i. e., whenever a language driver implementor finds that there's some feature missing to complete an AST to UAST conversion, this new feature is requested to be added through a Pull Request. While this process is quick, it doesn't scale well: a single language need is not the best background for a universal feature request discussion, and it easies a process that may degenerate in a UAST that becomes the sum of each single language feature, which is something we should avoid. That's why introducing a more formal process, requiring a larger background and use cases, should help visualising which features are general enough, how to generalize them and how to handle those features that are more specific.

## Workflow

A BIP process should begin with an author proposing an idea. This idea should be as focused as possible.

The author must:

* Discuss the idea publicly, gather early feedback and make sure that the idea is not guaranteed to be rejected.

* Present a draft. It should take the next available BIP number and its status should be **Draft**

* Reach consensus. Feedback should be added back to the proposal. You should open a Pull Request with the proposal in the [babelfish documentation project](https://github.com/bblfsh/documentation/) for this process.

Once the BIP has been completed and reviewed it may be **Accepted** or **Rejected** \(it may have not been a good idea after all\). In any case it's important to have a record of the discussion so the proposal will be archived. The final decision will be done by [mcuadros](https://github.com/mcuadros).

## Specification

Each BIP should have the following parts

1. Title
2. Header with metadata
3. Abstract -- a short description of the issue being addressed
4. Rationale -- proposal motivation. It should be clearly explain why current design is inadequate to address the problem that this BIP solves, and the use cases for it.
5. Specification -- The technical specification should clearly describe the change details. It should be detailed enough to allow language driver implementors understand whether and how they should make use of this new feature or change.
6. Alternatives -- The alternatives should describe how the issue is currently solved, how it could be solved if this change is not accepted, or what alternative designs could be implemented.
7. Impact -- Which languages are affected, how frequent this feature appears in each language, which existing drivers are affected and how they should be modified must be documented here.
8. References -- Links to language specs describing related features or other kind of related information.

### Format and Templates

BIPs must be written in Markdown format. [BIP 0](bip0-template.md) should be used as a template.

### Header

Each BIP must begin with a header that includes the following information:

* BIP number
* Title
* Author. List of authors with real names.
* Status. One of Draft \| Accepted \| Rejected
* Created. Date of creation
* Updated. Date of last update
* Target version. Target version for this feature to be implemented. It may not appear if the BIP is not about a releasable feature.

## References

This document is inspired by similar documents in other projects, specially [DEP](https://opendylan.org/proposals/) and [PEP](https://www.python.org/dev/peps/).

