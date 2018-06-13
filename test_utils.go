package datetime

import (
	"testing"
	"time"
)

func createDatatime(tm *time.Time) *Datetime {
	if tm == nil {
		now := time.Now()
		tm = &now
	}

	dt := New(tm)

	return dt
}

func assertTrue(t *testing.T, result bool, message string) {
	if result == true {
		return
	}

	t.Error(message)
}
