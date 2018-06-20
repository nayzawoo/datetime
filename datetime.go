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

// Time https://golang.org/pkg/time/#Time
func (dt DateTime) Time() time.Time {
	return dt.t
}
