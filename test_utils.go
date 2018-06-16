package datetime

import (
	"testing"
)

func assertTrue(t *testing.T, result bool, message string) {
	if result == true {
		return
	}

	t.Error(message)
}
