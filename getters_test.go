package datetime

import (
	"testing"
	"time"
)

func TestYearToNanoSecond(t *testing.T) {
	tm := time.Date(1994, time.May, 7, 20, 10, 5, 9, time.UTC)
	dm := createDatatime(&tm)
	assertTrue(t, dm.Year() == 1994, "Year => 1994")
	assertTrue(t, dm.Month() == 5, "Month => 5(may)")
	assertTrue(t, dm.Day() == 7, "Day => 7")
	assertTrue(t, dm.Hour() == 20, "Hour => 20")
	assertTrue(t, dm.Minute() == 10, "Minute => 10")
	assertTrue(t, dm.Second() == 5, "Second => 5")
	assertTrue(t, dm.Nanosecond() == 9, "Nanosecond => 9")
}
