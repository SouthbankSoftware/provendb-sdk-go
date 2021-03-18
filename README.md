# provendb-sdk-go
The ProvenDB SDK for Go.

## Contents
- [Getting Started](#getting-started)
- [Packages](#packages)
    - [anchor](#anchor)
    - [merkle](#merkle)

## Getting Started

Install the library with `go get github.com/SouthbankSoftware/provendb-sdk-go`.

All SDK packages reside in the `pkg` directory.

Examples are located [here](./examples).

## Packages

### anchor

`import "github.com/SouthbankSoftware/provendb-sdk-go/pkg/anchor/v1"`

The **anchor** package provides a client for our Anchor API Service.


### merkle

`import "github.com/SouthbankSoftware/provendb-sdk-go/pkg/merkle"`

The **merkle** package provides a library for you to construct your own Merkle Tree. Once constructed, you
can use the root hash of the tree and submit it to the [anchor](#anchor) service.