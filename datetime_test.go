package datetime

import (
	"testing"
	"time"
)

func TestTimeOutput(t *testing.T) {
	in := time.Now()
	dt := New(in)
	out := dt.Time()

	assertTrue(t, out.Equal(in), "output should equal given time")

	in = in.AddDate(1, 0, 0)

	assertTrue(t, !dt.Time().Equal(in), "output should't equal given time")

	changedOutput := dt.Time().AddDate(1, 0, 0)
	assertTrue(t, !dt.Time().Equal(changedOutput), "changedOutput should't equal given time")
}

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
