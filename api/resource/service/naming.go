package service

import (
	"regexp"

	"github.com/videocoin/common/api/resource"
)

// ErrInvalidName indicates that the service name is invalid.
var ErrInvalidName = resource.PatternError(NamePattern.String())

// NamePattern represents the service name pattern.
var NamePattern = regexp.MustCompile(`^[a-z]{3,20}.videocoin.network$`)

// IsValidName verifies whether the the service name is valid or not.
func IsValidName(name string) bool {
	return NamePattern.MatchString(name)
}
