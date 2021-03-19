package anchor

import "github.com/SouthbankSoftware/provendb-sdk-go/genproto/anchor"

const (
	BitcoinMainnet  AnchorType = 1
	BitcoinTestnet  AnchorType = 2
	EthereumMainnet AnchorType = 3
	EthereumTestnet AnchorType = 4
	Hyperledger     AnchorType = 5
)

const (
	ChainpointFormat ProofFormat = 1
)

// Type represents a single anchor type.
type AnchorType int

type ProofFormat int

// ProtoType returns the proto type of this anchor type.
func (t AnchorType) ProtoType() anchor.Anchor_Type {
	switch t {
	case BitcoinMainnet:
		return anchor.Anchor_BTC_MAINNET
	case BitcoinTestnet:
		return anchor.Anchor_BTC
	case EthereumMainnet:
		return anchor.Anchor_ETH_MAINNET
	case EthereumTestnet:
		return anchor.Anchor_ETH
	case Hyperledger:
		return anchor.Anchor_HYPERLEDGER
	default:
		panic("type not defined")
	}
}

// ProtoType returns the proto type of this format.
func (f ProofFormat) ProtoType() anchor.Proof_Format {
	switch f {
	case ChainpointFormat:
		return anchor.Proof_CHP_PATH
	default:
		panic("format not defined")
	}
}

// Proof struct.
type Proof struct {
	// The anchor type used to anchor the proof.
	AnchorType AnchorType `json:"anchor_type"`
	// Metadata is any proof related metadata including transaction information, prices, etc.
	AnchorData interface{} `json:"metadata"`
	// BatchID
	BatchID string `json:"batch_id"`
	// The format the proof is in.
	Format ProofFormat `json:"format"`
	// The submitted hash for this proof.
	Hash string
	// Status is the proof status.
	Status string `json:"status"`
	// Proof is the actual proof of the submitted hash.
	Proof interface{} `json:"proof"`
}

// BitcoinMetadata contains metadata related to the anchoring on the Bitcoin blockchains.
type BitcoinMetadata struct {
	TransactionID string `json:"transaction_id"`
}

// EthereumMetadata contains metadata related to anchoring on the Ethereum blockchains.
type EthereumMetadata struct {
	GasPrice string `json:"gas_price"`
}
