package merkle

import (
	"encoding/hex"
	"fmt"
	"hash"
	"strings"
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
	key     string
	hasher  hash.Hash
}

// Write implements the write method of the io.Writer interface.
func (w *Writer) Write(p []byte) (n int, err error) {
	return w.hasher.Write(p)
}

// Close completes the writing and adds the hashed sum to the tree.
func (w *Writer) Close() *Builder {
	// Add the hash and release the lock
	w.builder.add(w.key, w.hasher.Sum(nil), false)
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

func (b *Builder) add(key string, value []byte, doHash bool) *Builder {
	v := value
	if doHash {
		hasher := b.algorithm.Hash().New()
		hasher.Write(value)
		v = hasher.Sum(nil)
		b.leaves = append(b.leaves, key+":"+hex.EncodeToString(v))
	} else {
		b.leaves = append(b.leaves, key+":"+string(v))
	}
	return b
}

// Add adds data to the builder. Whatever data is passed here will be hashed
// with the algorithm specified in the builder.
func (b *Builder) Add(key string, value []byte) *Builder {
	return b.add(key, value, true)
}

// AddRaw adds data to the builder but will not hash the provided data.
func (b *Builder) AddRaw(key string, value string) *Builder {
	return b.add(key, []byte(value), false)
}

// AddBatch adds a batch of items to the tree.
func (b *Builder) AddBatch(data []*struct {
	Key   string
	Value []byte
}) *Builder {
	for _, l := range data {
		b.add(l.Key, l.Value, true)
	}
	return b
}

// Writer returns a writer object for writing streams of bytes.
func (b *Builder) Writer(key string) *Writer {
	return &Writer{
		builder: b,
		key:     key,
		hasher:  b.algorithm.Hash().New(),
	}
}

// Build constructs the tree and returns the tree struct.
func (b *Builder) Build() *Tree {
	levels := build(b.leaves, b.algorithm)
	return NewTree(b.algorithm, nil, levels)
}

// Build constructs an entire tree's levels from the single level provided.
func build(leaves []string, algorithm Hash) [][]string {
	levels := make([][]string, 0)
	levels = append(levels, leaves)
	level := leaves
	for len(level) > 1 {
		level = buildLevel(level, algorithm)
		levels = append(levels, level)
	}
	return levels
}

// BuildLevel constructs the next level from the given layer.
func buildLevel(level []string, algorithm Hash) []string {
	l := make([]string, 0)
	// Loop through the nodes with increments of 2 (pairs)
	for i := 0; i < len(level); i += 2 {
		// If we have an odd node, we promote it.
		if i+1 == len(level) {
			if strings.Contains(level[i], ":") {
				s := strings.Split(level[i], ":")
				l = append(l, s[1])
			} else {
				l = append(l, level[i])
			}
		} else {
			hasher := algorithm.Hash().New()
			left := level[i]
			right := level[i+1]
			if strings.Contains(level[i], ":") {
				left = strings.Split(level[i], ":")[1]
				right = strings.Split(level[i+1], ":")[1]
			}
			leftBytes, err := hex.DecodeString(left)
			if err != nil {
				panic("string is not hex")
			}
			rightBytes, err := hex.DecodeString(right)
			if err != nil {
				panic("string is not hex")
			}
			hasher.Write(leftBytes)
			hasher.Write(rightBytes)
			l = append(l, fmt.Sprintf("%x", hasher.Sum(nil)))
		}
	}
	return l
}
