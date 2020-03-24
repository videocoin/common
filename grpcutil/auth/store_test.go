package auth_test

import (
	"context"
	"testing"

	"github.com/grpc-ecosystem/go-grpc-middleware/util/metautils"
	"github.com/stretchr/testify/require"
	"github.com/videocoin/common/grpcutil/auth"
)

func TestGetUserRole(t *testing.T) {
	rstore := auth.NewRoleCache(auth.NewRoleRepo())
	ctx := make(metautils.NiceMD).Set("authorization", "Bearer eyJhbGciOiJIUzI1NiIsInR5cCI6IkpXVCJ9.eyJzdWIiOiI2NWIxMTA3Yy00N2E3LTQyOWItNjBkNS0yMThlZTg2OGYwODAifQ.1Rv-69do9Akb6lqJYKyWw40YxxEWTWR6AaUoBH0EvR0").ToIncoming(context.Background())
	role, err := rstore.GetUserRole(ctx, "65b1107c-47a7-429b-60d5-218ee868f080")
	require.NoError(t, err)
	require.Equal(t, "USER_ROLE_MINER", role)
}

func TestGetRole(t *testing.T) {
	rstore := auth.NewRoleCache(auth.NewRoleRepo())
	role, err := rstore.GetRole(nil, "USER_ROLE_MINER")
	require.NoError(t, err)
	require.Equal(t, &auth.Role{IncludedPermissions: []string{
		"iam.serviceAccountKeys.create",
		"iam.serviceAccountKeys.list",
		"iam.serviceAccountKeys.get",
	}}, role)
}
