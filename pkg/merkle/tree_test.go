package merkle

import (
	_ "crypto/sha256"
	"reflect"
	"testing"
)

// Variables for trees with 16 leaves.
var (
	batch16 = []*struct {
		Key   string
		Value []byte
	}{
		{Key: "a", Value: []byte("a")},
		{Key: "b", Value: []byte("b")},
		{Key: "c", Value: []byte("c")},
		{Key: "d", Value: []byte("d")},
		{Key: "e", Value: []byte("e")},
		{Key: "f", Value: []byte("f")},
		{Key: "g", Value: []byte("g")},
		{Key: "h", Value: []byte("h")},
		{Key: "i", Value: []byte("i")},
		{Key: "j", Value: []byte("j")},
		{Key: "k", Value: []byte("k")},
		{Key: "l", Value: []byte("l")},
		{Key: "m", Value: []byte("m")},
		{Key: "n", Value: []byte("n")},
		{Key: "o", Value: []byte("o")},
		{Key: "p", Value: []byte("p")},
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

// // Validates the path by ensuring the final calculated hash DOES match the path.
// func validate(t *testing.T, path *[]*Path, root string, leaf string) {
// 	ok, err := ValidatePath(*path, leaf, SHA256, root)
// 	if err != nil {
// 		t.Fatal(err)
// 	}
// 	if !ok {
// 		t.Fatalf("root mismatch for '%s'", leaf)
// 	}
// }

// // Invalidates the path by ensuring the final calculated hash does NOT match.
// func invalidate(t *testing.T, path []*Path, root string, leaf string) {
// 	ok, err := ValidatePath(path, leaf, SHA256, root)
// 	if err != nil {
// 		t.Fatal(err)
// 	}
// 	if ok {
// 		t.Fatalf("root should have been different for '%s'", a)
// 	}
// }

// func TestTree_Export(t *testing.T) {
// 	builder := NewBuilder(SHA256)
// 	builder.AddBatch(&batch16)
// 	tree := builder.Build()
// 	tree.AddProof(anchor.Proof{
// 		Hash: "",
// 		Data: "eJykkbGS0zAQQH+G1rG0K9lWqszwC1Q0nt3VGmsIlsfWBa4EGtp8w4XhYCiBkv/I3zC5wHU00D7pvZ3Z/XC/kzwVfVN+jqXM67auX2OKm7y8qGWkNM05TaU+4Knczvrl6SM6jbSO5x3IAC4y+YENoVFotEMGEs9BIahl21j1YgdqHXht0TjEBgJ2HDpwXy+ZPsV+ylHPTyIZQPFSSRd8Za2GihCGCghIIbKIuh8PynrDr1IpejV7Kt/AgK0MVuCfAW6d3Vp8/piXvPxj/mL+LX/PC00y6np893FPrPvvc+Rey9hfcF766/tdntfj209Xth7fP+zyTmj/+fe3FM+7VnCAltk2rlWN1nhpPEb0A6m6BkNk10Egb7yG0EqnTL7V0HaN+GCa082S1uP55Z8zXtObqIfNvOQ8EO91k3J90CUNt7WWsf7fkb8CAAD//5Arw98=",
// 		Batch: &anchor.Batch{
// 			Data: "{\"txnId\":\"0565eb2dd0c8c7e85864e937b9997b6d080f69f10cf71ad4b3113f9a707c6d7a\",\"txnUri\":\"https://rinkeby.etherscan.io/tx/0x0565eb2dd0c8c7e85864e937b9997b6d080f69f10cf71ad4b3113f9a707c6d7a\",\"blockTime\":1616716627,\"blockNumber\":8298881,\"endpoint\":\"https://rinkeby.infura.io/v3/ba25a62205f24e5bb74d4f9738910a83\",\"gasUsed\":21512,\"gasPrice\":1000000000}",
// 		},
// 	})
// 	if err := tree.Export("testing/merkle.json"); err != nil {
// 		t.Fail()
// 	}
// }

func TestTree_Path(t *testing.T) {
	builder := NewBuilder(SHA256)
	builder.AddBatch(batch16)
	tree := builder.Build()

	// Path of 'a'
	exp := batch16pathA
	act := tree.GetPath("a")
	equal(t, &exp, &act, a)

	// Path of 'b'
	exp = batch16pathB
	act = tree.GetPath("b")
	equal(t, &exp, &act, b)

	// Path of 'c'
	exp = batch16pathC
	act = tree.GetPath("c")
	equal(t, &exp, &act, c)

	// Path of 'd'
	exp = batch16pathD
	act = tree.GetPath("d")
	equal(t, &exp, &act, d)

	// Path of 'e'
	exp = batch16pathE
	act = tree.GetPath("e")
	equal(t, &exp, &act, e)

	// Path of 'f'
	exp = batch16pathF
	act = tree.GetPath("f")
	equal(t, &exp, &act, f)

	// Path of 'g'
	exp = batch16pathG
	act = tree.GetPath("g")
	equal(t, &exp, &act, g)

	// Path of 'h'
	exp = batch16pathH
	act = tree.GetPath("h")
	equal(t, &exp, &act, h)

	// Path of 'i'
	exp = batch16pathI
	act = tree.GetPath("i")
	equal(t, &exp, &act, i)

	// Path of 'j'
	exp = batch16pathJ
	act = tree.GetPath("j")
	equal(t, &exp, &act, j)

	// Path of 'k'
	exp = batch16pathK
	act = tree.GetPath("k")
	equal(t, &exp, &act, k)

	// Path of 'l'
	exp = batch16pathL
	act = tree.GetPath("l")
	equal(t, &exp, &act, l)

	// Path of 'm'
	exp = batch16pathM
	act = tree.GetPath("m")
	equal(t, &exp, &act, m)

	// Path of 'n'
	exp = batch16pathN
	act = tree.GetPath("n")
	equal(t, &exp, &act, n)

	// Path of 'o'
	exp = batch16pathO
	act = tree.GetPath("o")
	equal(t, &exp, &act, o)

	// Path of 'p'
	exp = batch16pathP
	act = tree.GetPath("p")
	equal(t, &exp, &act, p)
}

// func TestValidate(t *testing.T) {
// 	root := abcdefghijklmnop
// 	// Path of 'a'
// 	validate(t, &batch16pathA, root, a)
// 	// Path of 'b'
// 	validate(t, &batch16pathB, root, b)
// 	// Path of 'c'
// 	validate(t, &batch16pathC, root, c)
// 	// Path of 'd'
// 	validate(t, &batch16pathD, root, d)
// 	// Path of 'e'
// 	validate(t, &batch16pathE, root, e)
// 	// Path of 'f'
// 	validate(t, &batch16pathF, root, f)
// 	// Path of 'g'
// 	validate(t, &batch16pathG, root, g)
// 	// Path of 'h'
// 	validate(t, &batch16pathH, root, h)
// 	// Path of 'i'
// 	validate(t, &batch16pathI, root, i)
// 	// Path of 'j'
// 	validate(t, &batch16pathJ, root, j)
// 	// Path of 'k'
// 	validate(t, &batch16pathK, root, k)
// 	// Path of 'l'
// 	validate(t, &batch16pathL, root, l)
// 	// Path of 'm'
// 	validate(t, &batch16pathM, root, m)
// 	// Path of 'n'
// 	validate(t, &batch16pathN, root, n)
// 	// Path of 'o'
// 	validate(t, &batch16pathO, root, o)
// 	// Path of 'p'
// 	validate(t, &batch16pathP, root, p)
// }

// func TestValidate_InvalidPath(t *testing.T) {
// 	// Invalid path of 'a'
// 	path := []*Path{
// 		{L: b}, // this is invalid, should be R
// 		{R: cd},
// 		{R: efgh},
// 		{R: ijklmnop},
// 	}
// 	invalidate(t, path, batch16root, a)

// 	// Invalid path of 'f'
// 	path = []*Path{
// 		{L: e},
// 		{L: gh}, // this is invalid, should be R
// 		{L: abcd},
// 		{R: ijklmnop},
// 	}
// 	invalidate(t, path, batch16root, f)

// 	// Invalid path of 'k'
// 	path = []*Path{
// 		{R: l},
// 		{L: ij},
// 		{L: mnop}, // this is invalid, should be R
// 		{L: abcdefgh},
// 	}
// 	invalidate(t, path, batch16root, k)

// 	// Invalid path of 'p'
// 	path = []*Path{
// 		{L: o},
// 		{L: mn},
// 		{L: ijkl},
// 		{R: abcdefgh}, // this is invalid, should be L
// 	}
// 	invalidate(t, path, batch16root, p)
// }

func TestTree_Root(t *testing.T) {
	builder := NewBuilder(SHA256)
	builder.AddBatch(batch16)
	tree := builder.Build()
	if tree.GetRoot() != batch16root {
		t.Fail()
	}
}

func TestTree_Leaves(t *testing.T) {
	builder := NewBuilder(SHA256)
	builder.AddBatch(batch16)
	tree := builder.Build()
	exp := []string{
		a, b, c, d, e, f, g, h, i, j, k, l, m, n, o, p,
	}
	if !reflect.DeepEqual(exp, tree.GetLeaves()) {
		t.Fail()
	}
}

func TestTree_Verify(t *testing.T) {
	builder := NewBuilder(SHA256)
	builder.AddBatch(batch16)
	tree := builder.Build()
	if !tree.Verify(abcdefghijklmnop) {
		t.Fail()
	}
}

func TestTree_Level(t *testing.T) {
	builder := NewBuilder(SHA256)
	builder.AddBatch(batch16)
	tree := builder.Build()

	// level 0 (root)
	level := []string{abcdefghijklmnop}
	if !reflect.DeepEqual(level, tree.GetLevel(0)) {
		t.Fatal("no match for level 0")
	}

	// level 1
	level = []string{abcdefgh, ijklmnop}
	if !reflect.DeepEqual(level, tree.GetLevel(1)) {
		t.Fatal("no match for level 1")
	}

	// level 2
	level = []string{abcd, efgh, ijkl, mnop}
	if !reflect.DeepEqual(level, tree.GetLevel(2)) {
		t.Fatal("no match for level 2")
	}

	// level 3
	level = []string{ab, cd, ef, gh, ij, kl, mn, op}
	if !reflect.DeepEqual(level, tree.GetLevel(3)) {
		t.Fatal("no match for level 3")
	}

	// level 4 (leaves)
	level = []string{a, b, c, d, e, f, g, h, i, j, k, l, m, n, o, p}
	if !reflect.DeepEqual(level, tree.GetLevel(4)) {
		t.Fatal("no match for level 4")
	}
}

func TestTree_CountDepth(t *testing.T) {
	builder := NewBuilder(SHA256)
	builder.AddBatch(batch16)
	tree := builder.Build()
	if tree.NDepth() != 4 {
		t.Fail()
	}
}

func TestTree_CountLevels(t *testing.T) {
	builder := NewBuilder(SHA256)
	builder.AddBatch(batch16)
	tree := builder.Build()
	if tree.NLevels() != 5 {
		t.Fail()
	}
}

func TestTree_CountLeaves(t *testing.T) {
	builder := NewBuilder(SHA256)
	builder.AddBatch(batch16)
	tree := builder.Build()
	if tree.NLeaves() != 16 {
		t.Fail()
	}
}

func TestTree_CountNodes(t *testing.T) {
	builder := NewBuilder(SHA256)
	builder.AddBatch(batch16)
	tree := builder.Build()
	if tree.NNodes() != 15 {
		t.Fail()
	}
}
