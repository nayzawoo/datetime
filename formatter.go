package datetime

import (
	"strconv"
	"strings"
	"time"
)

var formatReplacer = strings.NewReplacer(
	// Year
	"{YYYY}", "2006",
	"{YY}", "06",

	// Month
	"{M}", "1",
	"{MM}", "01",
	"{MMM}", "Jan",
	"{MMMM}", "January",

	// Day
	"{D}", "2",
	"{DD}", "02",
	"{DDD}", "Mon",
	"{DDDD}", "Monday",

	// Hour
	"{h}", "3",
	"{hh}", "03",
	// {H}, "24 hour format without zero padded"
	"{HH}", "15",

	// Minute
	"{m}", "4",
	"{mm}", "04",

	// Second
	"{s}", "5",
	"{ss}", "05",

	// Meridian
	"{pm}", "pm",
	"{PM}", "PM",

	// Zone
	"{z}", "-07",
	"{zz}", "-0700",
	"{zzz}", "-070000",
	"{z:}", "-07:00",
	"{z::}", "-07:00:00",
	"{Z}", "MST",
)

// Format returns datetime string according to given format
func Format(t *time.Time, format string) string {
	date := t.Format(formatToStdLayout(format))

	// format non zero padded 24hr
	if strings.Contains(date, "{H}") {
		date = strings.Replace(date, "{H}", strconv.Itoa(t.Hour()), -1)
	}

	return date
}

// Format returns datetime string according to given format
func (dt *Datetime) Format(format string) string {
	t := dt.Time()
	return Format(t, format)
}

func formatToStdLayout(format string) string {

	layout := formatReplacer.Replace(format)

	return layout
}
