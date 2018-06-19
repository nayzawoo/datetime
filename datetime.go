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

// Time https://golang.org/pkg/time/#Time
func (dt DateTime) Time() time.Time {
	return dt.t
}
