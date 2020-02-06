package strings_test

import (
	"testing"

	"github.com/stretchr/testify/require"
	"github.com/videocoin/common/strings"
)

func TestJoin(t *testing.T) {
	tests := []struct {
		name   string
		strs   []string
		output string
	}{
		{
			name:   "empty",
			strs:   []string{},
			output: "",
		},
		{
			name:   "regular call",
			strs:   []string{"projects", "videocoin", "something"},
			output: "projectsvideocoinsomething",
		},
	}
	for _, test := range tests {
		t.Run(test.name, func(t *testing.T) {
			require.Equal(t, test.output, strings.Join(test.strs...))
		})
	}
}

func TestJoinWithSeparator(t *testing.T) {
	tests := []struct {
		name   string
		sep    string
		strs   []string
		output string
	}{
		{
			name:   "empty input",
			sep:    "/",
			strs:   []string{},
			output: "",
		},
		{
			name:   "empty separator",
			sep:    "",
			strs:   []string{"projects", "videocoin"},
			output: "projectsvideocoin",
		},
		{
			name:   "regular call",
			sep:    "/",
			strs:   []string{"projects", "videocoin"},
			output: "projects/videocoin",
		},
	}
	for _, test := range tests {
		t.Run(test.name, func(t *testing.T) {
			require.Equal(t, test.output, strings.JoinWithSeparator(test.sep, test.strs...))
		})
	}
}
