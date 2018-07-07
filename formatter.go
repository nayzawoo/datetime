package datetime

import (
	"strconv"
	"strings"
	"time"
)

// eg layout: Mon Jan 02 15:04:05 -0700 2006

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
func Format(t time.Time, format string) string {
	if isStdFormat(format) {
		return t.Format(format)
	}

	date := t.Format(formatToStdLayout(format))

	// format non zero padded 24hr
	if strings.Contains(date, "{H}") {
		date = strings.Replace(date, "{H}", strconv.Itoa(t.Hour()), -1)
	}

	return date
}

// Format returns datetime string according to given format
func (dt *DateTime) Format(format string) string {
	t := dt.Time()
	return Format(t, format)
}

func formatToStdLayout(format string) string {

	layout := formatReplacer.Replace(format)

	return layout
}

func fixLayoutFor24Hour(layout string) string {

	layout = strings.Replace(layout, "{H}", "15", -1)

	return layout
}

// DateTimeString returns datetime simple format eg: 2016-01-02 15:04:05
func (dt *DateTime) DateTimeString() string {
	return dt.Time().Format("2006-01-02 15:04:05")
}

func isStdFormat(format string) bool {
	switch format {
	case time.ANSIC,
		time.UnixDate,
		time.RubyDate,
		time.RFC822,
		time.RFC822Z,
		time.RFC850,
		time.RFC1123,
		time.RFC1123Z,
		time.RFC3339,
		time.RFC3339Nano,
		time.Kitchen,
		time.Stamp,
		time.StampMilli,
		time.StampMicro,
		time.StampNano:
		return true
	}

	return false
}
