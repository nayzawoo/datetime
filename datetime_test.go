package datetime

import (
	"testing"
	"time"
)

func TestNewFromFormatWithDate(t *testing.T) {
	dt, _ := NewFromFormat("Date: {YY} {M} {D}", "Date: 18 2 3", time.UTC)
	assertTrue(t, dt.Format("{YYYY} {MM} {DD}") == "2018 02 03", "date should 2018 02 03")

	dt, _ = NewFromFormat("Date: {YYYY} {MM} {DD}", "Date: 2018 12 10", time.UTC)
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

	dt = NewFromDate(2000, 2, 1, time.UTC)
	assertTrue(t, dt.DaysInMonth() == 29, "TestDaysInMonth for Leap Feb")

	dt = NewFromDate(2001, 2, 1, time.UTC)
	assertTrue(t, dt.DaysInMonth() == 28, "TestDaysInMonth")

	dt = NewFromDate(2001, 4, 1, time.UTC)
	assertTrue(t, dt.DaysInMonth() == 30, "TestDaysInMonth")
}

func TestStartOfs(t *testing.T) {
	tTime := time.Date(2016, 2, 3, 10, 20, 30, 40, time.UTC)

	dt := New(tTime)
	dt.StartOfMinute()

	// minute
	assertDate(t, dt, "2016-02-03 10:20:00")
	assertTrue(t, dt.Nanosecond() == 0, "")

	// hour
	dt.t = tTime
	dt.StartOfHour()
	assertTrue(t, dt.DateTimeString() == "2016-02-03 10:00:00", "start of hour")
	assertTrue(t, dt.Nanosecond() == 0, "")

	// day
	dt.t = tTime
	dt.StartOfDay()
	assertTrue(t, dt.DateTimeString() == "2016-02-03 00:00:00", "start of day")
	assertTrue(t, dt.Nanosecond() == 0, "")

	// month
	dt.t = tTime
	dt.StartOfMonth()
	assertTrue(t, dt.DateTimeString() == "2016-02-01 00:00:00", "start of month")
	assertTrue(t, dt.Nanosecond() == 0, "")

	// year
	dt.t = tTime
	dt.StartOfYear()
	assertTrue(t, dt.DateTimeString() == "2016-01-01 00:00:00", "start of month")
	assertTrue(t, dt.Nanosecond() == 0, "")
}

func TestEndOfs(t *testing.T) {
	tTime := time.Date(2016, 2, 3, 10, 20, 30, 40, time.UTC)
	dt := New(tTime)

	dt.EndOfMinute()
	assertDate(t, dt, "2016-02-03 10:20:59")
	assertTrue(t, dt.Nanosecond() == 999999999, "")

	// hour
	dt.t = tTime
	dt.EndOfHour()
	assertDate(t, dt, "2016-02-03 10:59:59")
	assertTrue(t, dt.Nanosecond() == 999999999, "")

	// day
	dt.t = tTime
	dt.EndOfDay()
	assertDate(t, dt, "2016-02-03 23:59:59")
	assertTrue(t, dt.Nanosecond() == 999999999, "")

	// month
	dt = NewFromDate(2017, 2, 1, time.UTC)
	dt.EndOfMonth()
	assertDate(t, dt, "2017-02-28 23:59:59")
	assertTrue(t, dt.Nanosecond() == 999999999, "")

	dt = NewFromDate(2000, 2, 1, time.UTC)
	dt.EndOfMonth()
	assertDate(t, dt, "2000-02-29 23:59:59")
	assertTrue(t, dt.Nanosecond() == 999999999, "")

	dt.t = dt.Time().Add(time.Nanosecond)
	assertDate(t, dt, "2000-03-01 00:00:00")

	// year
	dt = NewFromDate(2017, 2, 1, time.UTC)
	dt.EndOfYear()
	assertDate(t, dt, "2017-12-31 23:59:59")
	assertTrue(t, dt.Nanosecond() == 999999999, "")

}
