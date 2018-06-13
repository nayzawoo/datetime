package datetime

import (
	"testing"
	"time"
)

func TestYearToNanoSecond(t *testing.T) {
	tm := time.Date(2018, time.February, 4, 20, 10, 5, 9, time.UTC)
	dm := createDatatime(&tm)
	assertTrue(t, dm.Year() == 2018, "Year => 2018")
	assertTrue(t, dm.Month() == 2, "Month => 2(February)")
	assertTrue(t, dm.Day() == 4, "Day => 4")
	assertTrue(t, dm.Hour() == 20, "Hour => 20")
	assertTrue(t, dm.Minute() == 10, "Minute => 10")
	assertTrue(t, dm.Second() == 5, "Second => 5")
	assertTrue(t, dm.Nanosecond() == 9, "Nanosecond => 9")
	assertTrue(t, dm.DayOfWeek() == 0, "Weekday => 1(Sun)")
	assertTrue(t, dm.DayOfYear() == 35, "DayOfYear => 35")
	assertTrue(t, dm.ToLayout("2006-01-02T15:04:05") == "2018-02-04T20:10:05", "Format: 2006-01-02T15:04:05")
}

func TestFormat(t *testing.T) {
	tm := time.Date(2018, time.February, 4, 20, 10, 5, 9, time.UTC)
	dm := createDatatime(&tm)
	t.Log(dm.Format("F j, Y, g:i a"))
	assertTrue(t, dm.Format("F j, Y, g:i a") == "February 4, 2018, 8:10 pm", "Format: F j, Y, g:i a")
	// TODO
}
