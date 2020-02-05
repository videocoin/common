package project_test

import (
	"testing"

	"github.com/stretchr/testify/require"
	"github.com/videocoin/common/api/resource/project"
)

func TestName(t *testing.T) {
	tests := []struct {
		name   string
		projID string
		output string
		err    error
	}{
		{
			name:   "invalid id: starts with a dash",
			projID: "-123123",
			output: "",
			err:    project.ErrInvalidID,
		},
		{
			name:   "valid id: uuid",
			projID: "5325f2f9-a193-4b59-a539-8c06beb2eeb5",
			output: "projects/5325f2f9-a193-4b59-a539-8c06beb2eeb5",
			err:    nil,
		},
	}
	for _, test := range tests {
		t.Run(test.name, func(t *testing.T) {
			output, err := project.Name(test.projID)
			require.Equal(t, test.err, err)
			require.Equal(t, test.output, output)
		})
	}
}

func TestIDFromName(t *testing.T) {
	tests := []struct {
		name     string
		projName string
		output   string
		err      error
	}{
		{
			name:     "invalid name: incorrect collection id part 2",
			projName: "project/videocoin-123",
			output:   "",
			err:      project.ErrInvalidName,
		},
		{
			name:     "valid name",
			projName: "projects/5325f2f9-a193-4b59-a539-8c06beb2eeb5",
			output:   "5325f2f9-a193-4b59-a539-8c06beb2eeb5",
			err:      nil,
		},
	}
	for _, test := range tests {
		t.Run(test.name, func(t *testing.T) {
			output, err := project.IDFromName(test.projName)
			require.Equal(t, test.err, err)
			require.Equal(t, test.output, output)
		})
	}
}

func TestIsValidID(t *testing.T) {
	tests := []struct {
		name   string
		projID string
		output bool
	}{
		{
			name:   "invalid id: starts with a dash",
			projID: "-123123",
			output: false,
		},
		{
			name:   "invalid id: ends with a dash",
			projID: "videocoin-123-",
			output: false,
		},
		{
			name:   "valid id: uuid",
			projID: "5325f2f9-a193-4b59-a539-8c06beb2eeb5",
			output: true,
		},
		{
			name:   "valid id",
			projID: "videocoin-123",
			output: true,
		},
	}
	for _, test := range tests {
		t.Run(test.name, func(t *testing.T) {
			require.Equal(t, test.output, project.IsValidID(test.projID))
		})
	}
}

func TestIsValidName(t *testing.T) {
	tests := []struct {
		name     string
		projName string
		output   bool
	}{
		{
			name:     "invalid name: starts with a dash",
			projName: "projects/-123123",
			output:   false,
		},
		{
			name:     "invalid name: ends with a dash",
			projName: "projects/videocoin-123-",
			output:   false,
		},
		{
			name:     "invalid name: incorrect collection id",
			projName: "keys/videocoin-123",
			output:   false,
		},
		{
			name:     "invalid name: incorrect collection id part 2",
			projName: "project/videocoin-123",
			output:   false,
		},
		{
			name:     "invalid name: no dash",
			projName: "projectavideocoin-123",
			output:   false,
		},
		{
			name:     "valid name",
			projName: "projects/videocoin-123",
			output:   true,
		},
	}
	for _, test := range tests {
		t.Run(test.name, func(t *testing.T) {
			require.Equal(t, test.output, project.IsValidName(test.projName))
		})
	}
}
