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

func TestNewFromFormatWithDate(t *testing.T) {
	dt, _ := NewFromFormat("{YY}-{M}-{D} {H}:{m}:{s}", "18-2-3 10:20:30", time.UTC)
	assertDateTime(t, dt, "2018-02-03 10:20:30")

	dt, _ = NewFromFormat("{YY}-{M}-{D} {H}:{m}:{s}", "18-2-3 10:20:30", nil)
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

func TestUnix(t *testing.T) {
	dt := Unix(3600, 0)

	assertDateTime(t, dt.In(time.UTC), "1970-01-01 01:00:00")
	assertDateTime(t, dt.UTC(), "1970-01-01 01:00:00")
}

func TestStartOfs(t *testing.T) {
	tTime := time.Date(2016, 2, 3, 10, 20, 30, 40, time.UTC)

	dt := New(tTime)

	// minute
	dt = dt.StartOfMinute()
	assertDateTime(t, dt, "2016-02-03 10:20:00")
	assertTrue(t, dt.Nanosecond() == 0, "")

	// hour
	dt = dt.StartOfHour()
	assertDateTime(t, dt, "2016-02-03 10:00:00")
	assertTrue(t, dt.Nanosecond() == 0, "")

	// day
	dt = dt.StartOfDay()
	assertDateTime(t, dt, "2016-02-03 00:00:00")
	assertTrue(t, dt.Nanosecond() == 0, "")

	// month
	dt = dt.StartOfMonth()
	assertDateTime(t, dt, "2016-02-01 00:00:00")
	assertTrue(t, dt.Nanosecond() == 0, "")

	// year
	dt = dt.StartOfYear()
	assertDateTime(t, dt, "2016-01-01 00:00:00")
	assertTrue(t, dt.Nanosecond() == 0, "")
}

func TestEndOfs(t *testing.T) {
	tTime := time.Date(2016, 2, 3, 10, 20, 30, 40, time.UTC)
	dt := New(tTime)

	dt = dt.EndOfMinute()
	assertDateTime(t, dt, "2016-02-03 10:20:59")
	assertTrue(t, dt.Nanosecond() == 999999999, "")

	// hour
	dt = dt.EndOfHour()
	assertDateTime(t, dt, "2016-02-03 10:59:59")
	assertTrue(t, dt.Nanosecond() == 999999999, "")

	// day
	dt = dt.EndOfDay()
	assertDateTime(t, dt, "2016-02-03 23:59:59")
	assertTrue(t, dt.Nanosecond() == 999999999, "")

	// month
	dt = NewFromDate(2017, 2, 1, time.UTC)
	dt = dt.EndOfMonth()
	assertDateTime(t, dt, "2017-02-28 23:59:59")
	assertTrue(t, dt.Nanosecond() == 999999999, "")

	dt = NewFromDate(2000, 2, 1, time.UTC)
	dt = dt.EndOfMonth()
	assertDateTime(t, dt, "2000-02-29 23:59:59")
	assertTrue(t, dt.Nanosecond() == 999999999, "")

	dt.t = dt.Time().Add(time.Nanosecond)
	assertDateTime(t, dt, "2000-03-01 00:00:00")

	// year
	dt = NewFromDate(2017, 2, 1, time.UTC)
	dt = dt.EndOfYear()
	assertDateTime(t, dt, "2017-12-31 23:59:59")
	assertTrue(t, dt.Nanosecond() == 999999999, "")

}

func TestAddSubTime(t *testing.T) {
	tTime := time.Date(2016, 2, 3, 10, 20, 30, 40, time.UTC)
	dt := New(tTime)

	dt = dt.AddTime(-10, -20, -30)
	assertDateTime(t, dt, "2016-02-03 00:00:00")

	dt = dt.AddTime(10, 20, 30)
	assertDateTime(t, dt, "2016-02-03 10:20:30")

	dt = dt.SubTime(2, 5, 10)
	assertDateTime(t, dt, "2016-02-03 08:15:20")

	dt = dt.Add(time.Second * 10)
	assertDateTime(t, dt, "2016-02-03 08:15:30")

	dt = dt.AddSeconds(10)
	assertDateTime(t, dt, "2016-02-03 08:15:40")

	dt = dt.AddMinutes(5)
	assertDateTime(t, dt, "2016-02-03 08:20:40")

	dt = dt.AddHours(2)
	assertDateTime(t, dt, "2016-02-03 10:20:40")

	dt = dt.AddDays(1)
	assertDateTime(t, dt, "2016-02-04 10:20:40")

	dt = dt.AddMonths(1)
	assertDateTime(t, dt, "2016-03-04 10:20:40")

	dt = dt.AddYears(1)
	assertDateTime(t, dt, "2017-03-04 10:20:40")

	dt = dt.SubSeconds(10)
	assertDateTime(t, dt, "2017-03-04 10:20:30")

	dt = dt.SubMinutes(10)
	assertDateTime(t, dt, "2017-03-04 10:10:30")

	dt = dt.SubHours(10)
	assertDateTime(t, dt, "2017-03-04 00:10:30")

	dt = dt.SubDays(1)
	assertDateTime(t, dt, "2017-03-03 00:10:30")

	dt = dt.SubMonths(1)
	assertDateTime(t, dt, "2017-02-03 00:10:30")

	dt = dt.SubYears(1)
	assertDateTime(t, dt, "2016-02-03 00:10:30")

	dt = dt.Sub(time.Second * 20)
	assertDateTime(t, dt, "2016-02-03 00:10:10")

}

func TestAddSubDate(t *testing.T) {
	tTime := time.Date(2016, 2, 3, 10, 20, 30, 40, time.UTC)
	dt := New(tTime)

	dt = dt.AddDate(2, 1, 2)
	assertDateTime(t, dt, "2018-03-05 10:20:30")

	dt = dt.SubDate(2, 1, 2)
	assertDateTime(t, dt, "2016-02-03 10:20:30")
}

func TestComparisons(t *testing.T) {
	dt := NewFromDate(2017, 2, 1, time.UTC)

	// New York => -05:00
	newYork, _ := time.LoadLocation("America/New_York")
	dt2 := NewFromDate(2017, 2, 1, newYork)

	assertTrue(t, !dt.Eq(dt2), "")

	dt = dt.AddTime(5, 0, 0)
	assertTrue(t, dt.Eq(dt2), "")

	// Lt
	y2000 := NewFromDate(2000, 1, 1, time.UTC)
	y2001 := NewFromDate(2001, 1, 1, time.UTC)

	assertTrue(t, y2000.Lt(y2001), "")
	assertTrue(t, y2000.Lte(y2000), "")

	assertTrue(t, y2001.Gt(y2000), "")
	assertTrue(t, y2001.Gte(y2001), "")
}

func TestDiffs(t *testing.T) {
	dt := NewFromDate(2017, 2, 1, time.UTC)

	tenMinLater := NewFromDate(2017, 2, 1, time.UTC).AddMinutes(10)

	assertTrue(t, dt.Diff(tenMinLater).Minutes() == -10, "")
}

func TestIs(t *testing.T) {
	dt := NewFromDate(2017, 2, 1, time.UTC)
	assertTrue(t, !dt.IsLeapYear(), "")

	dt = NewFromDate(2000, 2, 1, time.UTC)
	assertTrue(t, dt.IsLeapYear(), "")
}
