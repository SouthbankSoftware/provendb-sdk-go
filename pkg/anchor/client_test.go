package anchor

import (
	"context"
	"testing"
)

func TestConnect(t *testing.T) {
	client, err := Connect(WithInsecure(true), WithAddress("localhost:10008"))
	if err != nil {
		t.Fail()
	}
	defer client.Close()
}

func TestClient_GetAnchors(t *testing.T) {
	client, err := Connect(WithInsecure(true), WithAddress("localhost:10008"))
	if err != nil {
		t.Fail()
	}
	defer client.Close()
	anchors, err := client.GetAnchors(context.Background())
	if err != nil {
		t.FailNow()
	}
	// There should always be a length above 0.
	if !(len(anchors) > 0) {
		t.FailNow()
	}
}

func TestClient_SubmitProof(t *testing.T) {
	client, err := Connect(WithInsecure(true), WithAddress("localhost:10008"))
	if err != nil {
		t.Fail()
	}
	defer client.Close()
	p, err := client.SubmitProof(context.Background(), "dffd6021bb2bd5b0af676290809ec3a53191dd81c7f70a4b28688a362182986f")
	if err != nil {
		t.Fail()
	}
	if p.Format != Proof_CHP_PATH {
		t.Fatal("wrong default format")
	}
	if p.AnchorType != Anchor_ETH {
		t.Fatal("wrong default anchor type")
	}
	confirmed := false
	callback := func(p *Proof, err error) {
		if err != nil {
			t.FailNow()
		}
		if p.BatchStatus == Batch_ERROR {
			t.FailNow()
		}
		if p.BatchStatus == Batch_CONFIRMED {
			confirmed = true
			return
		}
	}
	client.SubscribeProof(context.Background(), p, callback)
	for !confirmed {
	}
}
