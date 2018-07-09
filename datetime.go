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
func New(t time.Time) DateTime {
	dt := DateTime{
		t: t,
	}

	return dt
}

// Now returns the current time in given location or local time
func Now(loc *time.Location) DateTime {
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
func NewFromFormat(format, value string, loc *time.Location) (DateTime, error) {
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

// NewFromDate returns datetime from givem y,m,d
func NewFromDate(year, month, day int, loc *time.Location) DateTime {
	if loc == nil {
		loc = time.Local
	}

	t := time.Date(year, time.Month(month), day, 0, 0, 0, 0, loc)

	return New(t)
}

// Unix returns the local Time corresponding to the given Unix time
func Unix(sec int64, nsec int64) DateTime {
	t := time.Unix(sec, nsec)
	return New(t)
}

func (dt *DateTime) setTime(t time.Time) {
	dt.t = t
}

// In set current location to given location
func (dt DateTime) In(loc *time.Location) DateTime {
	t := dt.Time().In(loc)

	dt.setTime(t)
	return dt
}

// UTC set current location to UTC
func (dt DateTime) UTC() DateTime {
	return dt.In(time.UTC)
}

func (dt DateTime) String() string {
	return dt.DateTimeString()
}

// StartOfMinute returns 00s 0ns of current time
func (dt DateTime) StartOfMinute() DateTime {
	t := dt.Time().Truncate(time.Minute)
	dt.setTime(t)
	return dt
}

// StartOfHour returns 00m:00s 0ns of current time
func (dt DateTime) StartOfHour() DateTime {
	y, m, d := dt.Date()
	t := time.Date(y, m, d, dt.Hour(), 0, 0, 0, dt.Location())
	dt.setTime(t)
	return dt
}

// StartOfDay returns 00h:00m:00s 0ns of current time
func (dt DateTime) StartOfDay() DateTime {
	y, m, d := dt.Date()
	t := time.Date(y, m, d, 0, 0, 0, 0, dt.Location())
	dt.setTime(t)
	return dt
}

// StartOfMonth returns
func (dt DateTime) StartOfMonth() DateTime {
	y, m, _ := dt.Date()
	t := time.Date(y, m, 1, 0, 0, 0, 0, dt.Location())
	dt.setTime(t)
	return dt
}

// StartOfYear returns
func (dt DateTime) StartOfYear() DateTime {
	t := time.Date(dt.Year(), 1, 1, 0, 0, 0, 0, dt.Location())
	dt.setTime(t)
	return dt
}

// EndOfMinute returns
func (dt DateTime) EndOfMinute() DateTime {
	dt = dt.StartOfMinute()
	t := dt.Time().Add(time.Minute - time.Nanosecond)
	dt.setTime(t)
	return dt
}

// EndOfHour returns
func (dt DateTime) EndOfHour() DateTime {
	dt = dt.StartOfHour()
	t := dt.Time().Add(time.Hour - time.Nanosecond)
	dt.setTime(t)
	return dt
}

// EndOfDay returns
func (dt DateTime) EndOfDay() DateTime {
	y, m, d := dt.Date()
	t := time.Date(y, m, d+1, 0, 0, 0, 0, dt.Location()).Add(-time.Nanosecond)
	dt.setTime(t)
	return dt
}

// EndOfMonth returns
func (dt DateTime) EndOfMonth() DateTime {
	y, m, _ := dt.Date()
	t := time.Date(y, m+1, 1, 0, 0, 0, 0, dt.Location()).Add(-time.Nanosecond)
	dt.setTime(t)
	return dt
}

// EndOfYear returns
func (dt DateTime) EndOfYear() DateTime {
	t := time.Date(dt.Year()+1, 1, 1, 0, 0, 0, 0, dt.Location()).Add(-time.Nanosecond)
	dt.setTime(t)
	return dt
}

// Add add given duration
func (dt DateTime) Add(duration time.Duration) DateTime {
	t := dt.Time().Add(duration)

	dt.setTime(t)
	return dt
}

// Sub returns the datetime - duration
func (dt DateTime) Sub(duration time.Duration) DateTime {
	return dt.Add(-duration)
}

// AddSeconds returns the datetime t+d*time.Second.
func (dt DateTime) AddSeconds(seconds int) DateTime {
	t := dt.Time().Add(time.Duration(seconds) * time.Second)
	dt.setTime(t)
	return dt
}

// SubSeconds returns the datetime t-d*time.Second.
func (dt DateTime) SubSeconds(seconds int) DateTime {
	return dt.AddSeconds(-seconds)
}

// AddMinutes returns the datetime t+d*time.Minute.
func (dt DateTime) AddMinutes(minutes int) DateTime {
	return dt.Add(time.Duration(minutes) * time.Minute)
}

// SubMinutes returns the datetime t-d*time.Minute.
func (dt DateTime) SubMinutes(minutes int) DateTime {
	return dt.AddMinutes(-minutes)
}

// AddHours returns the datetime t+d*time.Hour.
func (dt DateTime) AddHours(hours int) DateTime {
	return dt.Add(time.Duration(hours) * time.Hour)
}

// SubHours returns the datetime t-d*time.Hour.
func (dt DateTime) SubHours(hours int) DateTime {
	return dt.AddHours(-hours)
}

// AddDays returns the datetime t+ days.
func (dt DateTime) AddDays(days int) DateTime {
	return dt.AddDate(0, 0, days)
}

// SubDays returns the datetime - days.
func (dt DateTime) SubDays(days int) DateTime {
	return dt.AddDate(0, 0, -days)
}

// AddMonths returns the datetime + months.
func (dt DateTime) AddMonths(months int) DateTime {
	return dt.AddDate(0, months, 0)
}

// SubMonths returns the datetime - months.
func (dt DateTime) SubMonths(months int) DateTime {
	return dt.AddDate(0, -months, 0)
}

// AddYears returns the datetime + years.
func (dt DateTime) AddYears(years int) DateTime {
	return dt.AddDate(years, 0, 0)
}

// SubYears returns the datetime + years.
func (dt DateTime) SubYears(years int) DateTime {
	return dt.AddDate(-years, 0, 0)
}

// AddTime returns the datetime t+-hr/min/sec.
func (dt DateTime) AddTime(hours, minutes, seconds int) DateTime {
	year, month, day := dt.Date()
	hour, minute, second := dt.Clock()

	t := time.Date(year, month, day, hour+hours, minute+minutes, second+seconds, dt.Nanosecond(), dt.Location())
	dt.setTime(t)
	return dt
}

// SubTime returns the datetime t+-hr/min/sec.
func (dt DateTime) SubTime(hours, minutes, seconds int) DateTime {
	return dt.AddTime(-hours, -minutes, -seconds)
}

// AddDate returns the time corresponding to adding the
// given number of years, months, and days to t.
func (dt DateTime) AddDate(years int, months int, days int) DateTime {
	t := dt.Time().AddDate(years, months, days)
	dt.setTime(t)
	return dt
}

// SubDate returns date-y,m,d
func (dt DateTime) SubDate(years int, months int, days int) DateTime {
	return dt.AddDate(-years, -months, -days)
}

// Eq reports whether t and u represent the same time instant.
func (dt DateTime) Eq(u DateTime) bool {
	return dt.Time().Equal(u.Time())
}

// Lt reports whether the datetime is before u
func (dt DateTime) Lt(u DateTime) bool {
	return dt.Time().Before(u.Time())
}

// Gt reports whether the datetime is after u.
func (dt DateTime) Gt(u DateTime) bool {
	return dt.Time().After(u.Time())
}

// Lte reports whether the datetime is before or same time instant.
func (dt DateTime) Lte(u DateTime) bool {
	return dt.Eq(u) || dt.Lt(u)
}

// Gte reports whether the datetime is after u or same time instant
func (dt DateTime) Gte(u DateTime) bool {
	return dt.Eq(u) || dt.Gt(u)
}

// Diff returns the duration t-u. If the result exceeds the maximum (or minimum)
// value that can be stored in a Duration, the maximum (or minimum) duration
// will be returned.
func (dt DateTime) Diff(u DateTime) time.Duration {
	return dt.Time().Sub(u.Time())
}

// IsLeapYear determine if current year is leap year
func (dt DateTime) IsLeapYear() bool {
	year := dt.Year()
	return year%4 == 0 && (year%100 != 0 || year%400 == 0)
}
