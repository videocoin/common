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
			authHeader: "Bearer eyJhbGciOiJIUzI1NiIsInR5cCI6IkpXVCJ9.eyJzdWIiOiIyYmI2YmQ5OS1jMjdjLTQ4ZWEtODViYy1hOGFmYmM5ZjM1Y2IiLCJuYW1lIjoiSm9obiBEb2UiLCJpYXQiOjE1MTYyMzkwMjJ9.StMn9-Nw_4xi635jsZgVWsomQaCo8W5rwjGr1MikYNM",
			secret:     "test",
			output:     "2bb6bd99-c27c-48ea-85bc-a8afbc9f35cb",
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
