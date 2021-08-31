package merkle

import (
	_ "crypto/sha256"
	"reflect"
	"testing"
)

const (
	// The leaves (a, b, c, ... p) are hashed using a utf8 encoded string, and all hashes following are hashed
	// using hex encoded strings. There should be enough data here to create your test cases.
	a     = "ca978112ca1bbdcafac231b39a23dc4da786eff8147c4e72b9807785afee48bb"
	aLeaf = "a:" + a
	b     = "3e23e8160039594a33894f6564e1b1348bbd7a0088d42c4acb73eeaed59c009d"
	bLeaf = "b:" + b
	c     = "2e7d2c03a9507ae265ecf5b5356885a53393a2029d241394997265a1a25aefc6"
	cLeaf = "c:" + c
	d     = "18ac3e7343f016890c510e93f935261169d9e3f565436429830faf0934f4f8e4"
	dLeaf = "d:" + d
	e     = "3f79bb7b435b05321651daefd374cdc681dc06faa65e374e38337b88ca046dea"
	eLeaf = "e:" + e
	f     = "252f10c83610ebca1a059c0bae8255eba2f95be4d1d7bcfa89d7248a82d9f111"
	fLeaf = "f:" + f
	g     = "cd0aa9856147b6c5b4ff2b7dfee5da20aa38253099ef1b4a64aced233c9afe29"
	gLeaf = "g:" + g
	h     = "aaa9402664f1a41f40ebbc52c9993eb66aeb366602958fdfaa283b71e64db123"
	hLeaf = "h:" + h
	i     = "de7d1b721a1e0632b7cf04edf5032c8ecffa9f9a08492152b926f1a5a7e765d7"
	iLeaf = "i:" + i
	j     = "189f40034be7a199f1fa9891668ee3ab6049f82d38c68be70f596eab2e1857b7"
	jLeaf = "j:" + j
	k     = "8254c329a92850f6d539dd376f4816ee2764517da5e0235514af433164480d7a"
	kLeaf = "k:" + k
	l     = "acac86c0e609ca906f632b0e2dacccb2b77d22b0621f20ebece1a4835b93f6f0"
	lLeaf = "l:" + l
	m     = "62c66a7a5dd70c3146618063c344e531e6d4b59e379808443ce962b3abd63c5a"
	mLeaf = "m:" + m
	n     = "1b16b1df538ba12dc3f97edbb85caa7050d46c148134290feba80f8236c83db9"
	nLeaf = "n:" + n
	o     = "65c74c15a686187bb6bbf9958f494fc6b80068034a659a9ad44991b08c58f2d2"
	oLeaf = "o:" + o
	p     = "148de9c5a7a44d19e56cd9ae1a554bf67847afb0c58f6e12fa29ac7ddfca9940"
	pLeaf = "p:" + p

	ab = "e5a01fee14e0ed5c48714f22180f25ad8365b53f9779f79dc4a3d7e93963f94a"
	cd = "bffe0b34dba16bc6fac17c08bac55d676cded5a4ade41fe2c9924a5dde8f3e5b"
	ef = "04fa33f8b4bd3db545fa04cdd51b462509f611797c7bfe5c944ee2bb3b2ed908"
	gh = "140257c1540113794d2ae3394879e586ca5caebca19663ff87417892cf36fd23"
	ij = "cf1970792b0aa5816da207cd936e21f594f4e9c68cb01ca97d843047e3107958"
	kl = "6c192582bc0f32bf1ba5833b200db8795b8fbe49228f9a73c09687a777dfca61"
	mn = "94ffc897da3f6a1098eb7b573721291eb9c58154e3fbd10e525c27baa0108bae"
	op = "91756679180e8130ee47a9cb4713261e4e3189d1aef40087dce9c393e38e84fd"

	mno = "7b0fbd42a983b6214192e1531e68881716c87c4672cd17799779b635e685a273"

	abcd = "14ede5e8e97ad9372327728f5099b95604a39593cac3bd38a343ad76205213e7"
	efgh = "8e2c530a100033894555cde1c7d4e36f7c6e553ee3914022ec7a13e1196baed2"
	ijkl = "9ed3e37faea35ec0ddf7bd4e7ea9e8e47ce83dfa84e13c8874646d83079c72aa"
	mnop = "bab7598e438c316f64a14876fb50be7177adf9eaaf257eabd60b84662dea09f8"

	abcdef = "1f7379539707bcaea00564168d1d4d626b09b73f8a2a365234c62d763f854da2"

	ijklmno = "6eeb6ef56df316c559cc627dd31358ce494fed3db575668ad93e4e4102d5025e"

	abcdefgh = "bd7c8a900be9b67ba7df5c78a652a8474aedd78adb5083e80e49d9479138a23f"
	ijklmnop = "5a2419accdde223b023d7bd53f4c58758207598b227c31e92c4120593c9e4ca3"

	abcdefghijklmno = "5486677cd239f0bde3a0bf517fef8de3cc04e75731be77642b30b6671833c76d"

	abcdefghijklmnop = "2eb4698fb52b0cd41d51e50f1878c2c23fdba3be61c73da456a8c40aea13003c"

	message       = "this is a message designed to be used as input to the builder's Writer() method"
	messageHashed = "2eedeb6c4d47f67831ffb0df5726f37d4137351bb39a88275cd5c1b7e0f024a4"
	messageLeaf   = "message:" + messageHashed

	abcdefghijklmnop_messageHashed = "d5d55c1dba8af00399a878abc75c21d328caa1815cb7fbaa5ad106e6eb9c0fea"
)

// func BenchmarkBuilder_Build(b *testing.B) {
// 	builder := NewBuilder(SHA256)
// 	for i := 0; i < b.N; i++ {
// 		builder.Add(string(i), []byte(strconv.Itoa(i)))
// 	}
// 	// Time the construction of the tree.
// 	b.StartTimer()
// 	builder.Build()
// 	b.StopTimer()
// }

// func TestBuilder(t *testing.T) {
// 	builder := NewBuilder(SHA256)
// 	// Add singularly
// 	builder.Add("a", []byte("a"))
// 	builder.Add("b", []byte("b"))
// 	builder.Add("c", []byte("c"))
// 	builder.Add("d", []byte("d"))
// 	builder.Add("e", []byte("e"))
// 	builder.Add("f", []byte("f"))
// 	builder.Add("g", []byte("g"))
// 	builder.Add("h", []byte("h"))
// 	// Add a batch
// 	builder.AddBatch(map[string][]byte{
// 		"i": []byte("i"),
// 		"j": []byte("j"),
// 		"k": []byte("k"),
// 		"l": []byte("l"),
// 		"m": []byte("m"),
// 		"n": []byte("n"),
// 		"o": []byte("o"),
// 		"p": []byte("p"),
// 	})
// 	// Write the last leaf, this should be promoted all the way to the top
// 	writer := builder.Writer("message")
// 	chars := []byte(message)
// 	// Write one byte at a time
// 	for i := 0; i < len(chars); i++ {
// 		n, e := writer.Write([]byte{chars[i]})
// 		if e != nil || n != 1 {
// 			t.Errorf("failed to write byte: %s", e.Error())
// 		}
// 	}
// 	writer.Close()

// 	tree := builder.Build()
// 	exp := [][]string{
// 		{aLeaf, bLeaf, cLeaf, dLeaf, eLeaf, fLeaf, gLeaf, hLeaf, iLeaf, jLeaf, kLeaf, lLeaf, mLeaf, nLeaf, oLeaf, pLeaf, messageLeaf},
// 		{ab, cd, ef, gh, ij, kl, mn, op, messageHashed},
// 		{abcd, efgh, ijkl, mnop, messageHashed},
// 		{abcdefgh, ijklmnop, messageHashed},
// 		{abcdefghijklmnop, messageHashed},
// 		{abcdefghijklmnop_messageHashed},
// 	}
// 	if !reflect.DeepEqual(exp, tree.GetLevels()) {
// 		t.Fail()
// 	}

// }
// func TestBuilder_Writer(t *testing.T) {
// 	builder := NewBuilder(SHA256)
// 	writer := builder.Writer("message")

// 	// Executing a single Add() in a separate go routine. This should be added after our writer call.
// 	chars := []byte(message)
// 	// Write one byte at a time
// 	for i := 0; i < len(chars); i++ {
// 		n, e := writer.Write([]byte{chars[i]})
// 		if e != nil || n != 1 {
// 			t.Errorf("failed to write byte: %s", e.Error())
// 		}
// 	}
// 	writer.Close()

// 	tree := builder.Build()
// 	exp := [][]string{{messageLeaf}}
// 	if !reflect.DeepEqual(exp, tree.GetLevels()) {
// 		t.Fail()
// 	}
// }

// func TestBuilder_Add(t *testing.T) {
// 	builder := NewBuilder(SHA256)
// 	// Add singularly
// 	builder.Add("a", []byte("a"))
// 	builder.Add("b", []byte("b"))
// 	builder.Add("c", []byte("c"))
// 	builder.Add("d", []byte("d"))
// 	builder.Add("e", []byte("e"))
// 	builder.Add("f", []byte("f"))
// 	builder.Add("g", []byte("g"))
// 	builder.Add("h", []byte("h"))
// 	builder.Add("i", []byte("i"))
// 	builder.Add("j", []byte("j"))
// 	builder.Add("k", []byte("k"))
// 	builder.Add("l", []byte("l"))
// 	builder.Add("m", []byte("m"))
// 	builder.Add("n", []byte("n"))
// 	builder.Add("o", []byte("o"))
// 	builder.Add("p", []byte("p"))

// 	tree := builder.Build()

// 	exp := [][]string{
// 		{aLeaf, bLeaf, cLeaf, dLeaf, eLeaf, fLeaf, gLeaf, hLeaf, iLeaf, jLeaf, kLeaf, lLeaf, mLeaf, nLeaf, oLeaf, pLeaf},
// 		{ab, cd, ef, gh, ij, kl, mn, op},
// 		{abcd, efgh, ijkl, mnop},
// 		{abcdefgh, ijklmnop},
// 		{abcdefghijklmnop},
// 	}

// 	if !reflect.DeepEqual(exp, tree.GetLevels()) {
// 		t.Fail()
// 	}
// }

// func TestBuilder_AddBatch(t *testing.T) {
// 	builder := NewBuilder(SHA256)
// 	builder.AddBatch(map[string][]byte{
// 		"a": []byte("a"),
// 		"b": []byte("b"),
// 		"c": []byte("c"),
// 		"d": []byte("d"),
// 		"e": []byte("e"),
// 		"f": []byte("f"),
// 		"g": []byte("g"),
// 		"h": []byte("h"),
// 	})
// 	builder.AddBatch(map[string][]byte{
// 		"i": []byte("i"),
// 		"j": []byte("j"),
// 		"k": []byte("k"),
// 		"l": []byte("l"),
// 		"m": []byte("m"),
// 		"n": []byte("n"),
// 		"o": []byte("o"),
// 		"p": []byte("p"),
// 	})

// 	tree := builder.Build()
// 	if !tree.Verify(abcdefghijklmnop) {
// 		t.Fail()
// 	}
// }

func TestBuild_Many(t *testing.T) {
	leaves := []string{aLeaf, bLeaf, cLeaf, dLeaf, eLeaf, fLeaf, gLeaf, hLeaf, iLeaf, jLeaf, kLeaf, lLeaf, mLeaf, nLeaf, oLeaf, pLeaf}
	exp := [][]string{
		leaves,
		{ab, cd, ef, gh, ij, kl, mn, op},
		{abcd, efgh, ijkl, mnop},
		{abcdefgh, ijklmnop},
		{abcdefghijklmnop},
	}
	act := build(leaves, SHA256)
	if !reflect.DeepEqual(exp, act) {
		t.Fail()
	}
}

func TestBuild_ManyPromoted(t *testing.T) {
	leaves := []string{aLeaf, bLeaf, cLeaf, dLeaf, eLeaf, fLeaf, gLeaf, hLeaf, iLeaf, jLeaf, kLeaf, lLeaf, mLeaf, nLeaf, oLeaf}
	exp := [][]string{
		leaves,
		{ab, cd, ef, gh, ij, kl, mn, o},
		{abcd, efgh, ijkl, mno},
		{abcdefgh, ijklmno},
		{abcdefghijklmno},
	}
	act := build(leaves, SHA256)
	if !reflect.DeepEqual(exp, act) {
		t.Fail()
	}
}

// func TestBuild_Three(t *testing.T) {
// 	leaves := []string{ab, cd, ef}
// 	exp := [][]string{
// 		leaves,
// 		{abcd, ef},
// 		{abcdef},
// 	}
// 	act := build(leaves, SHA256)
// 	if !reflect.DeepEqual(exp, act) {
// 		t.Fail()
// 	}
// }

func TestBuild_Two(t *testing.T) {
	leaves := []string{aLeaf, bLeaf}
	exp := [][]string{
		leaves,
		{ab},
	}
	act := build(leaves, SHA256)
	if !reflect.DeepEqual(exp, act) {
		t.Fail()
	}
}

// func TestBuild_One(t *testing.T) {
// 	leaves := []string{abcdefghijklmnop}
// 	exp := [][]string{leaves}
// 	act := build(leaves, SHA256)
// 	if !reflect.DeepEqual(exp, act) {
// 		t.Fail()
// 	}
// }

// // Tests the build of a level with many elements up to an even value.
// func TestBuildLevel_Many(t *testing.T) {
// 	level := []string{a, b, c, d, e, f, g, h, i, j, k, l, m, n, o, p}
// 	exp := []string{ab, cd, ef, gh, ij, kl, mn, op}
// 	act := buildLevel(level, SHA256)
// 	if !reflect.DeepEqual(exp, act) {
// 		t.Fail()
// 	}
// }

// Tests the build of a level with many elements up to an odd value (to ensure node promotion)
// func TestBuildLevel_ManyPromoted(t *testing.T) {
// 	level := []string{a, b, c, d, e, f, g, h, i}
// 	exp := []string{
// 		ab,
// 		cd,
// 		ef,
// 		gh,
// 		i,
// 	}
// 	act := BuildLevel(level, SHA256)
// 	if !reflect.DeepEqual(exp, act) {
// 		t.Fail()
// 	}
// }

// func TestBuildLevel_Promoted(t *testing.T) {
// 	level := []string{a, b, c}
// 	exp := []string{ab, c}
// 	act := BuildLevel(level, SHA256)
// 	if !reflect.DeepEqual(exp, act) {
// 		t.Fail()
// 	}
// }

// Tests the build of a level with two elements.
// func TestBuildLevel_Two(t *testing.T) {
// 	level := []string{a, b}
// 	exp := []string{ab}
// 	act := BuildLevel(level, SHA256)
// 	if !reflect.DeepEqual(exp, act) {
// 		t.Fail()
// 	}

// }

// // Tests the build of a level with a single element.
// func TestBuildLevel_One(t *testing.T) {
// 	level := []string{abcdefghijklmnop}
// 	exp := level
// 	act := BuildLevel(level, SHA256)
// 	if !reflect.DeepEqual(exp, act) {
// 		t.Fail()
// 	}
// }
