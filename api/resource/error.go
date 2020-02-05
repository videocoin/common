package resource

import (
	"fmt"

	"github.com/videocoin/common/strings"
)

// PatternError returns an error that indicates an invalid pattern.
func PatternError(exp string) error {
	return fmt.Errorf(strings.Join("value does not match regex pattern \"", exp, "\""))
}
