package types

import "fmt"

func (err BadRequestError) Error() string {
	return fmt.Sprintf("BadRequestError{Message: %v}", err.Message)
}
