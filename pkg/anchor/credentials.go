package anchor

import (
	"context"

	"google.golang.org/grpc/credentials"
)

// Credentials implements perRPC credentials.
type ServiceCredentials struct {
	credentials string
	secure      bool
}

// WithServiceCredentials will return an implementation of the PerRPCCredentials
// that will be used to attach your credentials to every RPC call.
func NewServiceCredentials(credentials string, secure bool) credentials.PerRPCCredentials {
	return &ServiceCredentials{credentials: credentials, secure: secure}
}

func (s *ServiceCredentials) GetRequestMetadata(ctx context.Context, uri ...string) (map[string]string, error) {
	return map[string]string{
		"authorization": "Bearer " + s.credentials,
	}, nil
}

// RequireTransportSecurity indicates whether the credentials requires
// transport security.
func (s *ServiceCredentials) RequireTransportSecurity() bool {
	return s.secure
}
