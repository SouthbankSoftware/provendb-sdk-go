package anchor

import (
	context "context"
	"os"

	"github.com/SouthbankSoftware/provendb-sdk-go/genproto/anchor"
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
		defaultAnchorType = anchor.Anchor_ETH
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
	if o.Credentials != "" {
		dialOpts = append(dialOpts, grpc.WithPerRPCCredentials(NewServiceCredentials(o.Credentials, o.Secure)))
	}
	if !o.Secure {
		dialOpts = append(dialOpts, grpc.WithInsecure())
	}
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

// GetAnchors will retrieve all the available anchors.
func (c *Client) GetAnchors() ([]anchor.Anchor, error) {
	return nil, nil
}

// GetAnchor will retreive information about a single anchor.
func (c *Client) GetAnchor(t anchor.Anchor_Type) (*anchor.Anchor, error) {
	return nil, nil
}

// GetBatch retrieves a single batch information.
func (c *Client) GetBatch(ctx context.Context, id string, anchorType anchor.Anchor_Type) (*anchor.Batch, error) {
	return nil, nil
}

// GetProof retrieves a proof matching the given hash and batch ID.
func (c *Client) GetProof(ctx context.Context, hash string, batchId string) (*anchor.Proof, error) {
	res, err := c.service.GetProof(ctx, &anchor.ProofRequest{
		Hash:      hash,
		BatchId:   batchId,
		WithBatch: true,
	})
	if err != nil {
		return nil, err
	}
	return res, nil
}

// SubmitProofOptions options.
type SubmitProofOptions struct {
	// The anchor type.
	AnchorType anchor.Anchor_Type
	// Whether to skip batching.
	SkipBatching bool
	// Format
	Format anchor.Proof_Format
}

type SubmitProofOption func(o *SubmitProofOptions)

func SubmitProofWithAnchorType(anchorType anchor.Anchor_Type) SubmitProofOption {
	return func(o *SubmitProofOptions) {
		o.AnchorType = anchorType
	}
}

func SubmitProofWithSkipBatching(skip bool) SubmitProofOption {
	return func(o *SubmitProofOptions) {
		o.SkipBatching = skip
	}
}

func SubmitProofWithFormat(format anchor.Proof_Format) SubmitProofOption {
	return func(o *SubmitProofOptions) {
		o.Format = format
	}
}

// SubmitProof submits a new proof to the anchor service.
func (c *Client) SubmitProof(ctx context.Context, hash string, opts ...SubmitProofOption) (*anchor.Proof, error) {
	// Prepare the options
	o := &SubmitProofOptions{
		AnchorType:   anchor.Anchor_ETH,
		SkipBatching: false,
		Format:       anchor.Proof_CHP_PATH,
	}
	for _, opt := range opts {
		opt(o)
	}
	// Submit the request
	req := &anchor.SubmitProofRequest{
		Hash:         hash,
		AnchorType:   o.AnchorType,
		SkipBatching: o.SkipBatching,
		Format:       o.Format,
	}
	res, err := c.service.SubmitProof(ctx, req)
	if err != nil {
		return nil, err
	}
	return res, nil
}

func (c *Client) SubscribeBatch(ctx context.Context, batchId string, anchorType anchor.Anchor_Type) (anchor.AnchorService_SubscribeBatchesClient, error) {
	return c.service.SubscribeBatches(ctx, &anchor.SubscribeBatchesRequest{
		Filter: &anchor.BatchRequest{
			BatchId:    batchId,
			AnchorType: anchorType,
		},
	})
}
