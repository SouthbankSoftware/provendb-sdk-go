package anchor

import (
	"context"

	"github.com/SouthbankSoftware/provendb-sdk-go/genproto/anchor/v1"
	"google.golang.org/grpc"
)

// import "github.com/SouthbankSoftware/proofable/pkg/protos/anchor"

// ProofHandle represents a handle on a single proof.
type ProofHandle struct {
	proof         *anchor.Proof
	err           error  // any errors
	serverAddress string // the anchor server address
	token         string // the auth token
}

// Proof wraps the proto.Proof from anchor service struct and provides some extra functionality.
type Proof struct {
	AnchorType string      `json:"anchorType"`
	Format     string      `json:"format"`
	Metadata   interface{} `json:"metadata"`
	Proof      interface{} `json:"proof"`
}

// MarshalJSON implements the json.Encoder interface.
func (p *Proof) MarshalJSON() ([]byte, error) {
	return nil, nil
}

// UnmarshalJSON implements the json.Decoder interface.
func (p *Proof) UnmarshalJSON(data []byte) error {
	return nil
}

// DecodeProof will decode the base64 encoded proof data.
func (p *Proof) DecodeProof() (b []byte, e error) {
	return
}

// NewProofHandle creates a new proof handle.
func NewProofHandle(proof *anchor.Proof) *ProofHandle {
	return &ProofHandle{proof: proof}
}

// Error returns the current error status, or nil if no error.
func (h *ProofHandle) Error() error {
	return h.err
}

// Update polls the anchor service and retrieves the current proof.
func (h *ProofHandle) Update(ctx context.Context) (*anchor.Proof, error) {
	conn, err := grpc.Dial(h.serverAddress)
	if err != nil {
		return nil, err
	}
	client := anchor.NewAnchorServiceClient(conn)
	h.proof, h.err = client.GetProof(ctx, &anchor.ProofRequest{
		Hash:      h.proof.Hash,
		BatchId:   h.proof.BatchId,
		WithBatch: true,
	})
	if h.err != nil {
		return nil, err
	}
	return h.proof, nil
}
