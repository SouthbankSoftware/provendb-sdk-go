package anchor

import (
	"bytes"
	"compress/zlib"
	"encoding/base64"
	"encoding/json"
	"fmt"

	"github.com/vmihailenco/msgpack"
)

/**
 * AnchorProof is a represention of a Proof object with tasks such as decoding already
 * performed and enums represented as strings for readability.
 */
type AnchorProof struct {
	Id         string
	AnchorType string
	BatchId    string
	Status     string
	Format     string
	Hash       string
	Metadata   interface{}
	Data       map[string]interface{}
}

func (a *AnchorProof) FromProof(proof *Proof) error {
	if proof.GetBatch() != nil {
		if proof.GetBatch().GetData() != "" {
			var data interface{}
			if e := json.Unmarshal([]byte(proof.GetBatch().GetData()), &data); e != nil {
				fmt.Print(e.Error())
				return e
			}
			// var meta map[string]interface{}
			// if e := json.Unmarshal([]byte(data), &meta); e != nil {
			// 	return e
			// }
			a.Metadata = data
		}
	}
	if proof.GetData() != "" {
		data, err := DecodeProof(proof.GetData())
		if err != nil {
			return err
		}
		a.Data = data
	}
	a.Id = generateProofId(proof.GetHash(), proof.GetBatchId())
	a.AnchorType = proof.GetAnchorType().String()
	a.Format = proof.GetFormat().String()
	a.BatchId = proof.GetBatchId()
	a.Hash = proof.GetHash()
	a.Status = proof.GetBatchStatus().String()
	return nil
}

func DecodeProof(data string) (map[string]interface{}, error) {
	b, err := base64.StdEncoding.DecodeString(data)
	if err != nil {
		return nil, err
	}

	z, err := zlib.NewReader(bytes.NewReader(b))
	if err != nil {
		return nil, err
	}
	defer z.Close()
	var m map[string]interface{}
	if err := msgpack.NewDecoder(z).UseJSONTag(true).Decode(&m); err != nil {
		return nil, err
	}
	return m, nil
}
