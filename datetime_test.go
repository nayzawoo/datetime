package datetime

import (
	"testing"
	"time"
)

func TestTime(t *testing.T) {
	in := time.Now()
	dt := createDatatime(&in)
	out := dt.Time()

	if !out.Equal(in) {
		t.Errorf("output time should = given time")
	}
}
