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
	// ErrInvalidUniqueID indicates that the service account id is invalid.
	ErrInvalidUniqueID = resource.PatternError(UniqueIDPattern.String())
	// ErrInvalidEmail indicates that the service account email is invalid.
	ErrInvalidEmail = resource.PatternError(EmailPattern.String())
	// ErrInvalidName indicates that the service account name is invalid.
	ErrInvalidName = resource.PatternError(NamePattern.String())
)

var (
	// UniqueIDPattern is the service account identifier pattern.
	UniqueIDPattern = regexp.MustCompile(`^[a-z][-a-z0-9]{4,28}[a-z0-9]$`)
	// EmailPattern is the service account identifier pattern.
	EmailPattern = regexp.MustCompile(`^[a-z][-a-z0-9]{4,28}[a-z0-9]@[a-z][-a-z0-9]{3,48}[a-z0-9]\.vserviceaccount\.com$`)
	// NamePattern is the service account name pattern.
	NamePattern = regexp.MustCompile(`^projects/[a-z][-a-z0-9]{3,48}[a-z0-9]/serviceAccounts/[a-z][-a-z0-9]{4,28}[a-z0-9]@[a-z][-a-z0-9]{3,48}[a-z0-9]\.vserviceaccount\.com$`)
)

// Name returns the service account's name given a project identifier and a
// unique identifier.
func Name(projID string, accEmail string) (string, error) {
	projName, err := project.Name(projID)
	if err != nil {
		return "", err
	}
	return cstr.JoinWithSeparator(resource.NameSeparator, projName, CollectionID, accEmail), nil
}

// Email returns the service account's email given a project identifier and a
// identifier.
func Email(projID string, accID string) string {
	return cstr.Join(accID, "@", projID, ".", Domain)
}

// ProjIDAndRefFromName returns the service account's project identifier and the
// account reference which is the unique identifier within the project or the
// email.
func ProjIDAndRefFromName(name string) (string, string) {
	nameParts := strings.SplitN(name, resource.NameSeparator, 4)
	return nameParts[1], nameParts[3]
}

// RefFromName returns the service account's reference which is the unique
// identifier within the project or the email.
func RefFromName(name string) string {
	return strings.SplitN(name, resource.NameSeparator, 4)[3]
}

// IsValidUniqueID reports whether a service account unique identifier is valid.
func IsValidUniqueID(ID string) bool {
	return UniqueIDPattern.MatchString(ID)
}

// IsValidEmail reports whether a service account email is valid.
func IsValidEmail(name string) bool {
	return NamePattern.MatchString(name)
}

// IsValidName reports whether a service account name is valid.
func IsValidName(name string) bool {
	return NamePattern.MatchString(name)
}
