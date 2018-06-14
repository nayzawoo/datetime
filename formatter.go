package datetime

import (
	"strings"
)

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
