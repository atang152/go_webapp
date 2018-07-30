// HELPER FUNCTION - Using bytes.Buffer for efficient string concatenation in Go
package helper

import (
	"bytes"
)

func Concat(values []string) string {

	var b bytes.Buffer
	for _, s := range values {
		b.WriteString(s)
	}
	return b.String()
}
