package anchor

import (
	"context"
	"os"

	"github.com/SouthbankSoftware/provendb-sdk-go/genproto/anchor/v1"
	"google.golang.org/grpc"
)

// Client represents an anchor service client.
type Client struct {
	service anchor.AnchorServiceClient
}

// ClientOptions represents the anchor options.
type ClientOptions struct {
	// The anchor service address.
	Address string
	// The credential for authentication.
	Credentials string
	// Establishes a secure connection.
	Secure bool
}

// ClientOption func.
type ClientOption func(*ClientOptions)

func WithAddress(address string) ClientOption {
	return func(o *ClientOptions) {
		o.Address = address
	}
}

func WithCredentials(credentials string) ClientOption {
	return func(o *ClientOptions) {
		o.Credentials = credentials
	}
}

func WithSecure(secure bool) ClientOption {
	return func(o *ClientOptions) {
		o.Secure = secure
	}
}

// NewClient creates a new anchor client.
func NewClient(opts ...ClientOption) (*Client, error) {
	const (
		defaultAddress    = "anchor.proofable.io:443"
		defaultSecure     = true
		defaultAnchorType = EthereumTestnet
	)
	// Default credentials is lookup from environment.
	credential := os.Getenv("PROVENDB_ANCHOR_CREDENTIALS")
	o := &ClientOptions{
		Address:     defaultAddress,
		Credentials: credential,
		Secure:      defaultSecure,
	}

	for _, opt := range opts {
		opt(o)
	}
	var dialOpts []grpc.DialOption
	conn, err := grpc.Dial(o.Address, dialOpts...)
	if err != nil {
		return nil, err
	}
	defer conn.Close()
	service := anchor.NewAnchorServiceClient(conn)
	return &Client{
		service: service,
	}, nil
}

// Proof returns a handle on a specific proof.
func (c *Client) Proof(p *Proof) *ProofHandle {
	return &ProofHandle{
		Proof:  p,
		client: c.service,
	}
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
func (c *Client) GetBatch(ctx context.Context, id string, anchorType AnchorType) (*anchor.Batch, error) {
	return nil, nil
}

// SubmitProofOptions options.
type SubmitProofOptions struct {
	// The anchor type.
	AnchorType AnchorType
	// Whether to skip batching.
	SkipBatching bool
	// Format
	Format ProofFormat
}

type SubmitProofOption func(o *SubmitProofOptions)

func SubmitProofWithAnchorType(anchorType AnchorType) SubmitProofOption {
	return func(o *SubmitProofOptions) {
		o.AnchorType = anchorType
	}
}

func SubmitProofWithSkipBatching(skip bool) SubmitProofOption {
	return func(o *SubmitProofOptions) {
		o.SkipBatching = skip
	}
}

func SubmitProofWithFormat(format ProofFormat) SubmitProofOption {
	return func(o *SubmitProofOptions) {
		o.Format = format
	}
}

// SubmitProof submits a new proof to the anchor service.
func (c *Client) SubmitProof(ctx context.Context, hash string, opts ...SubmitProofOption) (*ProofHandle, error) {
	return &ProofHandle{}, nil
}
