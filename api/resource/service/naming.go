package service

import (
	"regexp"

	cexp "github.com/videocoin/common/regexp"

	"github.com/videocoin/common/api/resource"
	"github.com/videocoin/common/strings"
)

// NamePattern represents the service name pattern.
const NamePattern = "[a-z]{3,20}.videocoin.network"

// ErrInvalidName indicates that the project name is invalid.
var ErrInvalidName = resource.PatternError(NamePattern)

var nameRegExp = regexp.MustCompile(strings.Join(cexp.Begin, NamePattern, cexp.End))

// IsValidName verifies whether the the project name is valid or not.
func IsValidName(name string) bool {
	return nameRegExp.MatchString(name)
}
