package anchor

import (
	"context"

	"github.com/SouthbankSoftware/provendb-sdk-go/genproto/anchor/v1"
)

// Client represents an anchor service client.
type Client struct {
	service anchor.AnchorServiceClient
}

// Options represents the anchor options.
type Options struct {
	Address string
	Token   string
	Secure  bool
}

// NewClient creates a new anchor client.
func NewClient(options *Options) (*Client, error) {
	return &Client{}, nil
}

// Proof returns a handle on a specific proof.
func (c *Client) Proof(p *anchor.Proof) *ProofHandle {
	return nil
}

// GetAnchors will retrieve all the available anchors.
func (c *Client) GetAnchors() ([]anchor.Anchor, error) {
	return nil, nil
}

// GetAnchor will retreive information about a single anchor.
func (c *Client) GetAnchor(t anchor.Anchor_Type) (*anchor.Anchor, error) {
	return nil, nil
}

// GetBatch retrieves a single batch information.
func (c *Client) GetBatch(ctx context.Context, req *anchor.BatchRequest) (*anchor.Batch, error) {
	return nil, nil
}

// SubmitProof submits a new proof to the anchor service.
func (c *Client) SubmitProof(ctx context.Context, req *anchor.SubmitProofRequest) (*ProofHandle, error) {
	return &ProofHandle{}, nil
}
