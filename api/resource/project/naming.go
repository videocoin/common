package project

import (
	"regexp"
	"strings"

	cexp "github.com/videocoin/common/regexp"
	cstr "github.com/videocoin/common/strings"

	"github.com/videocoin/common/api/resource"
)

const (
	// IDPattern represents the project identifier pattern.
	IDPattern = "[a-z][-a-z0-9]{3,48}[a-z0-9]"

	// CollectionID is the identifier of the resource that contains a list of
	// projects.
	CollectionID = "projects"
)

var (
	// ErrInvalidName indicates that the project name is invalid.
	ErrInvalidName = resource.PatternError(NamePattern)
	// ErrInvalidID indicates that the project identifier is invalid.
	ErrInvalidID = resource.PatternError(IDPattern)
)

var (
	// NamePattern represents the project name pattern.
	NamePattern = cstr.Join(CollectionID, resource.NameSeparator, IDPattern)
	nameRegExp  = regexp.MustCompile(cstr.Join(cexp.Begin, NamePattern, cexp.End))
	idRegExp    = regexp.MustCompile(cstr.Join(cexp.Begin, IDPattern, cexp.End))
)

// Name returns the project name given a project identifier.
func Name(projID string) string {
	return cstr.Join(CollectionID, resource.NameSeparator, projID)
}

// IDFromName derives the project identifier from the projects's
// name.
func IDFromName(name string) string {
	return strings.SplitN(name, resource.NameSeparator, 2)[1]
}

// IsValidName verifies whether the the project name is valid or not.
func IsValidName(name string) bool {
	return nameRegExp.MatchString(name)
}

// IsValidID verifies whether the the project identifier is valid or not.
func IsValidID(ID string) bool {
	return idRegExp.MatchString(ID)
}
