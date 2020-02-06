package serviceaccount_test

import (
	"fmt"
	"testing"

	sa "github.com/videocoin/common/api/resources/serviceaccount"

	"github.com/stretchr/testify/require"
)

func TestNewName(t *testing.T) {
	tests := []struct {
		projID, accEmail string
		output           sa.Name
	}{
		{
			projID:   "videocoin-123",
			accEmail: "account1@videocoin-123.vserviceaccount.com",
			output:   sa.Name("projects/videocoin-123/serviceAccounts/account1@videocoin-123.vserviceaccount.com"),
		},
	}
	for _, test := range tests {
		t.Run(fmt.Sprintf("proj id: %s, acc email: %s\n", test.projID, test.accEmail), func(t *testing.T) {
			require.Equal(t, test.output, sa.NewName(test.projID, test.accEmail))
		})
	}
}

func TestNewNameWildcard(t *testing.T) {
	tests := []struct {
		accEmail string
		output   sa.Name
	}{
		{
			accEmail: "account1@videocoin-123.vserviceaccount.com",
			output:   sa.Name("projects/-/serviceAccounts/account1@videocoin-123.vserviceaccount.com"),
		},
	}
	for _, test := range tests {
		t.Run(fmt.Sprintf("acc email: %s\n", test.accEmail), func(t *testing.T) {
			require.Equal(t, test.output, sa.NewNameWildcard(test.accEmail))
		})
	}
}

func TestNewEmail(t *testing.T) {
	tests := []struct {
		projID, accID string
		output        string
	}{
		{
			projID: "videocoin-123",
			accID:  "account1",
			output: "account1@videocoin-123.vserviceaccount.com",
		},
	}
	for _, test := range tests {
		t.Run(fmt.Sprintf("proj id: %s, acc id: %s\n", test.projID, test.accID), func(t *testing.T) {
			require.Equal(t, test.output, sa.NewEmail(test.projID, test.accID))
		})
	}
}

func TestNameEmail(t *testing.T) {
	tests := []struct {
		saName sa.Name
		output string
	}{
		{
			saName: sa.Name("projects/videocoin-123/serviceAccounts/account1@videocoin-123.vserviceaccount.com"),
			output: "account1@videocoin-123.vserviceaccount.com",
		},
	}
	for _, test := range tests {
		t.Run(fmt.Sprintf("acc name: %s\n", string(test.saName)), func(t *testing.T) {
			require.Equal(t, test.output, test.saName.Email())
		})
	}
}

func TestIsValidID(t *testing.T) {
	tests := []struct {
		name   string
		saID   string
		output bool
	}{
		{
			name:   "invalid id: must have at least 6 chars",
			saID:   "accou",
			output: false,
		},
		{
			name:   "invalid id: must have less than 30 chars",
			saID:   "account1sadsfafsdfasdfadsfasdfs",
			output: false,
		},
		{
			name:   "valid id",
			saID:   "account1",
			output: true,
		},
	}
	for _, test := range tests {
		t.Run(test.name, func(t *testing.T) {
			require.Equal(t, test.output, sa.IsValidID(test.saID))
		})
	}
}

func TestIsValidEmail(t *testing.T) {
	tests := []struct {
		name    string
		saEmail string
		output  bool
	}{
		{
			name:    "invalid email: incorrect domain",
			saEmail: "account1@videocoin-123.gserviceaccount.com",
			output:  false,
		},
		{
			name:    "valid email",
			saEmail: "account1@videocoin-123.vserviceaccount.com",
			output:  true,
		},
	}
	for _, test := range tests {
		t.Run(test.name, func(t *testing.T) {
			require.Equal(t, test.output, sa.IsValidEmail(test.saEmail))
		})
	}
}

func TestParseName(t *testing.T) {
	tests := []struct {
		name   string
		saName string
		output sa.Name
		err    error
	}{
		{
			name:   "invalid name: incorrect collection id",
			saName: "projects/videocoin-123/serviceAccount/account1@videocoin-123.vserviceaccount.com",
			output: "",
			err:    sa.ErrInvalidName,
		},
		{
			name:   "invalid name: incorrect domain",
			saName: "projects/videocoin-123/serviceAccounts/account1@videocoin-123.gserviceaccount.com",
			output: "",
			err:    sa.ErrInvalidName,
		},
		{
			name:   "invalid name: missing @",
			saName: "projects/videocoin-123/serviceAccounts/account1videocoin-123.vserviceaccount.com",
			output: "",
			err:    sa.ErrInvalidName,
		},
		{
			name:   "valid name wildcard",
			saName: "projects/-/serviceAccounts/account1@videocoin-123.vserviceaccount.com",
			output: sa.Name("projects/-/serviceAccounts/account1@videocoin-123.vserviceaccount.com"),
			err:    nil,
		},
		{
			name:   "valid name",
			saName: "projects/videocoin-123/serviceAccounts/account1@videocoin-123.vserviceaccount.com",
			output: sa.Name("projects/videocoin-123/serviceAccounts/account1@videocoin-123.vserviceaccount.com"),
			err:    nil,
		},
	}
	for _, test := range tests {
		t.Run(test.name, func(t *testing.T) {
			output, err := sa.ParseName(test.saName)
			require.Equal(t, test.err, err)
			require.Equal(t, test.output, output)
		})
	}
}
