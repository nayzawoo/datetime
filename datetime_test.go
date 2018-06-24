package datetime

import (
	"testing"
	"time"
)

func TestNewFromFormatWithDate(t *testing.T) {
	// 2. short y,m,d
	dt, err := NewFromFormat("Date: {YY} {M} {D}", "Date: 18 2 3", time.UTC)

	if err != nil {
		t.Error(err)
	}

	assertTrue(t, dt.Format("{YYYY} {MM} {DD}") == "2018 02 03", "date should 2018 02 03")

	// 1. long y,m,d
	dt, err = NewFromFormat("Date: {YYYY} {MM} {DD}", "Date: 2018 12 10", time.UTC)

	if err != nil {
		t.Error(err)
	}

	assertTrue(t, dt.Format("{YYYY} {MM} {DD}") == "2018 12 10", "date should 2018 12 10")

	assertTrue(t, dt.Format("{YYYY} {MM} {DD}") == "2018 12 10", "date should 2018 12 10")
}

func TestNewFromDate(t *testing.T) {
	dt := NewFromDate(2018, 12, 10, nil)

	assertTrue(t, dt.Format("{YYYY} {MM} {DD}") == "2018 12 10", "date should 2018 12 10")
}

func TestNewFromFormatWithTime(t *testing.T) {
	dt, err := NewFromFormat("Date: {h}:{m}:{s} {pm}", "Date: 1:2:3 pm", time.UTC)
	if err != nil {
		t.Error(err)
	}

	assertTrue(t, dt.Format("{HH}:{mm}:{ss}") == "13:02:03", "test time 13:02:03")
}

func TestSetDateTime(t *testing.T) {
	tTime := time.Date(2016, 1, 2, 10, 20, 30, 40, time.UTC)

	dt := New(tTime)

	assertTrue(t, dt.Nanosecond() == 40, "setdatetime test: ns")
	dt.set(41, "nanosecond")
	dt.set(31, "second")
	dt.set(21, "minute")
	dt.set(11, "hour")
	dt.set(3, "day")
	dt.set(2, "month")
	dt.set(2017, "year")

	assertTrue(t, dt.Format("{YYYY}-{M}-{D} {H}:{m}:{s}") == "2017-2-3 11:21:31", "test setDateTime")
	assertTrue(t, dt.Nanosecond() == 41, "setdatetime test: ns")

	dt.set(40, "nsec")
	dt.set(30, "sec")
	dt.set(20, "min")
	dt.set(10, "hr")
	assertTrue(t, dt.Format("{H}:{m}:{s}") == "10:20:30", "test setDateTime")
	assertTrue(t, dt.Nanosecond() == 40, "setdatetime test: ns")
}

func TestDaysInMonth(t *testing.T) {
	dt := NewFromDate(2000, 1, 1, time.UTC)
	assertTrue(t, dt.DaysInMonth() == 31, "TestDaysInMonth")

	//
	dt = NewFromDate(2000, 2, 1, time.UTC)
	assertTrue(t, dt.DaysInMonth() == 29, "TestDaysInMonth for Leap Feb")

	dt = NewFromDate(2001, 2, 1, time.UTC)
	assertTrue(t, dt.DaysInMonth() == 28, "TestDaysInMonth")

	dt = NewFromDate(2001, 4, 1, time.UTC)
	assertTrue(t, dt.DaysInMonth() == 30, "TestDaysInMonth")
}

func TestStartOfs(t *testing.T) {
	tTime := time.Date(2016, 1, 2, 10, 20, 30, 40, time.UTC)

	dt := New(tTime)
	dt.StartOfMinute()

	// Minute
	assertTrue(t, dt.Format("{h}:{m}:{s}") == "10:20:0", "start of minute")
	assertTrue(t, dt.Nanosecond() == 0, "start of minute")

	// Hour
	dt.t = tTime
	dt.StartOfHour()
	assertTrue(t, dt.Format("{D} {H}:{m}:{s}") == "2 10:0:0", "start of hour")
	assertTrue(t, dt.Nanosecond() == 0, "")

	// Day
	dt.t = tTime
	dt.StartOfDay()
	assertTrue(t, dt.Format("{D} {H}:{m}:{s}") == "2 0:0:0", "start of day")
	assertTrue(t, dt.Nanosecond() == 0, "")

	// Month
	dt.t = tTime
	dt.StartOfMonth()
	assertTrue(t, dt.DateTimeString() == "2016-01-01 00:00:00", "start of month")
	assertTrue(t, dt.Nanosecond() == 0, "")
}

func TestEndOfs(t *testing.T) {
	tTime := time.Date(2016, 1, 2, 10, 20, 30, 40, time.UTC)
	dt := New(tTime)

	// minute
	dt.EndOfMinute()
	assertTrue(t, dt.Format("{h}:{m}:{s}") == "10:20:59", "EndOfTest: min")
	assertTrue(t, dt.Nanosecond() == 999999999, "EndOfTest: min")

	// hour
	dt.t = tTime
	dt.EndOfHour()
	assertTrue(t, dt.Format("{h}:{m}:{s}") == "10:59:59", "EndOfTest: hr")
	assertTrue(t, dt.Nanosecond() == 999999999, "EndOfTest: hr")

	// day
	dt.t = tTime
	dt.EndOfDay()
	assertTrue(t, dt.Format("{H}:{m}:{s}") == "23:59:59", "EndOfTest: day")
	assertTrue(t, dt.Nanosecond() == 999999999, "EndOfTest: day")
	dt.t = tTime

	// test month
	dt.t = time.Date(2000, 2, 1, 10, 20, 30, 40, time.UTC)
	dt.EndOfMonth()
	assertTrue(t, dt.DateTimeString() == "2000-02-29 23:59:59", "EndOfTest: Month")
	dt.t = dt.Time().Add(time.Nanosecond)
	assertTrue(t, dt.DateTimeString() == "2000-03-01 00:00:00", "EndOfTest: Month")
	dt.t = tTime
}
