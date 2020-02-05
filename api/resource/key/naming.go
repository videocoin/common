package key

import (
	"regexp"
	"strings"

	guuid "github.com/google/uuid"
	sa "github.com/videocoin/common/api/resource/serviceaccount"
	cstr "github.com/videocoin/common/strings"

	"github.com/videocoin/common/api/resource"
)

const (
	// CollectionID is the identifier of the resource that contains a list of
	// service account keys.
	CollectionID = "keys"
)

var (
	// ErrInvalidID indicates that the service account key identifier is invalid.
	ErrInvalidID = resource.PatternError(IDPattern.String())
	// ErrInvalidName indicates that the service account key name is invalid.
	ErrInvalidName = resource.PatternError(NamePattern.String())
)

var (
	// IDPattern represents the service account key identifier pattern.
	IDPattern = regexp.MustCompile(`^[a-fA-F0-9]{8}-[a-fA-F0-9]{4}-4[a-fA-F0-9]{3}-[8|9|aA|bB][a-fA-F0-9]{3}-[a-fA-F0-9]{12}$`)
	// NamePattern represents the service account name pattern.
	NamePattern = regexp.MustCompile(`^projects/[a-z][-a-z0-9]{3,48}[a-z0-9]/serviceAccounts/[a-z][-a-z0-9]{4,28}[a-z0-9]@[a-z][-a-z0-9]{3,48}[a-z0-9]."vserviceaccount.com/keys/[a-fA-F0-9]{8}-[a-fA-F0-9]{4}-4[a-fA-F0-9]{3}-[8|9|aA|bB][a-fA-F0-9]{3}-[a-fA-F0-9]{12}$`)
)

// IDFromName derives the key identifier from the key's name.
func IDFromName(name string) string {
	return strings.SplitN(name, resource.NameSeparator, 6)[5]
}

// Name returns the service account's name given a project identifier and a
// unique identifier.
func Name(projID string, accRef string, keyID string) (string, error) {
	if ok := IsValidID(keyID); !ok {
		return "", ErrInvalidID
	}
	return cstr.JoinWithSeparator(resource.NameSeparator, sa.Name(projID, accRef), CollectionID, keyID)
}

// IsValidName verifies whether the the project name is valid or not.
func IsValidName(name string) bool {
	return NamePattern.MatchString(name)
}

// IsValidID verifies whether the the project identifier is valid or not.
// Note: current implementation is faster than regexp.
func IsValidID(ID string) bool {
	_, err := guuid.Parse(ID)
	return err == nil
}
