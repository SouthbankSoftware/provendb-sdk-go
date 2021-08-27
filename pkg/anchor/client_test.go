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
		t.Fatal(err)
	}
	if p.Format != Proof_CHP_PATH.String() {
		t.Fatal("wrong default format")
	}
	if p.AnchorType != Anchor_ETH.String() {
		t.Fatal("wrong default anchor type")
	}
	confirmed := false
	callback := func(p *AnchorProof, err error) {
		if err != nil {
			t.Fatal(err)
			return
		}
		if p.Status == Batch_CONFIRMED.String() {
			confirmed = true
			return
		}
	}
	client.SubscribeProof(context.Background(), p.Id, p.AnchorType, callback)
	for !confirmed {
	}
}

func TestClient_SubmitProofWithAwaitConfirmed(t *testing.T) {
	client, err := Connect(WithInsecure(true), WithAddress("localhost:10008"), WithCredentials("eyJhbGciOiJIUzI1NiIsInR5cCI6IkpXVCJ9.eyJhdWQiOiJhbmNob3IiLCJleHAiOjE3Nzk4NDgyODcsImp0aSI6ImFzMmNic3U1ODluZzU2ZzZ1em11bm8zOSIsInN1YiI6InByb3h5Iiwic2NvcGUiOiIwIiwicm9sZSI6IlBhaWQifQ.yVqpF8PvHVagAuOhQPVVWHDbzpnOZe_PoiYggsbFnwI"))
	if err != nil {
		t.Fail()
	}
	defer client.Close()
	p, err := client.SubmitProof(context.Background(), "dffd6021bb2bd5b0af676290809ec3a53191dd81c7f70a4b28688a362182986f", SubmitProofWithAwaitConfirmed(true))
	if err != nil {
		t.Fatal(err)
	}
	if p.Status != Batch_CONFIRMED.String() {
		t.Fatal("proof should have been confirmed")
	}
}
