package serviceaccount

import (
	"strings"

	cstrings "github.com/videocoin/common/strings"

	"github.com/videocoin/common/resource"
)

// Name returns the service account's name given a project identifier and a
// unique identifier.
func Name(projID string, accUniqueID string) string {
	return cstrings.Join("projects/", projID, "/serviceAccounts/", accUniqueID)
}

// Email returns the service account's email given a project identifier and a
// unique identifier.
func Email(projID string, accUniqueID string) string {
	return cstrings.Join(accUniqueID, "@", projID, ".vserviceaccount.com")
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
