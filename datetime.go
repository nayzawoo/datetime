package datetime

import (
	"strings"
	"time"
)

var (
	testNow *time.Time
)

// DateTime Wrapper Like Carbon, Moment, DaysJS
type DateTime struct {
	t              time.Time
	formatReplacer strings.Replacer
}

// New : Build
func New(t time.Time) *DateTime {
	dt := &DateTime{
		t: t,
	}

	return dt
}

// Now returns the current time in given location or local time
func Now(loc *time.Location) *DateTime {
	var t time.Time
	if HasTestNow() {
		t = *testNow
	} else {
		if loc == nil {
			t = time.Now()
		} else {
			t = time.Now().In(loc)
		}
	}

	return New(t)
}

// SetTestNow set current for testing
func SetTestNow(t time.Time) {
	testNow = &t
}

// HasTestNow check if mocking current time
func HasTestNow() bool {
	return testNow != nil
}

// ResetTestNow reset current for testing
func ResetTestNow() {
	testNow = nil
}

// NewFromFormat parses date and returns DatetimeObject
func NewFromFormat(format, value string, loc *time.Location) (*DateTime, error) {
	layout := formatToStdLayout(format)
	layout = fixLayoutFor24Hour(layout)

	t, err := time.Parse(layout, value)

	return New(t), err
}

// NewFromDate -
func NewFromDate(year, month, day int, loc *time.Location) *DateTime {
	if loc == nil {
		loc = time.UTC
	}

	t := time.Date(year, time.Month(month), day, 0, 0, 0, 0, loc)

	return New(t)
}

// Copy return copy of DateTime
func (dt *DateTime) Copy() *DateTime {
	return New(dt.Time())
}

func (dt *DateTime) String() string {
	return dt.DateTimeString()
}

// DaysInMonth returns the number of days in a current months
func (dt *DateTime) DaysInMonth() int {
	return dt.EndOfMonth().Day()
}

// StartOfMinute returns 00s 0ns of current time
func (dt *DateTime) StartOfMinute() *DateTime {
	dt.t = dt.t.Truncate(time.Minute)
	return dt
}

// StartOfHour returns 00m:00s 0ns of current time
func (dt *DateTime) StartOfHour() *DateTime {
	y, m, d := dt.Date()
	dt.t = time.Date(y, m, d, dt.Hour(), 0, 0, 0, dt.Location())
	return dt
}

// StartOfDay returns 00h:00m:00s 0ns of current time
func (dt *DateTime) StartOfDay() *DateTime {
	y, m, d := dt.Date()
	dt.t = time.Date(y, m, d, 0, 0, 0, 0, dt.Location())
	return dt
}

// StartOfMonth returns
func (dt *DateTime) StartOfMonth() *DateTime {
	y, m, _ := dt.Date()
	dt.t = time.Date(y, m, 1, 0, 0, 0, 0, dt.Location())
	return dt
}

// StartOfYear returns
func (dt *DateTime) StartOfYear() *DateTime {
	dt.t = time.Date(dt.Year(), 1, 1, 0, 0, 0, 0, dt.Location())
	return dt
}

// EndOfMinute returns
func (dt *DateTime) EndOfMinute() *DateTime {
	dt.StartOfMinute()
	dt.t = dt.Time().Add(time.Minute - time.Nanosecond)
	return dt
}

// EndOfHour returns
func (dt *DateTime) EndOfHour() *DateTime {
	dt.StartOfHour()
	dt.t = dt.Time().Add(time.Hour - time.Nanosecond)
	return dt
}

// EndOfDay returns
func (dt *DateTime) EndOfDay() *DateTime {
	y, m, d := dt.Date()

	t := time.Date(y, m, d+1, 0, 0, 0, 0, dt.Location()).Add(-time.Nanosecond)

	dt.t = t
	return dt
}

// EndOfMonth returns
func (dt *DateTime) EndOfMonth() *DateTime {
	y, m, _ := dt.Date()
	// startOfNextMonth - 1nsec
	dt.t = time.Date(y, m+1, 1, 0, 0, 0, 0, dt.Location()).Add(-time.Nanosecond)

	return dt
}

// EndOfYear returns
func (dt *DateTime) EndOfYear() *DateTime {
	dt.t = time.Date(dt.Year()+1, 1, 1, 0, 0, 0, 0, dt.Location()).Add(-time.Nanosecond)

	return dt
}

// AddTime returns the datetime t+-hr/min/sec.
func (dt *DateTime) AddTime(hour, minute, second int) *DateTime {
	y, m, d := dt.Date()
	_hr, _min, _sec := dt.Clock()

	dt.t = time.Date(y, m, d, _hr+hour, _min+minute, _sec+second, dt.Nanosecond(), dt.Location())

	return dt
}

// SubTime returns the datetime t+-hr/min/sec.
func (dt *DateTime) SubTime(hour, minute, second int) *DateTime {
	dt.AddTime(-hour, -minute, -second)
	return dt
}

// AddDate returns the time corresponding to adding the
// given number of years, months, and days to t.
func (dt *DateTime) AddDate(years int, months int, days int) *DateTime {
	dt.t = dt.Time().AddDate(years, months, days)

	return dt
}

// SubDate returns date-y,m,d
func (dt *DateTime) SubDate(years int, months int, days int) *DateTime {
	dt.t = dt.Time().AddDate(-years, -months, -days)
	return dt
}

// ===========
// Comparison
// ===========

// Eq return true if they are same
// Alias of go built in Equal
func (dt *DateTime) Eq(datetime *DateTime) bool {
	return dt.Time().Equal(datetime.Time())
}
