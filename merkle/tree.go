package merkle

import (
	"encoding/json"
	"os"
	"strings"

	"github.com/SouthbankSoftware/provendb-sdk-go/anchor"
)

// File is a complete representation of a merkle tree and it's related data.
type File struct {
	Algorithm string                `json:"algorithm"` // algorithm used to construct tree
	Proofs    []*anchor.AnchorProof `json:"proofs"`    // any associated tree proofs
	Data      [][]string            `json:"data"`      // the tree data
}

// Tree represents a single Merkle tree.
type Tree struct {
	// Algorithm used to perform tree hashing functions.
	Algorithm Hash
	// An array of proofs submitted for this tree.
	Proofs []*anchor.AnchorProof
	// The two-dimensional array of tree data, starting from the leaves (tree[0]) all the way
	// to the root (tree[tree.length - 1])
	Layers [][]string
}

// Leaf represents a single leaf in a tree.
type Leaf struct {
	Key   string
	Value string // hex encoded hash
}

// Path represents a merkle path. Either L or R will be populated so the final hash can be recalculated.
type Path struct {
	L string `json:"l,omitempty"` // the left leaf
	R string `json:"r,omitempty"` // the right leaf
}

/**
 * Converts string data to leaf.
 * @param data the data
 * @returns the leaf
 */
func toLeaf(data string) *Leaf {
	s := strings.Split(data, ":")
	return &Leaf{Key: s[0], Value: s[1]}
}

// NewTree creates a new Merkle Tree.
func NewTree(algorithm Hash, proofs []*anchor.AnchorProof, layers [][]string) *Tree {
	if proofs == nil {
		proofs = make([]*anchor.AnchorProof, 0)
	}
	if layers == nil {
		layers = make([][]string, 0)
	}
	return &Tree{
		Algorithm: algorithm,
		Proofs:    proofs,
		Layers:    layers,
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

// AddProof adds a proof for this tree.
func (t *Tree) AddProof(proof *anchor.AnchorProof) {
	t.Proofs = append(t.Proofs, proof)
}

// CountDepth returns the depth of the tree.
func (t *Tree) NDepth() int {
	return len(t.Layers) - 1
}

// CountLeaves returns the number of leaves in this tree.
func (t *Tree) NLeaves() int {
	return len((t.Layers)[0])
}

// CountNodes returns the number of nodes in this tree.
func (t *Tree) NNodes() int {
	nodes := 0
	for i := 1; i < len(t.Layers); i++ {
		nodes += len((t.Layers)[i])
	}
	return nodes
}

// CountLevels returns the number of levels in this tree.
func (t *Tree) NLevels() int {
	return len(t.Layers)
}

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

// GetLeaf returns a single leaf matching the given key.
func (t *Tree) GetLeaf(key string) *Leaf {
	for _, v := range (t.Layers)[0] {
		l := toLeaf(v)
		if l.Key == key {
			return l
		}
	}
	return nil
}

// GetLevels returns all the levels of this tree.
func (t *Tree) GetLevels() [][]string {
	return t.Layers
}

// GetLeaves returns the leaves of the tree.
func (t *Tree) GetLeaves() []*Leaf {
	leaves := make([]*Leaf, 0)
	for _, v := range (t.Layers)[0] {
		leaves = append(leaves, toLeaf(v))
	}
	return leaves
}

// GetAlgorithm returns the tree's hashing algorithm.
func (t *Tree) GetAlgorithm() Hash {
	return t.Algorithm
}

// GetLevel returns a specific level in the tree.
func (t *Tree) GetLevel(level int) []string {
	return (t.Layers)[len(t.Layers)-1-level]
}

// GetPath returns the path from a specific leaf all the way to the root hash.
// leaf must be the matching leaf value (hashed).
func (t *Tree) GetPath(key string) []*Path {
	path := make([]*Path, 0)
	index := -1
	// Find the leaf first
	leaves := (t.Layers)[0]
	for i := 0; i < len(leaves); i++ {
		leaf := toLeaf(leaves[i])
		if leaf.Key == key {
			index = i
		}
	}

	// If index is still -1, leaf not found. Return an empty path array.
	if index == -1 {
		return path
	}

	// Loop through each layer and get the index pair. Skip the root layer.
	for i := 0; i < len(t.Layers)-1; i++ {
		level := (t.Layers)[i]
		isRight := index%2 != 0

		if isRight {
			l := (level)[index-1]
			if strings.Contains(l, ":") {
				l = strings.Split(l, ":")[1]
			}
			path = append(path, &Path{L: l})
		} else {
			r := (level)[index+1]
			if strings.Contains(r, ":") {
				r = strings.Split(r, ":")[1]
			}
			path = append(path, &Path{R: r})
		}
		// Divide the index by 2 and truncate the float. Equivalent to math.Trunc()
		index = index/2 | 0
	}
	return path
}

// GetRoot returns the root hash of this tree.
func (t *Tree) GetRoot() string {
	return ((t.Layers)[len(t.Layers)-1])[0]
}

// Verify recalculates the root hash of this tree and returns the whether the calculated root hash
// matches the expected
func (t *Tree) Verify(expected string) bool {
	// Start with the leaves
	leaves := (t.Layers)[0]
	levels := build(leaves, t.Algorithm)
	return (levels[len(levels)-1])[0] == expected
}

// ValidatePath will validate the given path starting at the leaf matches the expected end result.
// func ValidatePath(path []*Path, leaf string, algorithm Hash, expected string) (bool, error) {
// 	current, err := hex.DecodeString(leaf)
// 	if err != nil {
// 		return false, err
// 	}
// 	for _, v := range path {
// 		hasher := algorithm.Hash().New()
// 		if v.L != "" {
// 			h, err := hex.DecodeString(v.L)
// 			if err != nil {
// 				return false, err
// 			}
// 			hasher.Write(h)
// 			hasher.Write(current)
// 		} else if v.R != "" {
// 			h, err := hex.DecodeString(v.R)
// 			if err != nil {
// 				return false, err
// 			}
// 			hasher.Write(current)
// 			hasher.Write(h)
// 		} else {
// 			return false, errors.New("either 'L' or 'R' must be provided in the path")
// 		}
// 		current = hasher.Sum(nil)
// 	}
// 	return hex.EncodeToString(current) == expected, nil
// }
