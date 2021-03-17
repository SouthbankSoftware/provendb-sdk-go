package merkle

import (
	_ "crypto/sha256"
	"reflect"
	"testing"
)

// Variables for trees with 16 leaves.
var (
	batch16 = [][]byte{
		[]byte("a"),
		[]byte("b"),
		[]byte("c"),
		[]byte("d"),
		[]byte("e"),
		[]byte("f"),
		[]byte("g"),
		[]byte("h"),
		[]byte("i"),
		[]byte("j"),
		[]byte("k"),
		[]byte("l"),
		[]byte("m"),
		[]byte("n"),
		[]byte("o"),
		[]byte("p"),
	}

	batch16root = abcdefghijklmnop

	batch16pathA = []*Path{
		{R: b},
		{R: cd},
		{R: efgh},
		{R: ijklmnop},
	}

	batch16pathB = []*Path{
		{L: a},
		{R: cd},
		{R: efgh},
		{R: ijklmnop},
	}

	batch16pathC = []*Path{
		{R: d},
		{L: ab},
		{R: efgh},
		{R: ijklmnop},
	}

	batch16pathD = []*Path{
		{L: c},
		{L: ab},
		{R: efgh},
		{R: ijklmnop},
	}

	batch16pathE = []*Path{
		{R: f},
		{R: gh},
		{L: abcd},
		{R: ijklmnop},
	}

	batch16pathF = []*Path{
		{L: e},
		{R: gh},
		{L: abcd},
		{R: ijklmnop},
	}

	batch16pathG = []*Path{
		{R: h},
		{L: ef},
		{L: abcd},
		{R: ijklmnop},
	}

	batch16pathH = []*Path{
		{L: g},
		{L: ef},
		{L: abcd},
		{R: ijklmnop},
	}

	batch16pathI = []*Path{
		{R: j},
		{R: kl},
		{R: mnop},
		{L: abcdefgh},
	}

	batch16pathJ = []*Path{
		{L: i},
		{R: kl},
		{R: mnop},
		{L: abcdefgh},
	}

	batch16pathK = []*Path{
		{R: l},
		{L: ij},
		{R: mnop},
		{L: abcdefgh},
	}

	batch16pathL = []*Path{
		{L: k},
		{L: ij},
		{R: mnop},
		{L: abcdefgh},
	}

	batch16pathM = []*Path{
		{R: n},
		{R: op},
		{L: ijkl},
		{L: abcdefgh},
	}

	batch16pathN = []*Path{
		{L: m},
		{R: op},
		{L: ijkl},
		{L: abcdefgh},
	}

	batch16pathO = []*Path{
		{R: p},
		{L: mn},
		{L: ijkl},
		{L: abcdefgh},
	}

	batch16pathP = []*Path{
		{L: o},
		{L: mn},
		{L: ijkl},
		{L: abcdefgh},
	}
)

// Checks for equality between both Path slices.
func equal(t *testing.T, exp *[]*Path, act *[]*Path, leaf string) {
	if !reflect.DeepEqual(exp, act) {
		t.Fatalf("path mismatch for '%s'", leaf)
	}

}

// Validates the path by ensuring the final calculated hash DOES match the path.
func validate(t *testing.T, path *[]*Path, root string, leaf string) {
	ok, err := ValidatePath(*path, leaf, SHA256, root)
	if err != nil {
		t.Fatal(err)
	}
	if !ok {
		t.Fatalf("root mismatch for '%s'", leaf)
	}
}

// Invalidates the path by ensuring the final calculated hash does NOT match.
func invalidate(t *testing.T, path []*Path, root string, leaf string) {
	ok, err := ValidatePath(path, leaf, SHA256, root)
	if err != nil {
		t.Fatal(err)
	}
	if ok {
		t.Fatalf("root should have been different for '%s'", a)
	}
}

func TestTree_Path(t *testing.T) {
	builder := NewBuilder(SHA256)
	builder.AddBatch(&batch16)
	tree := builder.Build()

	// Path of 'a'
	exp := batch16pathA
	act := tree.Path(a)
	equal(t, &exp, &act, a)

	// Path of 'b'
	exp = batch16pathB
	act = tree.Path(b)
	equal(t, &exp, &act, b)

	// Path of 'c'
	exp = batch16pathC
	act = tree.Path(c)
	equal(t, &exp, &act, c)

	// Path of 'd'
	exp = batch16pathD
	act = tree.Path(d)
	equal(t, &exp, &act, d)

	// Path of 'e'
	exp = batch16pathE
	act = tree.Path(e)
	equal(t, &exp, &act, e)

	// Path of 'f'
	exp = batch16pathF
	act = tree.Path(f)
	equal(t, &exp, &act, f)

	// Path of 'g'
	exp = batch16pathG
	act = tree.Path(g)
	equal(t, &exp, &act, g)

	// Path of 'h'
	exp = batch16pathH
	act = tree.Path(h)
	equal(t, &exp, &act, h)

	// Path of 'i'
	exp = batch16pathI
	act = tree.Path(i)
	equal(t, &exp, &act, i)

	// Path of 'j'
	exp = batch16pathJ
	act = tree.Path(j)
	equal(t, &exp, &act, j)

	// Path of 'k'
	exp = batch16pathK
	act = tree.Path(k)
	equal(t, &exp, &act, k)

	// Path of 'l'
	exp = batch16pathL
	act = tree.Path(l)
	equal(t, &exp, &act, l)

	// Path of 'm'
	exp = batch16pathM
	act = tree.Path(m)
	equal(t, &exp, &act, m)

	// Path of 'n'
	exp = batch16pathN
	act = tree.Path(n)
	equal(t, &exp, &act, n)

	// Path of 'o'
	exp = batch16pathO
	act = tree.Path(o)
	equal(t, &exp, &act, o)

	// Path of 'p'
	exp = batch16pathP
	act = tree.Path(p)
	equal(t, &exp, &act, p)
}

func TestValidate(t *testing.T) {
	root := abcdefghijklmnop
	// Path of 'a'
	validate(t, &batch16pathA, root, a)
	// Path of 'b'
	validate(t, &batch16pathB, root, b)
	// Path of 'c'
	validate(t, &batch16pathC, root, c)
	// Path of 'd'
	validate(t, &batch16pathD, root, d)
	// Path of 'e'
	validate(t, &batch16pathE, root, e)
	// Path of 'f'
	validate(t, &batch16pathF, root, f)
	// Path of 'g'
	validate(t, &batch16pathG, root, g)
	// Path of 'h'
	validate(t, &batch16pathH, root, h)
	// Path of 'i'
	validate(t, &batch16pathI, root, i)
	// Path of 'j'
	validate(t, &batch16pathJ, root, j)
	// Path of 'k'
	validate(t, &batch16pathK, root, k)
	// Path of 'l'
	validate(t, &batch16pathL, root, l)
	// Path of 'm'
	validate(t, &batch16pathM, root, m)
	// Path of 'n'
	validate(t, &batch16pathN, root, n)
	// Path of 'o'
	validate(t, &batch16pathO, root, o)
	// Path of 'p'
	validate(t, &batch16pathP, root, p)
}

func TestValidate_InvalidPath(t *testing.T) {
	// Invalid path of 'a'
	path := []*Path{
		{L: b}, // this is invalid, should be R
		{R: cd},
		{R: efgh},
		{R: ijklmnop},
	}
	invalidate(t, path, batch16root, a)

	// Invalid path of 'f'
	path = []*Path{
		{L: e},
		{L: gh}, // this is invalid, should be R
		{L: abcd},
		{R: ijklmnop},
	}
	invalidate(t, path, batch16root, f)

	// Invalid path of 'k'
	path = []*Path{
		{R: l},
		{L: ij},
		{L: mnop}, // this is invalid, should be R
		{L: abcdefgh},
	}
	invalidate(t, path, batch16root, k)

	// Invalid path of 'p'
	path = []*Path{
		{L: o},
		{L: mn},
		{L: ijkl},
		{R: abcdefgh}, // this is invalid, should be L
	}
	invalidate(t, path, batch16root, p)
}

func TestTree_Root(t *testing.T) {
	builder := NewBuilder(SHA256)
	builder.AddBatch(&batch16)
	tree := builder.Build()
	if tree.Root() != batch16root {
		t.Fail()
	}
}

func TestTree_Leaves(t *testing.T) {
	builder := NewBuilder(SHA256)
	builder.AddBatch(&batch16)
	tree := builder.Build()
	exp := []string{
		a, b, c, d, e, f, g, h, i, j, k, l, m, n, o, p,
	}
	if !reflect.DeepEqual(exp, tree.Leaves()) {
		t.Fail()
	}
}

func TestTree_Verify(t *testing.T) {
	builder := NewBuilder(SHA256)
	builder.AddBatch(&batch16)
	tree := builder.Build()
	if !tree.Verify(abcdefghijklmnop) {
		t.Fail()
	}
}

func TestTree_Level(t *testing.T) {
	builder := NewBuilder(SHA256)
	builder.AddBatch(&batch16)
	tree := builder.Build()

	// level 0 (root)
	level := []string{abcdefghijklmnop}
	if !reflect.DeepEqual(level, tree.Level(0)) {
		t.Fatal("no match for level 0")
	}

	// level 1
	level = []string{abcdefgh, ijklmnop}
	if !reflect.DeepEqual(level, tree.Level(1)) {
		t.Fatal("no match for level 1")
	}

	// level 2
	level = []string{abcd, efgh, ijkl, mnop}
	if !reflect.DeepEqual(level, tree.Level(2)) {
		t.Fatal("no match for level 2")
	}

	// level 3
	level = []string{ab, cd, ef, gh, ij, kl, mn, op}
	if !reflect.DeepEqual(level, tree.Level(3)) {
		t.Fatal("no match for level 3")
	}

	// level 4 (leaves)
	level = []string{a, b, c, d, e, f, g, h, i, j, k, l, m, n, o, p}
	if !reflect.DeepEqual(level, tree.Level(4)) {
		t.Fatal("no match for level 4")
	}
}

func TestTree_CountDepth(t *testing.T) {
	builder := NewBuilder(SHA256)
	builder.AddBatch(&batch16)
	tree := builder.Build()
	if tree.CountDepth() != 4 {
		t.Fail()
	}
}

func TestTree_CountLevels(t *testing.T) {
	builder := NewBuilder(SHA256)
	builder.AddBatch(&batch16)
	tree := builder.Build()
	if tree.CountLevels() != 5 {
		t.Fail()
	}
}

func TestTree_CountLeaves(t *testing.T) {
	builder := NewBuilder(SHA256)
	builder.AddBatch(&batch16)
	tree := builder.Build()
	if tree.CountLeaves() != 16 {
		t.Fail()
	}
}

func TestTree_CountNodes(t *testing.T) {
	builder := NewBuilder(SHA256)
	builder.AddBatch(&batch16)
	tree := builder.Build()
	if tree.CountNodes() != 15 {
		t.Fail()
	}
}
