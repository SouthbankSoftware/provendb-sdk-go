package merkle

import (
	"encoding/hex"
	"fmt"
	"hash"
)

// Builder represents a merkle tree builder.
type Builder struct {
	algorithm   Hash     // the algorithm
	leaves      []string // the leaves
	description string   // a description
}

// Writer for writing streams of data to the tree.
type Writer struct {
	builder *Builder
	hasher  hash.Hash
}

// Write implements the write method of the io.Writer interface.
func (w *Writer) Write(p []byte) (n int, err error) {
	return w.hasher.Write(p)
}

// Close completes the writing and adds the hashed sum to the tree.
func (w *Writer) Close() *Builder {
	// Add the hash and release the lock
	w.builder.add(w.hasher.Sum(nil), false)
	return w.builder
}

// NewBuilder creates a new merkle tree builder.
func NewBuilder(algorithm Hash) *Builder {
	return &Builder{
		algorithm: algorithm,
		leaves:    []string{},
	}
}

// Description sets the description to the final tree.
func (b *Builder) Description(d string) *Builder {
	b.description = d
	return b
}

func (b *Builder) add(data []byte, doHash bool) *Builder {
	v := data
	if doHash {
		hasher := b.algorithm.Hash().New()
		hasher.Write(data)
		v = hasher.Sum(nil)
	}
	b.leaves = append(b.leaves, hex.EncodeToString(v))
	return b
}

// Add adds data to the builder. Whatever data is passed here will be hashed
// with the algorithm specified in the builder.
func (b *Builder) Add(data []byte) *Builder {
	return b.add(data, true)
}

// AddHash will add the hash to the builder and will not re-hash it. If the value is not
// a hex string, and not the same size as the algorithm size, it will panic.
func (b *Builder) AddHash(h string) *Builder {
	d, err := hex.DecodeString(h)
	if err != nil {
		panic(err)
	}
	if l := len(d); l != b.algorithm.Hash().Size() {
		panic(fmt.Sprintf("invalid hex size '%d' for algorithm '%s'", l, b.algorithm))
	}
	b.add(d, false)
	return b
}

// AddBatch adds a batch of items to the tree.
func (b *Builder) AddBatch(data *[][]byte) *Builder {
	for _, v := range *data {
		b.add(v, true)
	}
	return b
}

// Writer returns a writer object for writing streams of bytes.
func (b *Builder) Writer() *Writer {
	return &Writer{
		builder: b,
		hasher:  b.algorithm.Hash().New(),
	}
}

// Build constructs the tree and returns the tree struct.
func (b *Builder) Build() *Tree {
	levels := Build(b.leaves, b.algorithm)
	return NewTree(b.algorithm, nil, levels)
}

// Build constructs an entire tree's levels from the single level provided.
func Build(leaves []string, algorithm Hash) [][]string {
	levels := make([][]string, 0)
	levels = append(levels, leaves)
	level := leaves
	for len(level) > 1 {
		level = BuildLevel(level, algorithm)
		levels = append(levels, level)
	}
	return levels
}

// BuildLevel constructs the next level from the given layer.
func BuildLevel(level []string, algorithm Hash) []string {
	l := make([]string, 0)
	// Loop through the nodes with increments of 2 (pairs)
	for i := 0; i < len(level); i += 2 {
		// If we have an odd node, we promote it.
		if i+1 == len(level) {
			l = append(l, level[i])
		} else {
			hasher := algorithm.Hash().New()
			left, err := hex.DecodeString(level[i])
			if err != nil {
				panic("string is not hex")
			}
			right, err := hex.DecodeString(level[i+1])
			if err != nil {
				panic("string is not hex")
			}
			hasher.Write(left)
			hasher.Write(right)
			l = append(l, fmt.Sprintf("%x", hasher.Sum(nil)))
		}
	}
	return l
}
