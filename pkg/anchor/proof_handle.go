package anchor

import (
	"context"
	"errors"
	"io"

	"github.com/SouthbankSoftware/provendb-sdk-go/genproto/anchor"
)

type ProofHandle struct {
	Proof *anchor.Proof

	client *Client
}

func NewProofHandle(proof *anchor.Proof, client *Client) (*ProofHandle, error) {
	if proof.Batch == nil {
		return nil, errors.New("proof must include batch")
	}
	return &ProofHandle{
		Proof:  proof,
		client: client,
	}, nil
}

// Update performs a call to retrieve the latest proof information.
func (p *ProofHandle) Update(ctx context.Context) (*anchor.Proof, error) {
	res, err := p.client.GetProof(ctx, p.Proof.Hash, p.Proof.BatchId)
	if err != nil {
		return nil, err
	}
	p.Proof = res
	return res, nil
}

// Subscribe subscribes to changes on this proof.
func (p *ProofHandle) Subscribe(ctx context.Context, callback func(onChange *anchor.Proof, onError error)) {
	c, err := p.client.SubscribeBatches(ctx, &SubscribeBatchesRequest{
		Filter: &BatchRequest{
			BatchId:    p.Proof.BatchId,
			AnchorType: p.Proof.AnchorType,
		},
	})
	if err != nil {
		callback(nil, err)
	}
	go func() {
		for {
			b, err := c.Recv()
			if err == io.EOF {
				return
			}
			if err != nil {
				callback(nil, err)
				return
			}
			if p.Proof.Batch.Status != b.Status {
				p, err := p.Update(ctx)
				if err != nil {
					callback(nil, err)
					return
				}
				callback(p, nil)
				// Complete func if completed or error
				if p.Batch.Status == Batch_CONFIRMED || p.Batch.Status == Batch_ERROR {
					return
				}
			}
		}
	}()
}
