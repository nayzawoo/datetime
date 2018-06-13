package datetime

import (
	"testing"
	"time"
)

func TestFormatToStdLayout(t *testing.T) {
	assertTrue(t, FormatToStdLayout("") == "", "Format Blank")
	assertTrue(t, FormatToStdLayout("Y") == "2006", "Format YYYY")
	assertTrue(t, FormatToStdLayout("y") == "06", "Format YY")
}

func TestToLayout(t *testing.T) {
	tm := time.Date(2018, time.February, 4, 20, 10, 5, 9, time.UTC)
	dm := createDatatime(&tm)
	assertTrue(t, dm.ToLayout("2006-01-02T15:04:05") == "2018-02-04T20:10:05", "Format: 2006-01-02T15:04:05")
}
func TestFormat(t *testing.T) {
	tm := time.Date(2018, time.February, 4, 20, 10, 5, 9, time.UTC)
	dm := createDatatime(&tm)
	t.Log(dm.Format("F j, Y, g:i a"))
	assertTrue(t, dm.Format("F j, Y, g:i a") == "February 4, 2018, 8:10 pm", "Format: F j, Y, g:i a")
	// TODO
}
