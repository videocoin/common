package service_test

import (
	"testing"

	"github.com/stretchr/testify/require"
	"github.com/videocoin/cloud-pkg/api/resources/service"
)

func TestIsValidName(t *testing.T) {
	tests := []struct {
		name    string
		svcName string
		output  bool
	}{
		{
			name:    "invalid name: suffix must be videocoin.network",
			svcName: "iam.network.videocoin",
			output:  false,
		},
		{
			name:    "invalid name: suffix must be videocoin.network (test special character part 1)",
			svcName: "iam.videocoinanetwork",
			output:  false,
		},
		{
			name:    "invalid name: suffix must be videocoin.network (test special character part 2)",
			svcName: "iamavideocoin.network",
			output:  false,
		},
		{
			name:    "invalid name: name has numbers",
			svcName: "iam123.network.videocoin",
			output:  false,
		},
		{
			name:    "invalid name: id must have at least 3 characters",
			svcName: "ia.network.videocoin",
			output:  false,
		},
		{
			name:    "invalid name: id must have at less than 20 characters",
			svcName: "iasalfsdfdiaoshfoasdh.network.videocoin",
			output:  false,
		},
		{
			name:    "valid name",
			svcName: "iam.videocoin.network",
			output:  true,
		},
	}
	for _, test := range tests {
		t.Run(test.name, func(t *testing.T) {
			require.Equal(t, test.output, service.IsValidName(test.svcName))
		})
	}
}
