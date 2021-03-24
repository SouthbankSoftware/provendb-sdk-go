package anchor

type AnchorHandle struct {
	client AnchorServiceClient
}

func NewAnchorHandle(address string, credentials string) *AnchorHandle {
	return &AnchorHandle{}
}
