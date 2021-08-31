package merkle

import "crypto"

const (
	SHA224   Hash = "sha-224"
	SHA256   Hash = "sha-256"
	SHA384   Hash = "sha-384"
	SHA512   Hash = "sha-512"
	SHA3_224 Hash = "sha3-224"
	SHA3_256 Hash = "sha3-256"
	SHA3_384 Hash = "sha3-384"
	SHA3_512 Hash = "sha3-512"
)

// Hash represents the hash functions.
type Hash string

// Hash returns the crypto.Hash of the given hash.
func (h Hash) Hash() crypto.Hash {
	switch h {
	case SHA224:
		return crypto.SHA224
	case SHA256:
		return crypto.SHA256
	case SHA384:
		return crypto.SHA384
	case SHA512:
		return crypto.SHA512
	case SHA3_224:
		return crypto.SHA3_224
	case SHA3_256:
		return crypto.SHA3_256
	case SHA3_384:
		return crypto.SHA3_384
	case SHA3_512:
		return crypto.SHA3_512
	default:
		panic("unknown hash")
	}
}
