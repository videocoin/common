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

func TestJWTAuthNZ(t *testing.T) {
	testCases := []struct {
		name       string
		authHeader string
		secret     string
		fullMethod string
		err        error
	}{
		{
			name:       "HMAC - valid token",
			authHeader: "Bearer eyJhbGciOiJIUzI1NiIsInR5cCI6IkpXVCJ9.eyJzdWIiOiIyYmI2YmQ5OS1jMjdjLTQ4ZWEtODViYy1hOGFmYmM5ZjM1Y2IiLCJuYW1lIjoiSm9obiBEb2UiLCJpYXQiOjE1MTYyMzkwMjJ9.StMn9-Nw_4xi635jsZgVWsomQaCo8W5rwjGr1MikYNM",
			secret:     "test",
			fullMethod: "/videocoin.iam.v1.IAM/CreateKey",
			err:        nil,
		},
		{
			name:       "HMAC - permission denied",
			authHeader: "Bearer eyJhbGciOiJIUzI1NiIsInR5cCI6IkpXVCJ9.eyJzdWIiOiIyYmI2YmQ5OS1jMjdjLTQ4ZWEtODViYy1hOGFmYmM5ZjM1Y2IiLCJuYW1lIjoiSm9obiBEb2UiLCJpYXQiOjE1MTYyMzkwMjJ9.StMn9-Nw_4xi635jsZgVWsomQaCo8W5rwjGr1MikYNM",
			secret:     "test",
			fullMethod: "/videocoin.iam.v1.IAM/DeleteKey",
			err:        status.Error(codes.PermissionDenied, "Permission iam.serviceAccountKeys.delete is required to perform this operation on account 2bb6bd99-c27c-48ea-85bc-a8afbc9f35cb"),
		},
	}
	for _, testCase := range testCases {
		t.Run(fmt.Sprintf("name: %s\n", testCase.name), func(t *testing.T) {
			ctx := make(metautils.NiceMD).Set("authorization", testCase.authHeader).ToIncoming(context.Background())
			ctx, err := auth.JWTAuthNZ("", "", testCase.secret)(ctx, testCase.fullMethod)
			require.Equal(t, testCase.err, err)
			if err == nil {
				require.NotNil(t, ctx)
				require.NotNil(t, ctx.Value("token"))
			}
		})
	}
}
