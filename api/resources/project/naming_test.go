package project_test

import (
	"fmt"
	"testing"

	"github.com/stretchr/testify/require"
	"github.com/videocoin/cloud-pkg/api/resources/project"
)

func TestNameID(t *testing.T) {
	tests := []struct {
		projName project.Name
		output   string
	}{
		{
			projName: project.Name("projects/5325f2f9-a193-4b59-a539-8c06beb2eeb5"),
			output:   "5325f2f9-a193-4b59-a539-8c06beb2eeb5",
		},
	}
	for _, test := range tests {
		t.Run(fmt.Sprintf("project name: %s\n", test.projName), func(t *testing.T) {
			require.Equal(t, test.output, test.projName.ID())
		})
	}
}

func TestParseName(t *testing.T) {
	tests := []struct {
		name     string
		projName string
		output   project.Name
		err      error
	}{
		{
			name:     "invalid name: starts with a dash",
			projName: "projects/-123123",
			output:   "",
			err:      project.ErrInvalidName,
		},
		{
			name:     "invalid name: ends with a dash",
			projName: "projects/videocoin-123-",
			output:   "",
			err:      project.ErrInvalidName,
		},
		{
			name:     "invalid name: incorrect collection id",
			projName: "keys/videocoin-123",
			output:   "",
			err:      project.ErrInvalidName,
		},
		{
			name:     "invalid name: incorrect collection id part 2",
			projName: "project/videocoin-123",
			output:   "",
			err:      project.ErrInvalidName,
		},
		{
			name:     "invalid name: no dash",
			projName: "projectavideocoin-123",
			output:   "",
			err:      project.ErrInvalidName,
		},
		{
			name:     "valid name",
			projName: "projects/videocoin-123",
			output:   project.Name("projects/videocoin-123"),
			err:      nil,
		},
	}
	for _, test := range tests {
		t.Run(test.name, func(t *testing.T) {
			output, err := project.ParseName(test.projName)
			require.Equal(t, test.err, err)
			require.Equal(t, test.output, output)
		})
	}
}

func TestNewName(t *testing.T) {
	tests := []struct {
		projID string
		output project.Name
	}{
		{
			projID: "5325f2f9-a193-4b59-a539-8c06beb2eeb5",
			output: project.Name("projects/5325f2f9-a193-4b59-a539-8c06beb2eeb5"),
		},
	}
	for _, test := range tests {
		t.Run(fmt.Sprintf("project id: %s\n", test.projID), func(t *testing.T) {
			require.Equal(t, test.output, project.NewName(test.projID))
		})
	}
}

func TestNewNameWildcard(t *testing.T) {
	require.Equal(t, project.Name("projects/-"), project.NewNameWildcard())
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
