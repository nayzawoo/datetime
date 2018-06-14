package datetime

import (
	"strings"
	"time"
)

// Common Formats
// https://carbon.nesbot.com/docs/#api-commonformats
// http://php.net/manual/en/class.datetime.php
// https://ruby-doc.org/core-2.2.0/Time.html
// https://docs.python.org/3/library/datetime.html

// Go Time Layout:  Mon Jan 2 15:04:05 2006
// time/format.go
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
	return dt.Time().Format(layout)
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

// ToRfc822String returns Rfc822
// example: 02 Jan 06 15:04 MST
func (dt *Datetime) ToRfc822String() string {
	return dt.Format(time.RFC822)
}

// ToRfc822ZString returns Rfc822, With Numeric Zone
// example: 02 Jan 06 15:04 -0700
func (dt *Datetime) ToRfc822ZString() string {
	return dt.Format(time.RFC822Z)
}

// ToRfc850String returns Rfc8659 Format
// example: Monday, 02-Jan-06 15:04:05 MST
func (dt *Datetime) ToRfc850String() string {
	return dt.Format(time.RFC850)
}

// ToRfc2822String returns Rfc2822 Format
// example: Mon, 02 Jan 2006 15:04:05 -0700
func (dt *Datetime) ToRfc2822String() string {
	const rfc2822Layout = "Mon, 02 Jan 2006 15:04:05 -0700"
	return dt.Time().Format(rfc2822Layout)
}
