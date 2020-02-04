package strings

import "strings"

// Join concatenates a variable number of strings.
func Join(strs ...string) string {
	var sb strings.Builder
	for _, str := range strs {
		sb.WriteString(str)
	}
	return sb.String()
}
