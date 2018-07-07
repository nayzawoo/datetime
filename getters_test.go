package datetime

import (
	"testing"
	"time"
)

func TestYearToNanoSecond(t *testing.T) {
	tm := time.Date(2018, time.February, 4, 20, 10, 5, 9, time.UTC)
	dt := New(tm)
	assertTrue(t, dt.Year() == 2018, "Year => 2018")
	assertTrue(t, dt.Month() == 2, "Month => 2(February)")
	assertTrue(t, dt.Day() == 4, "Day => 4")
	assertTrue(t, dt.Hour() == 20, "Hour => 20")
	assertTrue(t, dt.Minute() == 10, "Minute => 10")
	assertTrue(t, dt.Second() == 5, "Second => 5")
	assertTrue(t, dt.Nanosecond() == 9, "Nanosecond => 9")
	assertTrue(t, dt.Weekday() == 0, "Weekday => 1(Sun)")
	assertTrue(t, dt.YearDay() == 35, "DayOfYear => 35")

	y, m, d := dt.Date()

	assertTrue(t, y == 2018, "")
	assertTrue(t, m == time.February, "")
	assertTrue(t, d == 4, "")

	h, i, s := dt.Clock()

	assertTrue(t, h == 20, "")
	assertTrue(t, i == 10, "")
	assertTrue(t, s == 5, "")
}

func TestGetLocation(t *testing.T) {
	loc, _ := time.LoadLocation("America/New_York")
	dt := NewFromDate(2000, 1, 1, loc)
	assertTrue(t, dt.Location().String() == "America/New_York", "test get location")
}

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

func TestDaysInYearMonth(t *testing.T) {
	dt := NewFromDate(2000, 1, 1, time.UTC)
	assertTrue(t, dt.DaysInMonth() == 31, "TestDaysInMonth")

	dt = NewFromDate(2000, 2, 1, time.UTC)
	assertTrue(t, dt.DaysInMonth() == 29, "TestDaysInMonth for Leap Feb")

	dt = NewFromDate(2001, 2, 1, time.UTC)
	assertTrue(t, dt.DaysInMonth() == 28, "TestDaysInMonth")

	dt = NewFromDate(2001, 4, 1, time.UTC)
	assertTrue(t, dt.DaysInMonth() == 30, "TestDaysInMonth")

	dt = NewFromDate(2000, 2, 1, time.UTC)
	assertTrue(t, dt.DaysInYear() == 366, "TestDaysInYear")
}
