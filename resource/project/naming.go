package project

import (
	"strings"

	cstrings "github.com/videocoin/common/strings"

	"github.com/videocoin/common/resource"
)

// Name returns the project name given a project identifier.
func Name(projID string) string {
	return cstrings.Join("projects/", projID)
}

// IDFromName derives the project identifier from the projects's
// name.
func IDFromName(name string) string {
	return strings.SplitN(name, resource.NameSeparator, 2)[1]
}
