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
		t = time.Now()
	}

	if loc != nil {
		t = t.In(loc)
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

	var t time.Time
	var err error

	if loc == nil {
		t, err = time.Parse(layout, value)
	} else {
		t, err = time.ParseInLocation(layout, value, loc)
	}

	return New(t), err
}

// NewFromDate -
func NewFromDate(year, month, day int, loc *time.Location) *DateTime {
	if loc == nil {
		loc = time.Local
	}

	t := time.Date(year, time.Month(month), day, 0, 0, 0, 0, loc)

	return New(t)
}

// Unix returns the local Time corresponding to the given Unix time
func Unix(sec int64, nsec int64) *DateTime {
	t := time.Unix(sec, nsec)
	return New(t)
}

// In set current location to given location
func (dt *DateTime) In(loc *time.Location) *DateTime {
	dt.t = dt.Time().In(loc)
	return dt
}

// UTC set current location to UTC
func (dt *DateTime) UTC() *DateTime {
	return dt.In(time.UTC)
}

// Copy return new datetime
func (dt *DateTime) Copy() *DateTime {
	return New(dt.Time())
}

func (dt *DateTime) String() string {
	return dt.DateTimeString()
}

// StartOfMinute returns 00s 0ns of current time
func (dt *DateTime) StartOfMinute() *DateTime {
	dt.t = dt.Time().Truncate(time.Minute)
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

// Add add given duration
func (dt *DateTime) Add(duration time.Duration) *DateTime {
	dt.t = dt.Time().Add(duration)

	return dt
}

// Sub returns the datetime - duration
func (dt *DateTime) Sub(duration time.Duration) *DateTime {
	dt.Add(-duration)

	return dt
}

// AddSeconds returns the datetime t+d*time.Second.
func (dt *DateTime) AddSeconds(seconds int) *DateTime {
	dt.t = dt.Time().Add(time.Duration(seconds) * time.Second)

	return dt
}

// SubSeconds returns the datetime t-d*time.Second.
func (dt *DateTime) SubSeconds(seconds int) *DateTime {
	dt.AddSeconds(-seconds)

	return dt
}

// AddMinutes returns the datetime t+d*time.Minute.
func (dt *DateTime) AddMinutes(minutes int) *DateTime {
	dt.t = dt.Time().Add(time.Duration(minutes) * time.Minute)

	return dt
}

// SubMinutes returns the datetime t-d*time.Minute.
func (dt *DateTime) SubMinutes(minutes int) *DateTime {
	dt.AddMinutes(-minutes)

	return dt
}

// AddHours returns the datetime t+d*time.Hour.
func (dt *DateTime) AddHours(hours int) *DateTime {
	dt.t = dt.Time().Add(time.Duration(hours) * time.Hour)

	return dt
}

// SubHours returns the datetime t-d*time.Hour.
func (dt *DateTime) SubHours(hours int) *DateTime {
	dt.AddHours(-hours)

	return dt
}

// AddDays returns the datetime t+ days.
func (dt *DateTime) AddDays(days int) *DateTime {
	dt.AddDate(0, 0, days)

	return dt
}

// SubDays returns the datetime - days.
func (dt *DateTime) SubDays(days int) *DateTime {
	dt.AddDate(0, 0, -days)

	return dt
}

// AddMonths returns the datetime + months.
func (dt *DateTime) AddMonths(months int) *DateTime {
	dt.AddDate(0, months, 0)

	return dt
}

// SubMonths returns the datetime - months.
func (dt *DateTime) SubMonths(months int) *DateTime {
	dt.AddDate(0, -months, 0)

	return dt
}

// AddYears returns the datetime + years.
func (dt *DateTime) AddYears(years int) *DateTime {
	dt.AddDate(years, 0, 0)

	return dt
}

// SubYears returns the datetime + years.
func (dt *DateTime) SubYears(years int) *DateTime {
	dt.AddDate(-years, 0, 0)

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
// Comparison & difference
// ===========

// Eq reports whether t and u represent the same time instant.
func (dt *DateTime) Eq(u *DateTime) bool {
	return dt.Time().Equal(u.Time())
}

// Lt reports whether the datetime is before u
func (dt *DateTime) Lt(u *DateTime) bool {
	return dt.Time().Before(u.Time())
}

// Lte reports whether the datetime is before or same time instant.
func (dt *DateTime) Lte(u *DateTime) bool {
	return dt.Time().Equal(u.Time()) || dt.Time().Before(u.Time())
}

// Gt reports whether the datetime is after u.
func (dt *DateTime) Gt(u *DateTime) bool {
	return dt.Time().After(u.Time())
}

// Gte reports whether the datetime is after u or same time instant
func (dt *DateTime) Gte(u *DateTime) bool {
	return dt.Time().Equal(u.Time()) || dt.Time().After(u.Time())
}

// Diff returns the duration t-u. If the result exceeds the maximum (or minimum)
// value that can be stored in a Duration, the maximum (or minimum) duration
// will be returned.
func (dt *DateTime) Diff(u *DateTime) time.Duration {
	return dt.Time().Sub(u.Time())
}

// IsLeapYear determine if current year is leap year
func (dt *DateTime) IsLeapYear() bool {
	year := dt.Year()
	return year%4 == 0 && (year%100 != 0 || year%400 == 0)
}
