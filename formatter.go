package datetime

import (
	"strconv"
	"strings"
	"time"
)

// Common Formats
// https://carbon.nesbot.com/docs/#api-commonformats
// http://php.net/manual/en/class.datetime.php

// time format to build time layout replacer
// time/format.go
// https://ruby-doc.org/core-2.2.0/Time.html
// https://docs.python.org/3/library/datetime.html
var simpleFormatReplacer = strings.NewReplacer(
	// Year
	"%Y", "2006",
	"%y", "06",

	// Month
	"%-m", "1",
	"%m", "01",
	"%b", "Jan",
	"%B", "January",

	// Day
	"%-d", "2",
	"%d", "02",

	// Weekday:
	"%a", "Mon",
	"%A", "Monday",

	// Hour
	"%H", "15",
	"%-I", "3",
	"%I", "03",

	// Minute
	"%-M", "4",
	"%M", "04",

	// Second
	"%-S", "5",
	"%S", "05",

	// Meridian
	"%p", "pm",
	"%P", "PM",
)

// ToLayout returns formatted datetime string according to given layout.
func (dt *Datetime) ToLayout(layout string) string {
	date := dt.Time().Format(layout)

	if dt.Hour() < 10 && strings.Contains(date, "%-H") {
		date = strings.Replace(date, "%-H", strconv.Itoa(dt.Hour()), -1)
	}

	return date
}

// Format returns formatted datetime string according to given format
func (dt *Datetime) Format(format string) string {
	return dt.ToLayout(dt.formatToStdLayout(format))
}

func (dt *Datetime) getFormatReplacer() *strings.Replacer {
	if dt.formatReplacer == nil {
		dt.formatReplacer = simpleFormatReplacer
	}

	return dt.formatReplacer
}

func (dt *Datetime) formatToStdLayout(format string) string {

	layout := dt.getFormatReplacer().Replace(format)

	return layout
}

// ToAtomString returns Atom
// example: 2005-08-15T15:52:01+00:00)
func (dt *Datetime) ToAtomString() string {
	const atomLayout = "2006-01-02T15:04:05-07:00"
	return dt.Time().Format(atomLayout)
}

// ToCookieString returns HTTP Cookie Format
// example: Monday, 02-Jan-2006 15:04:05 MST
func (dt *Datetime) ToCookieString() string {
	return dt.Time().Format("Monday, 02-Jan-2006 15:04:05 MST")
}

// ToRFC822String returns RFC822
// example: 02 Jan 06 15:04 MST
func (dt *Datetime) ToRFC822String() string {
	return dt.Format(time.RFC822)
}

// ToRFC822ZString returns RFC822, With Numeric Zone
// example: 02 Jan 06 15:04 -0700
func (dt *Datetime) ToRFC822ZString() string {
	return dt.Format(time.RFC822Z)
}

// ToRFC850String returns RFC8659 Format
// example: Monday, 02-Jan-06 15:04:05 MST
func (dt *Datetime) ToRFC850String() string {
	return dt.Format(time.RFC850)
}

// ToRFC2822String returns RFC2822 Format
// example: Mon, 02 Jan 2006 15:04:05 -0700
func (dt *Datetime) ToRFC2822String() string {
	const rfc2822Layout = "Mon, 02 Jan 2006 15:04:05 -0700"
	return dt.Time().Format(rfc2822Layout)
}

// ToRFC3339String returns RFC3339 Format
// 2006-01-02T15:04:05+07:00
func (dt *Datetime) ToRFC3339String() string {
	return dt.Time().Format("2006-01-02T15:04:05-07:00")
}

// Todo
// RFC1123
// RSS
// W3C
