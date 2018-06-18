package datetime

import (
	"testing"
	"time"
)

func TestYearToNanoSecond(t *testing.T) {
	tm := time.Date(2018, time.February, 4, 20, 10, 5, 9, time.UTC)
	dm := New(tm)
	assertTrue(t, dm.Year() == 2018, "Year => 2018")
	assertTrue(t, dm.Month() == 2, "Month => 2(February)")
	assertTrue(t, dm.Day() == 4, "Day => 4")
	assertTrue(t, dm.Hour() == 20, "Hour => 20")
	assertTrue(t, dm.Minute() == 10, "Minute => 10")
	assertTrue(t, dm.Second() == 5, "Second => 5")
	assertTrue(t, dm.Nanosecond() == 9, "Nanosecond => 9")
	assertTrue(t, dm.Weekday() == 0, "Weekday => 1(Sun)")
	assertTrue(t, dm.YearDay() == 35, "DayOfYear => 35")
}
