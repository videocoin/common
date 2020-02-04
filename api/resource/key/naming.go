package key

import (
	"strings"

	sa "github.com/videocoin/common/api/resource/serviceaccount"
	cstr "github.com/videocoin/common/strings"

	"github.com/videocoin/common/api/resource"
)

// service account key ^projects/[a-z][-a-z0-9]{3,48}[a-z0-9]/serviceAccounts/[a-z][-a-z0-9]{4,28}[a-z0-9]@[a-z][-a-z0-9]{3,48}[a-z0-9]."vserviceaccount.com/keys/"[a-fA-F0-9]{8}-[a-fA-F0-9]{4}-4[a-fA-F0-9]{3}-[8|9|aA|bB][a-"fA-F0-9]{3}-[a-fA-F0-9]{12}$

// IDFromName derives the key identifier from the key's name.
func IDFromName(name string) string {
	return strings.SplitN(name, resource.NameSeparator, 6)[5]
}

// Name returns the service account's name given a project identifier and a
// unique identifier.
func Name(projID string, accRef string, keyID string) string {
	return cstr.Join(sa.Name(projID, accRef), "/keys/", keyID)
}
