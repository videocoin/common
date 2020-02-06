package key_test

import (
	"testing"

	"github.com/stretchr/testify/require"
	"github.com/videocoin/common/api/resources/key"
)

func TestNewName(t *testing.T) {
	tests := []struct {
		name                    string
		projID, accEmail, keyID string
		output                  key.Name
	}{
		{
			name:     "valid project id, account email and key id",
			projID:   "videocoin-123",
			accEmail: "account1@videocoin-123.vserviceaccount.com",
			keyID:    "5325f2f9-a193-4b59-a539-8c06beb2eeb5",
			output:   key.Name("projects/videocoin-123/serviceAccounts/account1@videocoin-123.vserviceaccount.com/keys/5325f2f9-a193-4b59-a539-8c06beb2eeb5"),
		},
	}
	for _, test := range tests {
		t.Run(test.name, func(t *testing.T) {
			require.Equal(t, test.output, key.NewName(test.projID, test.accEmail, test.keyID))
		})
	}
}

func TestNewNameWildcard(t *testing.T) {
	tests := []struct {
		name            string
		accEmail, keyID string
		output          key.Name
	}{
		{
			name:     "valid account email and key id",
			accEmail: "account1@videocoin-123.vserviceaccount.com",
			keyID:    "5325f2f9-a193-4b59-a539-8c06beb2eeb5",
			output:   key.Name("projects/-/serviceAccounts/account1@videocoin-123.vserviceaccount.com/keys/5325f2f9-a193-4b59-a539-8c06beb2eeb5"),
		},
	}
	for _, test := range tests {
		t.Run(test.name, func(t *testing.T) {
			require.Equal(t, test.output, key.NewNameWildcard(test.accEmail, test.keyID))
		})
	}
}

func TestNameID(t *testing.T) {
	tests := []struct {
		name    string
		keyName key.Name
		output  string
	}{
		{
			name:    "valid key name",
			keyName: key.Name("projects/videocoin-123/serviceAccounts/account1@videocoin-123.vserviceaccount.com/keys/5325f2f9-a193-4b59-a539-8c06beb2eeb5"),
			output:  "5325f2f9-a193-4b59-a539-8c06beb2eeb5",
		},
	}
	for _, test := range tests {
		t.Run(test.name, func(t *testing.T) {
			require.Equal(t, test.output, test.keyName.ID())
		})
	}
}

func TestIsValidID(t *testing.T) {
	tests := []struct {
		name   string
		keyID  string
		output bool
	}{
		{
			name:   "invalid id",
			keyID:  "a193-4b59-a539-8c06beb2eeb5",
			output: false,
		},
		{
			name:   "valid id",
			keyID:  "5325f2f9-a193-4b59-a539-8c06beb2eeb5",
			output: true,
		},
	}
	for _, test := range tests {
		t.Run(test.name, func(t *testing.T) {
			require.Equal(t, test.output, key.IsValidID(test.keyID))
		})
	}
}

func TestParseName(t *testing.T) {
	tests := []struct {
		name    string
		keyName string
		output  key.Name
		err     error
	}{
		{
			name:    "invalid name: invalid key id",
			keyName: "projects/videocoin-123/serviceAccounts/account1@videocoin-123.vserviceaccount.com/keys/a193-4b59-a539-8c06beb2eeb5",
			output:  "",
			err:     key.ErrInvalidName,
		},
		{
			name:    "valid name wildcard",
			keyName: "projects/-/serviceAccounts/account1@videocoin-123.vserviceaccount.com/keys/5325f2f9-a193-4b59-a539-8c06beb2eeb5",
			output:  key.Name("projects/-/serviceAccounts/account1@videocoin-123.vserviceaccount.com/keys/5325f2f9-a193-4b59-a539-8c06beb2eeb5"),
			err:     nil,
		},
		{
			name:    "valid name",
			keyName: "projects/videocoin-123/serviceAccounts/account1@videocoin-123.vserviceaccount.com/keys/5325f2f9-a193-4b59-a539-8c06beb2eeb5",
			output:  key.Name("projects/videocoin-123/serviceAccounts/account1@videocoin-123.vserviceaccount.com/keys/5325f2f9-a193-4b59-a539-8c06beb2eeb5"),
			err:     nil,
		},
	}
	for _, test := range tests {
		t.Run(test.name, func(t *testing.T) {
			output, err := key.ParseName(test.keyName)
			require.Equal(t, test.err, err)
			require.Equal(t, test.output, output)
		})
	}
}
