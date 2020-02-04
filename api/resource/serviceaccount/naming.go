package serviceaccount

import (
	"strings"

	cstrings "github.com/videocoin/common/strings"

	"github.com/videocoin/common/api/resource"
	"github.com/videocoin/common/api/resource/project"
)

// account id ^[a-z][-a-z0-9]{4,28}[a-z0-9]$
// service account ^projects/[a-z][-a-z0-9]{3,48}[a-z0-9]/serviceAccounts/"[a-z][-a-z0-9]{4,28}[a-z0-9]@[a-z][-a-z0-9]{3,48}[a-z0-9]"\.vserviceaccount\.com$

// Name returns the service account's name given a project identifier and a
// unique identifier.
func Name(projID string, accRef string) string {
	return cstrings.Join(project.Name(projID), resource.NameSeparator, "serviceAccounts", resource.NameSeparator, accRef)
}

// Email returns the service account's email given a project identifier and a
// identifier.
func Email(projID string, accID string) string {
	return cstrings.Join(accID, "@", projID, ".vserviceaccount.com")
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
