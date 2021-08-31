package anchor

import (
	"errors"
	"fmt"
)

/**
 * Generates a proof ID using the proof.hash and proof.batchId
 * @param hash the hash
 * @param batchId the batch ID
 * @returns the proof ID
 */
func generateProofId(hash string, batchId string) string {
	return fmt.Sprintf("%s:%s", hash, batchId)
}

// Retrieves the anchor type from either string or proto.Anchor_Type
func getAnchorType(anchorType interface{}) (Anchor_Type, error) {
	switch anchorType.(type) {
	case string:
		return Anchor_Type(Anchor_Type_value[anchorType.(string)]), nil
	case Anchor_Type:
		return anchorType.(Anchor_Type), nil
	default:
		return Anchor_ETH, errors.New("invalid anchorType provided")
	}
}
