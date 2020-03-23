package auth_test

import (
	"context"
	"errors"
	"fmt"
	"testing"

	"github.com/stretchr/testify/require"
	"github.com/videocoin/common/grpcutil/auth"
)

func TestRBACAuthZ(t *testing.T) {
	testCases := []struct {
		name       string
		ctx        context.Context
		principal  string
		fullMethod string
		err        error
	}{
		{
			name:       "valid",
			ctx:        nil,
			principal:  "2bb6bd99-c27c-48ea-85bc-a8afbc9f35cb",
			fullMethod: "/videocoin.iam.v1.IAM/CreateKey",
			err:        nil,
		},
		{
			name:       "permission denied",
			ctx:        nil,
			principal:  "2bb6bd99-c27c-48ea-85bc-a8afbc9f35cb",
			fullMethod: "/videocoin.iam.v1.IAM/DeleteKey",
			err:        errors.New("Permission iam.serviceAccountKeys.delete is required to perform this operation on account 2bb6bd99-c27c-48ea-85bc-a8afbc9f35cb"),
		},
	}
	for _, testCase := range testCases {
		t.Run(fmt.Sprintf("name: %s\n", testCase.name), func(t *testing.T) {
			require.Equal(t, testCase.err, auth.RBACAuthZ()(testCase.ctx, testCase.principal, testCase.fullMethod))
		})
	}
}
