package datetime

import (
	"testing"
	"time"
)

func TestNow(t *testing.T) {
	dt := Now(nil)
	assertDateTime(t, dt, time.Now().Format("2006-01-02 15:04:05"))

	SetTestNow(time.Now())
	SetTestNow(time.Date(1900, 1, 1, 0, 0, 0, 0, time.UTC))
	dt = Now(nil)
	assertTrue(t, HasTestNow(), "")
	assertDateTime(t, dt, "1900-01-01 00:00:00")

	ResetTestNow()
	assertTrue(t, !HasTestNow(), "")
	dt = Now(nil)
	assertDateTime(t, dt, time.Now().Format("2006-01-02 15:04:05"))

	newYork, _ := time.LoadLocation("America/New_York")

	local := Now(nil).StartOfMinute()
	utc := Now(time.UTC).StartOfMinute()
	newYorkTime := Now(newYork).StartOfMinute()
	assertTrue(t, local.Eq(utc), "")
	assertTrue(t, local.Eq(utc), "")
	assertTrue(t, utc.Eq(newYorkTime), "")
}

func TestCopy(t *testing.T) {
	dt := NewFromDate(2018, 12, 10, nil)

	dt2 := dt.Copy()

	// Modify
	dt.StartOfMonth()
	dt2.AddDate(1, 0, 0)

	assertDateTime(t, dt, "2018-12-01 00:00:00")
	assertDateTime(t, dt2, "2019-12-10 00:00:00")
}

func TestNewFromFormatWithDate(t *testing.T) {
	dt, _ := NewFromFormat("{YY}-{M}-{D} {H}:{m}:{s}", "18-2-3 10:20:30", time.UTC)
	assertDateTime(t, dt, "2018-02-03 10:20:30")
}

func TestNewFromDate(t *testing.T) {
	dt := NewFromDate(2018, 12, 10, nil)

	assertDateTime(t, dt, "2018-12-10 00:00:00")
	assertTrue(t, dt.Format("{YYYY} {MM} {DD}") == "2018 12 10", "date should 2018 12 10")
}

func TestNewFromFormatWithTime(t *testing.T) {
	dt, err := NewFromFormat("Date: {h}:{m}:{s} {pm}", "Date: 1:2:3 pm", time.UTC)
	if err != nil {
		t.Error(err)
	}

	assertTrue(t, dt.Format("{HH}:{mm}:{ss}") == "13:02:03", "test time 13:02:03")
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
	assertDateTime(t, dt, "2016-02-03 10:20:00")
	assertTrue(t, dt.Nanosecond() == 0, "")

	// hour
	dt.t = tTime
	dt.StartOfHour()
	assertDateTime(t, dt, "2016-02-03 10:00:00")
	assertTrue(t, dt.Nanosecond() == 0, "")

	// day
	dt.t = tTime
	dt.StartOfDay()
	assertDateTime(t, dt, "2016-02-03 00:00:00")
	assertTrue(t, dt.Nanosecond() == 0, "")

	// month
	dt.t = tTime
	dt.StartOfMonth()
	assertDateTime(t, dt, "2016-02-01 00:00:00")
	assertTrue(t, dt.Nanosecond() == 0, "")

	// year
	dt.t = tTime
	dt.StartOfYear()
	assertDateTime(t, dt, "2016-01-01 00:00:00")
	assertTrue(t, dt.Nanosecond() == 0, "")
}

func TestEndOfs(t *testing.T) {
	tTime := time.Date(2016, 2, 3, 10, 20, 30, 40, time.UTC)
	dt := New(tTime)

	dt.EndOfMinute()
	assertDateTime(t, dt, "2016-02-03 10:20:59")
	assertTrue(t, dt.Nanosecond() == 999999999, "")

	// hour
	dt.t = tTime
	dt.EndOfHour()
	assertDateTime(t, dt, "2016-02-03 10:59:59")
	assertTrue(t, dt.Nanosecond() == 999999999, "")

	// day
	dt.t = tTime
	dt.EndOfDay()
	assertDateTime(t, dt, "2016-02-03 23:59:59")
	assertTrue(t, dt.Nanosecond() == 999999999, "")

	// month
	dt = NewFromDate(2017, 2, 1, time.UTC)
	dt.EndOfMonth()
	assertDateTime(t, dt, "2017-02-28 23:59:59")
	assertTrue(t, dt.Nanosecond() == 999999999, "")

	dt = NewFromDate(2000, 2, 1, time.UTC)
	dt.EndOfMonth()
	assertDateTime(t, dt, "2000-02-29 23:59:59")
	assertTrue(t, dt.Nanosecond() == 999999999, "")

	dt.t = dt.Time().Add(time.Nanosecond)
	assertDateTime(t, dt, "2000-03-01 00:00:00")

	// year
	dt = NewFromDate(2017, 2, 1, time.UTC)
	dt.EndOfYear()
	assertDateTime(t, dt, "2017-12-31 23:59:59")
	assertTrue(t, dt.Nanosecond() == 999999999, "")

}

func TestAddSubTime(t *testing.T) {
	tTime := time.Date(2016, 2, 3, 10, 20, 30, 40, time.UTC)
	dt := New(tTime)

	dt.AddTime(-10, -20, -30)
	assertDateTime(t, dt, "2016-02-03 00:00:00")

	dt.AddTime(10, 20, 30)
	assertDateTime(t, dt, "2016-02-03 10:20:30")

	dt.SubTime(2, 5, 10)
	assertDateTime(t, dt, "2016-02-03 08:15:20")

	dt.Add(time.Second * 10)
	assertDateTime(t, dt, "2016-02-03 08:15:30")
}

func TestAddSubDate(t *testing.T) {
	tTime := time.Date(2016, 2, 3, 10, 20, 30, 40, time.UTC)
	dt := New(tTime)

	dt.AddDate(2, 1, 2)
	assertDateTime(t, dt, "2018-03-05 10:20:30")

	dt.SubDate(2, 1, 2)
	assertDateTime(t, dt, "2016-02-03 10:20:30")
}

func TestEqual(t *testing.T) {
	dt := NewFromDate(2017, 2, 1, time.UTC)

	// New York => -05:00
	newYork, _ := time.LoadLocation("America/New_York")
	dt2 := NewFromDate(2017, 2, 1, newYork)

	assertTrue(t, !dt.Eq(dt2), "")

	dt.AddTime(5, 0, 0)
	assertTrue(t, dt.Eq(dt2), "")
}
