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
	dt.setDateTime(41, "nanosecond")
	dt.setDateTime(31, "second")
	dt.setDateTime(21, "minute")
	dt.setDateTime(11, "hour")
	dt.setDateTime(3, "day")
	dt.setDateTime(2, "month")
	dt.setDateTime(2017, "year")

	assertTrue(t, dt.Format("{YYYY}-{M}-{D} {H}:{m}:{s}") == "2017-2-3 11:21:31", "test setDateTime")
	assertTrue(t, dt.Nanosecond() == 41, "setdatetime test: ns")

	dt.setDateTime(40, "nsec")
	dt.setDateTime(30, "sec")
	dt.setDateTime(20, "min")
	dt.setDateTime(10, "hr")
	assertTrue(t, dt.Format("{H}:{m}:{s}") == "10:20:30", "test setDateTime")
	assertTrue(t, dt.Nanosecond() == 40, "setdatetime test: ns")
}

func TestStartOfs(t *testing.T) {
	tTime := time.Date(2016, 1, 2, 10, 20, 30, 40, time.UTC)

	dt := New(tTime)
	dt.StartOfMinute()

	// Start Of Minute
	assertTrue(t, dt.Format("{h}:{m}:{s}") == "10:20:0", "start of minute")
	assertTrue(t, dt.Nanosecond() == 0, "start of minute")

	// Start Of Hour
	dt.t = tTime
	dt.StartOfHour()
	assertTrue(t, dt.Format("{D} {H}:{m}:{s}") == "2 10:0:0", "start of hour")
	assertTrue(t, dt.Nanosecond() == 0, "")

	// Start Of Day
	dt.t = tTime
	dt.StartOfDay()
	assertTrue(t, dt.Format("{D} {H}:{m}:{s}") == "2 0:0:0", "start of day")
	assertTrue(t, dt.Nanosecond() == 0, "")
}

func TestEndOfs(t *testing.T) {
	tTime := time.Date(2016, 1, 2, 10, 20, 30, 40, time.UTC)
	dt := New(tTime)

	// Start Of Minute
	dt.EndOfMinute()
	assertTrue(t, dt.Format("{h}:{m}:{s}") == "10:20:59", "EndOfTest: min")
	assertTrue(t, dt.Nanosecond() == 999999999, "EndOfTest: min")
	dt.t = tTime

	dt.EndOfHour()
	assertTrue(t, dt.Format("{h}:{m}:{s}") == "10:59:59", "EndOfTest: hr")
	assertTrue(t, dt.Nanosecond() == 999999999, "EndOfTest: hr")
	dt.t = tTime
}
