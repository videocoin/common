package auth_test

import (
	"context"
	"fmt"
	"testing"

	"github.com/grpc-ecosystem/go-grpc-middleware/util/metautils"
	"github.com/stretchr/testify/require"
	"github.com/videocoin/common/grpcutil/auth"
)

func TestJWTAuthNZ(t *testing.T) {
	testCases := []struct {
		name       string
		authHeader string
		secret     string
		subject    string
		context    context.Context
		err        error
	}{
		{
			name:       "valid token with HMAC protection",
			authHeader: "Bearer eyJhbGciOiJIUzI1NiIsInR5cCI6IkpXVCJ9.eyJzdWIiOiIxMjM0NTY3ODkwIiwibmFtZSI6IkpvaG4gRG9lIiwiaWF0IjoxNTE2MjM5MDIyfQ.5mhBHqs5_DTLdINd9p5m7ZJ6XD0Xc55kIaCRY5r6HRA",
			secret:     "test",
			subject:    "1234567890",
			context:    nil,
			err:        nil,
		},
	}
	for _, testCase := range testCases {
		t.Run(fmt.Sprintf("name: %s\n", testCase.name), func(t *testing.T) {
			ctx := make(metautils.NiceMD).Set("authorization", testCase.authHeader).ToIncoming(context.Background())
			ctx, err := auth.JWTAuthNZ("", "", testCase.secret)(ctx)
			require.Equal(t, testCase.err, err)
			require.NotNil(t, ctx)
			if err == nil {
				require.NotNil(t, ctx.Value("token"))
			}
		})
	}
}
