package naming_test

import (
	"testing"
)

func TestName(t *testing.T) {
	tests := []struct {
		projID, accEmail, keyID string
		output                  string
	}{
		{
			projID:   "test",
			accEmail: "email",
			keyID:    "id",
			output:   "",
		},
	}
}

func IsValidID(t *testing.T) {
	tests := []struct {
		keyID  string
		output string
	}{
		{
			keyID:  "id",
			output: "",
		},
	}
}
