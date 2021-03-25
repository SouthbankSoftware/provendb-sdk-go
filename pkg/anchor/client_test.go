package anchor

import (
	"context"
	"testing"
)

func TestConnect(t *testing.T) {
	a, e := Connect()
	p, er := a.SubmitProof(context.Background(), "", SubmitProofWithAnchorType(Anchor_BTC))
}
