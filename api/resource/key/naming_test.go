package key_test

import (
	"testing"

	sa "github.com/videocoin/common/api/resource/serviceaccount"

	"github.com/stretchr/testify/require"
	"github.com/videocoin/common/api/resource/key"
	"github.com/videocoin/common/api/resource/project"
)

func TestName(t *testing.T) {
	tests := []struct {
		name                    string
		projID, accEmail, keyID string
		output                  string
		err                     error
	}{
		{
			name:     "invalid account email: id must have at least 3 chars",
			projID:   "videocoin-123",
			accEmail: "acc@videocoin-123.vserviceaccount.com",
			keyID:    "5325f2f9-a193-4b59-a539-8c06beb2eeb5",
			output:   "",
			err:      sa.ErrInvalidEmail,
		},
		{
			name:     "invalid project id: starts with dash",
			projID:   "-videocoin-123",
			accEmail: "account1@videocoin-123.vserviceaccount.com",
			keyID:    "5325f2f9-a193-4b59-a539-8c06beb2eeb5",
			output:   "",
			err:      project.ErrInvalidID,
		},
		{
			name:     "invalid key id",
			projID:   "videocoin-123",
			accEmail: "account1@videocoin-123.vserviceaccount.com",
			keyID:    "a193-4b59-a539-8c06beb2eeb5",
			output:   "",
			err:      key.ErrInvalidID,
		},
		{
			name:     "valid project id, account email and key id",
			projID:   "videocoin-123",
			accEmail: "account1@videocoin-123.vserviceaccount.com",
			keyID:    "5325f2f9-a193-4b59-a539-8c06beb2eeb5",
			output:   "projects/videocoin-123/serviceAccounts/account1@videocoin-123.vserviceaccount.com/keys/5325f2f9-a193-4b59-a539-8c06beb2eeb5",
			err:      nil,
		},
	}
	for _, test := range tests {
		t.Run(test.name, func(t *testing.T) {
			output, err := key.Name(test.projID, test.accEmail, test.keyID)
			require.Equal(t, test.err, err)
			require.Equal(t, test.output, output)
		})
	}
}

func TestIDFromName(t *testing.T) {
	tests := []struct {
		name    string
		keyName string
		output  string
		err     error
	}{
		{
			name:    "invalid name: invalid key id",
			keyName: "projects/videocoin-123/serviceAccounts/account1@videocoin-123.vserviceaccount.com/keys/a193-4b59-a539-8c06beb2eeb5",
			output:  "",
			err:     key.ErrInvalidName,
		},
		{
			name:    "valid name",
			keyName: "projects/videocoin-123/serviceAccounts/account1@videocoin-123.vserviceaccount.com/keys/5325f2f9-a193-4b59-a539-8c06beb2eeb5",
			output:  "5325f2f9-a193-4b59-a539-8c06beb2eeb5",
			err:     nil,
		},
	}
	for _, test := range tests {
		t.Run(test.name, func(t *testing.T) {
			output, err := key.IDFromName(test.keyName)
			require.Equal(t, test.err, err)
			require.Equal(t, test.output, output)
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

func TestIsValidName(t *testing.T) {
	tests := []struct {
		name    string
		keyName string
		output  bool
	}{
		{
			name:    "invalid name: invalid key id",
			keyName: "projects/videocoin-123/serviceAccounts/account1@videocoin-123.vserviceaccount.com/keys/a193-4b59-a539-8c06beb2eeb5",
			output:  false,
		},
		{
			name:    "valid name",
			keyName: "projects/videocoin-123/serviceAccounts/account1@videocoin-123.vserviceaccount.com/keys/5325f2f9-a193-4b59-a539-8c06beb2eeb5",
			output:  true,
		},
	}
	for _, test := range tests {
		t.Run(test.name, func(t *testing.T) {
			require.Equal(t, test.output, key.IsValidName(test.keyName))
		})
	}
}
