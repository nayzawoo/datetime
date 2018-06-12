package datetime_test

import (
	"github.com/nayzawoo/datetime"
	"testing"
	"time"
)

func createDatatime(tm *time.Time) *datetime.Datetime {
	if tm == nil {
		now := time.Now()
		tm = &now
	}

	dt := datetime.New(tm)

	return dt
}

func TestTime(t *testing.T) {
	in := time.Now()
	dt := createDatatime(&in)
	out := dt.Time()

	if !out.Equal(in) {
		t.Errorf("output time should = given time")
	}
}
