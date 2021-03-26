package merkle

import (
	"encoding/json"

	"github.com/SouthbankSoftware/provendb-sdk-go/pkg/anchor"
)

type Proof interface {
	// Metadata provides information about the anchoring.
	Metadata() map[string]interface{}
	// Data is the actual receipt.
	Data() interface{}
}

type proof struct {
	Meta map[string]interface{} `json:"metadata"`
	Raw  interface{}            `json:"data"`
}

func (p *proof) Metadata() map[string]interface{} {
	return p.Meta
}

func (p *proof) Data() interface{} {
	return p.Data
}

func FromAnchorProof(p *anchor.Proof) Proof {
	x := &proof{}
	if p.Batch != nil {
		if p.Batch.Data != "" {
			// This data is JSON string. Unmarshall to json.
			var meta map[string]interface{}
			if err := json.Unmarshal([]byte(p.Batch.Data), &meta); err != nil {
				panic("failed to unmarshal anchor proof batch data")
			}
			x.Meta = meta
		}
	}
	if p.Data != "" {
		raw, err := anchor.DecodeProof(p.Data)
		if err != nil {
			panic("invalid data")
		}
		x.Raw = raw
	}
	return x
}
