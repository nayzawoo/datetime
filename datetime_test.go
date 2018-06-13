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

func TestFormatToStdLayout(t *testing.T) {
	assertTrue(t, FormatToStdLayout("") == "", "Format Blank")
	assertTrue(t, FormatToStdLayout("Y") == "2006", "Format YYYY")
	assertTrue(t, FormatToStdLayout("y") == "06", "Format YY")
}
