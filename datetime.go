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

func (dt *DateTime) setDateTime(value int, setType string) {
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
	dt.setDateTime(0, "nsec")
	dt.setDateTime(0, "sec")
	return dt
}

// StartOfHour return 00m:00s 0ns of current time
func (dt *DateTime) StartOfHour() *DateTime {
	dt.StartOfMinute()
	dt.setDateTime(0, "min")
	return dt
}

// StartOfDay return 00h:00m:00s 0ns of current time
func (dt *DateTime) StartOfDay() *DateTime {
	dt.StartOfHour()
	dt.setDateTime(0, "hour")
	return dt
}

// EndOfMinute returns 59s ns of current time
func (dt *DateTime) EndOfMinute() *DateTime {
	dt.StartOfMinute()
	dt.t = dt.Time().Add(time.Minute - time.Nanosecond)
	return dt
}

// EndOfHour returns 59s ns of current time
func (dt *DateTime) EndOfHour() *DateTime {
	dt.StartOfHour()
	dt.t = dt.Time().Add(time.Hour - time.Nanosecond)
	return dt
}
