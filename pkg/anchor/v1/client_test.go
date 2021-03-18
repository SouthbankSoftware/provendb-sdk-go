package anchor

import (
	"testing"
)

func TestNewClient(t *testing.T) {
	_, _ = NewClient(
		WithSecure(true),
		WithAddress("localhost:8000"),
		WithCredentials("n2m8a91llaks91mnsmla"))

}
