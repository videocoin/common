package auth_test

import (
	"context"
	"fmt"
	"testing"

	"github.com/grpc-ecosystem/go-grpc-middleware/util/metautils"
	"github.com/stretchr/testify/require"
	"github.com/videocoin/common/grpcutil/auth"
	"google.golang.org/grpc/codes"
	"google.golang.org/grpc/status"
)

func TestHMACAuthN(t *testing.T) {
	testCases := []struct {
		authHeader string
		secret     string
		output     string
		err        error
	}{
		{
			authHeader: "",
			secret:     "",
			output:     "",
			err:        status.Errorf(codes.Unauthenticated, "Request unauthenticated with bearer"),
		},
		{
			authHeader: "Basic eyJhbGciOiJIUzI1NiIsInR5cCI6IkpXVCJ9.eyJzdWIiOiIxMjM0NTY3ODkwIiwibmFtZSI6IkpvaG4gRG9lIiwiaWF0IjoxNTE2MjM5MDIyfQ.5mhBHqs5_DTLdINd9p5m7ZJ6XD0Xc55kIaCRY5r6HRA",
			secret:     "test",
			output:     "",
			err:        status.Errorf(codes.Unauthenticated, "Request unauthenticated with bearer"),
		},
		{
			authHeader: "Bearer eyJhbGciOiJIUzI1NiIsInR5cCI6IkpXVCJ9.eyJzdWIiOiIxMjM0NTY3ODkwIiwibmFtZSI6IkpvaG4gRG9lIiwiaWF0IjoxNTE2MjM5MDIyfQ.5mhBHqs5_DTLdINd9p5m7ZJ6XD0Xc55kIaCRY5r6HRA",
			secret:     "test",
			output:     "1234567890",
			err:        nil,
		},
	}
	for _, testCase := range testCases {
		t.Run(fmt.Sprintf("auth header: %s, secret: %s\n", testCase.authHeader, testCase.secret), func(t *testing.T) {
			ctx := make(metautils.NiceMD).Set("authorization", testCase.authHeader).ToIncoming(context.Background())
			sub, err := auth.HMACAuthN(testCase.secret)(ctx)
			require.Equal(t, testCase.err, err)
			require.Equal(t, testCase.output, sub)
		})
	}
}

func TestServiceAccountAuthN(t *testing.T) {}
