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
