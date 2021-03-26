package anchor

import (
	context "context"
	"io"
	"os"

	grpc "google.golang.org/grpc"
	"google.golang.org/protobuf/types/known/emptypb"
)

// Client holds the grpc anchor service client and provides simplified functionality to
// perform anchor operations.
type Client struct {
	conn   *grpc.ClientConn
	anchor AnchorServiceClient
}

// ClientOptions represents the anchor options.
type ClientOptions struct {
	// The anchor service address.
	Address string
	// The credential for authentication.
	Credentials string
	// Establishes an insecure connection.
	Insecure bool
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

func WithInsecure(insecure bool) ClientOption {
	return func(o *ClientOptions) {
		o.Insecure = insecure
	}
}

// Connect creates a new anchor client and performs all the grpc connections.
func Connect(opts ...ClientOption) (*Client, error) {
	const (
		defaultAddress    = "anchor.proofable.io:443"
		defaultInsecure   = false
		defaultAnchorType = Anchor_ETH
	)
	// Default credentials is lookup from environment.
	credential := os.Getenv("PROVENDB_ANCHOR_CREDENTIALS")
	o := &ClientOptions{
		Address:     defaultAddress,
		Credentials: credential,
		Insecure:    defaultInsecure,
	}

	for _, opt := range opts {
		opt(o)
	}
	var dialOpts []grpc.DialOption
	if o.Credentials != "" {
		dialOpts = append(dialOpts, grpc.WithPerRPCCredentials(NewServiceCredentials(o.Credentials, !o.Insecure)))
	}
	if o.Insecure {
		dialOpts = append(dialOpts, grpc.WithInsecure())
	}
	conn, err := grpc.Dial(o.Address, dialOpts...)
	if err != nil {
		return nil, err
	}
	service := NewAnchorServiceClient(conn)
	return &Client{
		conn:   conn,
		anchor: service,
	}, nil
}

// Close the anchor connection gracefully.
func (c *Client) Close() error {
	return c.conn.Close()
}

// GetAnchors will retrieve all the available anchors.
func (c *Client) GetAnchors(ctx context.Context) ([]*Anchor, error) {
	res, err := c.anchor.GetAnchors(ctx, &emptypb.Empty{})
	if err != nil {
		return nil, err
	}
	anchors := make([]*Anchor, 0)
	for {
		a, err := res.Recv()
		if err == io.EOF {
			return anchors, nil
		}
		if err != nil {
			return nil, err
		}
		anchors = append(anchors, a)
	}
}

// GetAnchor will retreive information about a single anchor.
func (c *Client) GetAnchor(ctx context.Context, anchorType Anchor_Type) (*Anchor, error) {
	return c.anchor.GetAnchor(ctx, &AnchorRequest{
		Type: anchorType,
	})
}

// GetBatch retrieves a single batch information.
func (c *Client) GetBatch(ctx context.Context, batchId string, anchorType Anchor_Type) (*Batch, error) {
	return c.anchor.GetBatch(ctx, &BatchRequest{
		BatchId:    batchId,
		AnchorType: anchorType,
	})
}

// GetProof retrieves a proof matching the given hash and batch ID.
func (c *Client) GetProof(ctx context.Context, hash string, batchId string) (*Proof, error) {
	res, err := c.anchor.GetProof(ctx, &ProofRequest{
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
	AnchorType Anchor_Type
	// Whether to skip batching.
	SkipBatching bool
	// Format
	Format Proof_Format
}

type SubmitProofOption func(o *SubmitProofOptions)

func SubmitProofWithAnchorType(anchorType Anchor_Type) SubmitProofOption {
	return func(o *SubmitProofOptions) {
		o.AnchorType = anchorType
	}
}

func SubmitProofWithSkipBatching(skip bool) SubmitProofOption {
	return func(o *SubmitProofOptions) {
		o.SkipBatching = skip
	}
}

func SubmitProofWithFormat(format Proof_Format) SubmitProofOption {
	return func(o *SubmitProofOptions) {
		o.Format = format
	}
}

// SubmitProof submits a new proof to the anchor service.
func (c *Client) SubmitProof(ctx context.Context, hash string, opts ...SubmitProofOption) (*Proof, error) {
	// Set default options
	o := &SubmitProofOptions{
		AnchorType:   Anchor_ETH,
		SkipBatching: false,
		Format:       Proof_CHP_PATH,
	}
	for _, opt := range opts {
		opt(o)
	}
	// Submit the request
	req := &SubmitProofRequest{
		Hash:         hash,
		AnchorType:   o.AnchorType,
		SkipBatching: o.SkipBatching,
		Format:       o.Format,
	}
	res, err := c.anchor.SubmitProof(ctx, req)
	if err != nil {
		return nil, err
	}
	return res, nil
}

func (c *Client) SubscribeBatch(ctx context.Context, batchId string, anchorType Anchor_Type, callback func(batch *Batch, err error)) {
	c.anchor.SubscribeBatches(ctx, &SubscribeBatchesRequest{
		Filter: &BatchRequest{
			BatchId:    batchId,
			AnchorType: anchorType,
		},
	})
}

func (c *Client) SubscribeBatches(ctx context.Context, callback func(batch *Batch, err error)) {
	// TODO
	return
}

// Subsribe proof will listen for changes to the proof and return either the updated proof, or an error.
// Function will complete once proof status returned is either CONFIRMED or ERROR, or context expired.
func (c *Client) SubscribeProof(ctx context.Context, p *Proof, callback func(proof *Proof, err error)) {
	// TODO check context expiration
	res, err := c.anchor.SubscribeBatches(ctx, &SubscribeBatchesRequest{
		Filter: &BatchRequest{
			BatchId:    p.BatchId,
			AnchorType: p.AnchorType,
		},
	})
	if err != nil {
		callback(nil, err)
	}
	go func() {
		for {
			b, err := res.Recv()
			if err == io.EOF {
				return
			}
			if err != nil {
				callback(nil, err)
				return
			}
			if p.BatchStatus != b.Status {
				proof, err := c.GetProof(ctx, p.Hash, p.BatchId)
				if err != nil {
					callback(nil, err)
					return
				}
				callback(proof, nil)
				// Complete func if completed or error
				if proof.Batch.Status == Batch_CONFIRMED || proof.Batch.Status == Batch_ERROR {
					return
				}
			}
		}
	}()
}
