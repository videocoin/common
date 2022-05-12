package types

// BadRequestError is an internal type (TBD...)
type BadRequestError struct {
	Message string `json:"message,required"`
}

// GetMessage is an internal getter (TBD...)
func (v *BadRequestError) GetMessage() (o string) {
	if v != nil {
		return v.Message
	}
	return
}
