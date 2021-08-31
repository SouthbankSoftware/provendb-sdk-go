# provendb-sdk-go

The ProvenDB SDK for Go.

**NOTE**: The SDK is is in development and not recommended for production use.  For any bugs, please raise an [issue](https://github.com/SouthbankSoftware/provendb-sdk-node/issues).

The use of our services requires an API key. Visit [provendb.com](https://provendb.com) to sign in/register and create one.

## Installation

`go get github.com/SouthbankSoftware/provendb-sdk-go`

## Libraries

| Name | Description | Import |
| :--- | :---------- | :----- |
| [anchor](./anchor) | The ProvenDB Anchor client. | `import "github.com/SouthbankSoftware/provendb-sdk-go/anchor"` |
| [merkle](./merkle) | A merkle tree library. | `import "github.com/SouthbankSoftware/provendb-sdk-go/merkle"` |

## Examples

## Hello World!

This Hello, World example uses both the [anchor](./anchor) and [merkle](./merkle) libraries to generate
a merkle tree and submit the tree's root hash to Hedera via the ProvenDB Anchor service.

```go
package main

import (
    "github.com/SouthbankSoftware/provendb-sdk-go/anchor"
    "github.com/SouthbankSoftware/provendb-sdk-go/merkle"
)

func main() {
    // Create the new builder and add your data.
    builder := NewBuilder(merkle.SHA256)
    builder.Add("key1", []byte("Hello, "))
    builder.Add("key2", []byte("World, !"))

    // Construct the tree.
    tree := builder.Build()

    // Create a new anchor client using your credentials
    client := anchor.Connect(anchor.WithCredentials("YOUR_API_KEY"));

    // Submit your proof.
    proof, err := client.SubmitProof(tree.GetRoot(), 
        anchor.SubmitProofWithAnchorType(anchor.Anchor.Type.HEDERA_MAINNET), // Optional. Add your anchor type.
        anchor.SubmitProofWithAwaitConfirmed(true)); // Optional. Return proof only when the proof is confirmed.

    if err != nil {
        // handle error
        panic(err)
    }

    // Add your proof to the tree.
    tree.AddProof(proof);

    // Export it to file
    tree.Export("./merkle.json");
}
```

Congratulations, you have successfully created, anchored and exported your first `Hello, World!` proof!

## Documentation

Full API documentation not yet available.
