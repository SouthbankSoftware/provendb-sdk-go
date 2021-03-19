package anchor

import (
	"context"

	"github.com/SouthbankSoftware/provendb-sdk-go/genproto/anchor"
)

// ProofHandle represents a handle on a single proof.
type ProofHandle struct {
	Proof  *Proof
	err    error // any errors
	client anchor.AnchorServiceClient
}

// Error returns the current error status, or nil if no error.
func (h *ProofHandle) Error() error {
	return h.err
}

// Update polls the anchor service and retrieves the current proof.
func (h *ProofHandle) Update(ctx context.Context) (*Proof, error) {
	_, err := h.client.GetProof(ctx, &anchor.ProofRequest{
		Hash:      h.Proof.Hash,
		BatchId:   h.Proof.BatchID,
		WithBatch: true,
	})
	if h.err != nil {
		return nil, err
	}
	return h.Proof, nil
}

// Subscribe to changes/updates on this proof.
func (h *ProofHandle) Subscribe(ctx context.Context, callback func(p Proof, e error)) {}
