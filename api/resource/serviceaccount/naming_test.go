package serviceaccount_test

import (
	"testing"

	sa "github.com/videocoin/common/api/resource/serviceaccount"

	"github.com/stretchr/testify/require"
	"github.com/videocoin/common/api/resource/project"
)

func TestName(t *testing.T) {
	tests := []struct {
		name             string
		projID, accEmail string
		output           string
		err              error
	}{
		{
			name:     "invalid project id: starts with dash",
			projID:   "-videocoin-123",
			accEmail: "account1@videocoin-123.vserviceaccount.com",
			output:   "",
			err:      project.ErrInvalidID,
		},
		{
			name:     "invalid account email: id must have at least 6 chars",
			projID:   "videocoin-123",
			accEmail: "acc@videocoin-123.vserviceaccount.com",
			output:   "",
			err:      sa.ErrInvalidEmail,
		},
		{
			name:     "valid project id and service account email",
			projID:   "videocoin-123",
			accEmail: "account1@videocoin-123.vserviceaccount.com",
			output:   "projects/videocoin-123/serviceAccounts/account1@videocoin-123.vserviceaccount.com",
			err:      nil,
		},
	}
	for _, test := range tests {
		t.Run(test.name, func(t *testing.T) {
			output, err := sa.Name(test.projID, test.accEmail)
			require.Equal(t, test.err, err)
			require.Equal(t, test.output, output)
		})
	}
}

func TestEmail(t *testing.T) {
	tests := []struct {
		name          string
		projID, accID string
		output        string
		err           error
	}{
		{
			name:   "invalid project id: starts with dash",
			projID: "-videocoin-123",
			accID:  "account1",
			output: "",
			err:    project.ErrInvalidID,
		},
		{
			name:   "invalid account id: must have at least 3 chars",
			projID: "videocoin-123",
			accID:  "acc",
			output: "",
			err:    sa.ErrInvalidID,
		},
		{
			name:   "valid project id and service account id",
			projID: "videocoin-123",
			accID:  "account1",
			output: "account1@videocoin-123.vserviceaccount.com",
			err:    nil,
		},
	}
	for _, test := range tests {
		t.Run(test.name, func(t *testing.T) {
			output, err := sa.Email(test.projID, test.accID)
			require.Equal(t, test.err, err)
			require.Equal(t, test.output, output)
		})
	}
}

func TestEmailFromName(t *testing.T) {
	tests := []struct {
		name   string
		saName string
		output string
		err    error
	}{
		{
			name:   "invalid name: incorrect domain",
			saName: "projects/videocoin-123/serviceAccounts/account1@videocoin-123.gserviceaccount.com",
			output: "",
			err:    sa.ErrInvalidName,
		},
		{
			name:   "valid name",
			saName: "projects/videocoin-123/serviceAccounts/account1@videocoin-123.vserviceaccount.com",
			output: "account1@videocoin-123.vserviceaccount.com",
			err:    nil,
		},
	}
	for _, test := range tests {
		t.Run(test.name, func(t *testing.T) {
			output, err := sa.EmailFromName(test.saName)
			require.Equal(t, test.err, err)
			require.Equal(t, test.output, output)
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

func TestIsValidName(t *testing.T) {
	tests := []struct {
		name   string
		saName string
		output bool
	}{
		{
			name:   "invalid name: incorrect collection id",
			saName: "projects/videocoin-123/serviceAccount/account1@videocoin-123.vserviceaccount.com",
			output: false,
		},
		{
			name:   "invalid name: incorrect domain",
			saName: "projects/videocoin-123/serviceAccounts/account1@videocoin-123.gserviceaccount.com",
			output: false,
		},
		{
			name:   "invalid name: missing @",
			saName: "projects/videocoin-123/serviceAccounts/account1videocoin-123.vserviceaccount.com",
			output: false,
		},
		{
			name:   "valid name",
			saName: "projects/videocoin-123/serviceAccounts/account1@videocoin-123.vserviceaccount.com",
			output: true,
		},
	}
	for _, test := range tests {
		t.Run(test.name, func(t *testing.T) {
			require.Equal(t, test.output, sa.IsValidName(test.saName))
		})
	}
}
