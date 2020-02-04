package key

import (
	"strings"

	cstrings "github.com/videocoin/common/strings"

	"github.com/videocoin/common/resource"
)

// IDFromName derives the key identifier from the key's name.
func IDFromName(name string) string {
	return strings.SplitN(name, resource.NameSeparator, 6)[5]
}

// Name returns the service account's name given a project identifier and a
// unique identifier.
func Name(projID string, accUniqueID string, keyID string) string {
	return cstrings.Join("projects/", projID, "/serviceAccounts/", accUniqueID, "/keys/", keyID)
}
