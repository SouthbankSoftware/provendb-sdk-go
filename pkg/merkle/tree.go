package merkle

import (
	"encoding/hex"
	"encoding/json"
	"errors"
	"os"

	"github.com/SouthbankSoftware/provendb-sdk-go/pkg/anchor"
)

// Tree represents a single Merkle tree.
type Tree struct {
	// Description is a custom description of this tree.
	Description string `json:"description"`
	// Algorithm used to perform tree hashing functions.
	Algorithm Hash `json:"algorithm"`
	// An array of proofs submitted for this tree.
	Proofs []Proof `json:"proofs"`
	// The two-dimensional array of tree data, starting from the leaves (tree[0]) all the way
	// to the root (tree[tree.length - 1])
	Levels [][]string `json:"levels"`
}

// Path represents a merkle path. Either L or R will be populated so the final hash can be recalculated.
type Path struct {
	L string `json:"l,omitempty"` // the left leaf
	R string `json:"r,omitempty"` // the right leaf
}

// NewTree creates a new Merkle Tree.
func NewTree(algorithm Hash, proofs []Proof, levels [][]string) *Tree {
	if proofs == nil {
		proofs = make([]Proof, 0)
	}
	if levels == nil {
		levels = make([][]string, 0)
	}
	return &Tree{
		Algorithm: algorithm,
		Proofs:    proofs,
		Levels:    levels,
	}
}

// NewTreeFromFile creates a new tree from an existing file.
func NewTreeFromFile(path string) (*Tree, error) {
	file, err := os.Open(path)
	if err != nil {
		return nil, err
	}
	tree := new(Tree)
	if err := json.NewDecoder(file).Decode(tree); err != nil {
		return nil, err
	}
	return tree, nil
}

// AddProof adds a proof for this tree. If the proof provided is an anchor.Proof, this library will
// perform the necessary functions to convert it to a generic proof.
func (t *Tree) AddProof(proof interface{}) {
	switch proof.(type) {
	case anchor.Proof:
		p := proof.(anchor.Proof)
		t.Proofs = append(t.Proofs, FromAnchorProof(&p))
	case Proof:
		t.Proofs = append(t.Proofs, proof.(Proof))
	default:
		panic("proof not recognised")
	}
}

// CountDepth returns the depth of the tree.
func (t *Tree) CountDepth() int {
	return len(t.Levels) - 1
}

// CountLeaves returns the number of leaves in this tree.
func (t *Tree) CountLeaves() int {
	return len((t.Levels)[0])
}

// CountNodes returns the number of nodes in this tree.
func (t *Tree) CountNodes() int {
	nodes := 0
	for i := 1; i < len(t.Levels); i++ {
		nodes += len((t.Levels)[i])
	}
	return nodes
}

// CountLevels returns the number of levels in this tree.
func (t *Tree) CountLevels() int {
	return len(t.Levels)
}

// type CreateProofOptions struct {
// 	// The address of the anchor service.
// 	Address string
// 	// The anchor type to use.
// 	AnchorType anchor.Anchor_Type
// 	// The credentials to use.
// 	Credentials string
// 	// The proof format.
// 	Format anchor.Proof_Format
// 	// Enable a secure connection. Default true.
// 	Secure bool
// 	// SkipBatching
// 	SkipBatching bool
// }

// type CreateProofOption func(opts *CreateProofOptions)

// func CreateProofWithAddress(address string) CreateProofOption {
// 	return func(opts *CreateProofOptions) {
// 		opts.Address = address
// 	}
// }

// func CreateProofWithAnchorType(anchorType proto.Anchor_Type) CreateProofOption {
// 	return func(opts *CreateProofOptions) {
// 		opts.AnchorType = anchorType
// 	}
// }

// func CreateProofWithCredentials(credentials string) CreateProofOption {
// 	return func(opts *CreateProofOptions) {
// 		opts.Credentials = credentials
// 	}
// }

// func CreateProofWithSkipBatching(skipBatching bool) CreateProofOption {
// 	return func(opts *CreateProofOptions) {
// 		opts.SkipBatching = skipBatching
// 	}
// }

// func CreateProofWithSecure(secure bool) CreateProofOption {
// 	return func(opts *CreateProofOptions) {
// 		opts.Secure = secure
// 	}
// }

// func CreateProofWithFormat(format proto.Proof_Format) CreateProofOption {
// 	return func(opts *CreateProofOptions) {
// 		opts.Format = format
// 	}
// }

// // CreateProof creates a proof by submitting the root hash of this tree.
// func (t *Tree) CreateProof(ctx context.Context, anchor *anchor.AnchorServiceClient, opts ...CreateProofOptions) (*anchor.ProofHandle, error) {
// 	options := &CreateProofOptions{
// 		Address:      "anchor.proofable.io:443",
// 		AnchorType:   anchor.Anchor_ETH,
// 		Credentials:  os.Getenv("PROVENDB_ANCHOR_CREDENTIALS"),
// 		Secure:       true,
// 		Format:       anchor.Proof_CHP_PATH,
// 		SkipBatching: false,
// 	}

// 	for _, opt := range opts {
// 		opt(options)
// 	}

// 	var dialOpts []grpc.DialOption
// 	if !options.Secure {
// 		dialOpts = append(dialOpts, grpc.WithInsecure())
// 	}
// 	conn, err := grpc.Dial(options.Address, dialOpts...)
// 	if err != nil {
// 		return nil, err
// 	}
// 	defer conn.Close()
// 	service := anchor.NewAnchorServiceClient(conn)

// 	p, err := service.SubmitProof(ctx, &anchor.SubmitProofRequest{
// 		Hash:       t.Root(),
// 		AnchorType: options.AnchorType,
// 		Format:     options.Format,
// 	})
// 	if err != nil {
// 		return nil, err
// 	}
// 	return anchor.NewProofHandle(p, service)
// }

// Export exports this tree to file.
func (t *Tree) Export(path string) error {
	file, err := os.Create(path)
	defer file.Close()
	if err != nil {
		return err
	}
	e := json.NewEncoder(file)
	if err := e.Encode(t); err != nil {
		return err
	}
	return nil
}

// Leaves returns the leaves of the tree.
func (t *Tree) Leaves() []string {
	return (t.Levels)[0]
}

// Level returns a specific level in the tree.
func (t *Tree) Level(level int) []string {
	return (t.Levels)[len(t.Levels)-1-level]
}

// Path returns the path from a specific leaf all the way to the root hash.
// leaf must be the matching leaf value (hashed).
func (t *Tree) Path(leaf string) []*Path {
	path := make([]*Path, 0)
	index := -1
	// Find the leaf first
	leaves := (t.Levels)[0]
	for i := 0; i < len(leaves); i++ {
		if (leaves)[i] == leaf {
			index = i
		}
	}

	// If index is still -1, leaf not found. Return an empty path array.
	if index == -1 {
		return path
	}

	// Loop through each layer and get the index pair. Skip the root layer.
	for i := 0; i < len(t.Levels)-1; i++ {
		level := (t.Levels)[i]
		isRight := index%2 != 0

		if isRight {
			l := (level)[index-1]
			path = append(path, &Path{L: l})
		} else {
			r := (level)[index+1]
			path = append(path, &Path{R: r})
		}
		// Divide the index by 2 and truncate the float. Equivalent to math.Trunc()
		index = index/2 | 0
	}
	return path
}

// Root returns the root hash of this tree.
func (t *Tree) Root() string {
	return ((t.Levels)[len(t.Levels)-1])[0]
}

// Verify recalculates the root hash of this tree and returns the whether the calculated root hash
// matches the expected
func (t *Tree) Verify(expected string) bool {
	// Start with the leaves
	leaves := (t.Levels)[0]
	levels := Build(leaves, t.Algorithm)
	return (levels[len(levels)-1])[0] == expected
}

// ValidatePath will validate the given path starting at the leaf matches the expected end result.
func ValidatePath(path []*Path, leaf string, algorithm Hash, expected string) (bool, error) {
	current, err := hex.DecodeString(leaf)
	if err != nil {
		return false, err
	}
	for _, v := range path {
		hasher := algorithm.Hash().New()
		if v.L != "" {
			h, err := hex.DecodeString(v.L)
			if err != nil {
				return false, err
			}
			hasher.Write(h)
			hasher.Write(current)
		} else if v.R != "" {
			h, err := hex.DecodeString(v.R)
			if err != nil {
				return false, err
			}
			hasher.Write(current)
			hasher.Write(h)
		} else {
			return false, errors.New("either 'L' or 'R' must be provided in the path")
		}
		current = hasher.Sum(nil)
	}
	return hex.EncodeToString(current) == expected, nil
}
