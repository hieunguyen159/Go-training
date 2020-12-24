package helpers

import (
	"bytes"
	"fmt"
)

func CreateKeyValuePairs(m map[string]float64) string {
	b := new(bytes.Buffer)
	for key, value := range m {
		fmt.Fprintf(b, "%v=%f\n", key, value)
	}
	return b.String()
}
