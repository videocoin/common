package serviceaccount

import (
	"regexp"
	"strings"

	cstr "github.com/videocoin/common/strings"

	"github.com/videocoin/common/api/resource"
	"github.com/videocoin/common/api/resource/project"
)

const (
	// CollectionID is the identifier of the resource that contains a list of
	// service accounts.
	CollectionID = "serviceAccounts"

	// Domain represents the service account email domain.
	Domain = "vserviceaccount.com"
)

var (
	// ErrInvalidID indicates that the service account id is invalid.
	ErrInvalidID = resource.PatternError(IDPattern.String())
	// ErrInvalidEmail indicates that the service account email is invalid.
	ErrInvalidEmail = resource.PatternError(EmailPattern.String())
	// ErrInvalidName indicates that the service account name is invalid.
	ErrInvalidName = resource.PatternError(NamePattern.String())
)

var (
	// IDPattern is the service account identifier pattern.
	IDPattern = regexp.MustCompile(`^[a-z][-a-z0-9]{4,28}[a-z0-9]$`)
	// EmailPattern is the service account email pattern.
	EmailPattern = regexp.MustCompile(`^[a-z][-a-z0-9]{4,28}[a-z0-9]@[a-z][-a-z0-9]{3,48}[a-z0-9]\.vserviceaccount\.com$`)
	// NamePattern is the service account name pattern.
	NamePattern = regexp.MustCompile(`^projects/[a-z][-a-z0-9]{3,48}[a-z0-9]/serviceAccounts/[a-z][-a-z0-9]{4,28}[a-z0-9]@[a-z][-a-z0-9]{3,48}[a-z0-9]\.vserviceaccount\.com$`)
)

// Name returns the service account's name given a project identifier and a
// service account identifier.
func Name(projID string, accEmail string) (string, error) {
	projName, err := project.Name(projID)
	if err != nil {
		return "", err
	}
	if ok := IsValidEmail(accEmail); !ok {
		return "", ErrInvalidEmail
	}

	return cstr.JoinWithSeparator(resource.NameSeparator, projName, CollectionID, accEmail), nil
}

// Email returns the service account's email given a project identifier and a
// service account identifier.
func Email(projID string, accID string) (string, error) {
	if ok := project.IsValidID(projID); !ok {
		return "", project.ErrInvalidID
	}
	if ok := IsValidID(accID); !ok {
		return "", ErrInvalidID
	}

	return cstr.Join(accID, "@", projID, ".", Domain), nil
}

// EmailFromName derives the service account's email from its name.
func EmailFromName(name string) (string, error) {
	if ok := IsValidName(name); !ok {
		return "", ErrInvalidName
	}
	return strings.SplitN(name, resource.NameSeparator, 4)[3], nil
}

// IsValidID reports whether a service account identifier is valid.
func IsValidID(ID string) bool {
	return IDPattern.MatchString(ID)
}

// IsValidEmail reports whether a service account email is valid.
func IsValidEmail(name string) bool {
	return EmailPattern.MatchString(name)
}

// IsValidName reports whether a service account name is valid.
func IsValidName(name string) bool {
	return NamePattern.MatchString(name)
}
