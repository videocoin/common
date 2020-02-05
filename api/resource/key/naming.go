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
	// IDPattern is the key identifier pattern.
	IDPattern = regexp.MustCompile(`^[a-fA-F0-9]{8}-[a-fA-F0-9]{4}-4[a-fA-F0-9]{3}-[8|9|aA|bB][a-fA-F0-9]{3}-[a-fA-F0-9]{12}$`)
	// NamePattern is the key name pattern.
	NamePattern = regexp.MustCompile(`^projects/[a-z][-a-z0-9]{3,48}[a-z0-9]/serviceAccounts/[a-z][-a-z0-9]{4,28}[a-z0-9]@[a-z][-a-z0-9]{3,48}[a-z0-9]\.vserviceaccount\.com/keys/[a-fA-F0-9]{8}-[a-fA-F0-9]{4}-4[a-fA-F0-9]{3}-[8|9|aA|bB][a-fA-F0-9]{3}-[a-fA-F0-9]{12}$`)
)

// IDFromName derives the key's identifier from its name.
func IDFromName(name string) (string, error) {
	if ok := IsValidName(name); !ok {
		return "", ErrInvalidName
	}
	return strings.SplitN(name, resource.NameSeparator, 6)[5], nil
}

// Name returns the key's name given a project identifier and a
// service account email.
func Name(projID string, accEmail string, keyID string) (string, error) {
	saName, err := sa.Name(projID, accEmail)
	if err != nil {
		return "", err
	}
	if ok := IsValidID(keyID); !ok {
		return "", ErrInvalidID
	}
	return cstr.JoinWithSeparator(resource.NameSeparator, saName, CollectionID, keyID), nil
}

// IsValidName reports whether a key name is valid.
func IsValidName(name string) bool {
	return NamePattern.MatchString(name)
}

// IsValidID reports whether a key identifier is valid.
// Note: current implementation is faster than regexp.
func IsValidID(ID string) bool {
	_, err := guuid.Parse(ID)
	return err == nil
}
