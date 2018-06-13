package datetime

import (
	"strings"
	"time"
)

// Datetime Wrapper Like Carbon, Moment, DaysJS
type Datetime struct {
	t *time.Time
}

// New : Build
func New(t *time.Time) *Datetime {
	dt := &Datetime{
		t: t,
	}

	return dt
}

// Time https://golang.org/pkg/time/#Time
func (dt *Datetime) Time() *time.Time {
	return dt.t
}

// FormatToStdLayout returns buildin standard time layout
// Formats =>
// Y : A full numeric representation of a year, 4 digits
// y : A two digit representation of a year
// d : Day of the month as a zero-padded decimal number.
func FormatToStdLayout(format string) string {
	r := strings.NewReplacer(
		// Day
		"d", "02",
		"D", "Mon",
		"j", "2",
		"l", "Monday",

		// Month
		"F", "January",
		"m", "01",
		"M", "Jan",
		"n", "1",

		// Year
		"Y", "2006",
		"y", "06",

		// Time
		"g", "3",
		"h", "03",
		"H", "15",
		"i", "04",
		"s", "05",
		"a", "pm",
		"A", "PM",
	)

	layout := r.Replace(format)

	return layout
}
