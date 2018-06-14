package datetime

import (
	"testing"
	"time"
)

func TestFormatToStdLayout(t *testing.T) {
	tm := time.Now()
	dm := createDatatime(&tm)

	assertTrue(t, dm.formatToStdLayout("") == "", "Format Blank")
	assertTrue(t, dm.formatToStdLayout("%Y") == "2006", "Format Y")
	assertTrue(t, dm.formatToStdLayout("%y") == "06", "Format y")
}

func TestToLayout(t *testing.T) {
	tm := time.Date(2018, time.February, 4, 20, 10, 5, 9, time.UTC)
	dm := createDatatime(&tm)
	assertTrue(t, dm.ToLayout("2006-01-02T15:04:05") == "2018-02-04T20:10:05", "Format: 2006-01-02T15:04:05")
}
func TestFormat(t *testing.T) {
	tm := time.Date(2018, time.February, 4, 20, 7, 5, 9, time.UTC)
	dm := createDatatime(&tm)
	assertTrue(t, dm.Format("%Y-%m-%d %I:%M:%S") == "2018-02-04 08:07:05", "Test: Zero-Padded")
	assertTrue(t, dm.Format("%Y-%-m-%-d %-I:%-M:%-S") == "2018-2-4 8:7:5", "Test: Non-Zero-Padded")
	assertTrue(t, dm.Format("%b %B") == "Feb February", "Test: Month Name")
	assertTrue(t, dm.Format("%p %P") == "pm PM", "Test: Meridian")
	assertTrue(t, dm.Format("%a %A") == "Sun Sunday", "Test: Weekday")
}
