package datetime

import (
	"strings"
	"time"
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

// NewFromFormat parses date and returns DatetimeObject
func NewFromFormat(format, value string, loc *time.Location) (DateTime, error) {
	layout := formatToStdLayout(format)
	layout = fixLayoutFor24Hour(layout)

	t, err := time.Parse(layout, value)

	return New(t), err
}

// NewFromDate -
func NewFromDate(year, month, day int, loc *time.Location) DateTime {
	if loc == nil {
		loc = time.UTC
	}

	t := time.Date(year, time.Month(month), day, 0, 0, 0, 0, loc)

	return New(t)
}

// DaysInMonth returns the number of days in a current months
func (dt *DateTime) DaysInMonth() int {
	return dt.EndOfMonth().Day()
}

func (dt *DateTime) set(value int, setType string) {
	t := dt.Time()
	Y, M, D := t.Date()
	h, m, s := t.Clock()
	ns := dt.Nanosecond()
	switch setType {
	case "nanosecond", "nsec":
		ns = value
		break
	case "second", "sec":
		s = value
		break
	case "minute", "min":
		m = value
		break
	case "hour", "hr":
		h = value
		break
	case "day":
		D = value
		break
	case "month":
		M = time.Month(value)
		break
	case "year":
		Y = value
	default:
		panic("unrecognized type")
	}

	dt.t = time.Date(Y, M, D, h, m, s, ns, t.Location())
}

// StartOfMinute returns 00s 0ns of current time
func (dt *DateTime) StartOfMinute() *DateTime {
	dt.set(0, "nsec")
	dt.set(0, "sec")
	return dt
}

// StartOfHour return 00m:00s 0ns of current time
func (dt *DateTime) StartOfHour() *DateTime {
	dt.StartOfMinute()
	dt.set(0, "min")
	return dt
}

// StartOfDay return 00h:00m:00s 0ns of current time
func (dt *DateTime) StartOfDay() *DateTime {
	dt.StartOfHour()
	dt.set(0, "hour")
	return dt
}

// StartOfMonth return
func (dt *DateTime) StartOfMonth() *DateTime {
	y, m, _ := dt.Time().Date()
	dt.t = time.Date(y, m, 1, 0, 0, 0, 0, dt.Time().Location())
	return dt
}

// StartOfYear return
func (dt *DateTime) StartOfYear() *DateTime {
	dt.t = time.Date(dt.Time().Year(), 1, 1, 0, 0, 0, 0, dt.Time().Location())
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
	y, m, d := dt.Time().Date()

	t := time.Date(y, m, d+1, 0, 0, 0, 0, dt.Time().Location()).Add(-time.Nanosecond)

	dt.t = t
	return dt
}

// EndOfMonth returns
func (dt *DateTime) EndOfMonth() *DateTime {
	t := dt.Time()
	y, m, _ := t.Date()
	// startOfNextMonth - 1nsec
	dt.t = time.Date(y, m+1, 1, 0, 0, 0, 0, t.Location()).Add(-time.Nanosecond)

	return dt
}

// EndOfYear returns
func (dt *DateTime) EndOfYear() *DateTime {
	t := dt.Time()
	dt.t = time.Date(t.Year()+1, 1, 1, 0, 0, 0, 0, t.Location()).Add(-time.Nanosecond)

	return dt
}
