package datetime

import (
	"strings"
)

var formatToLayoutReplacer = strings.NewReplacer(
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

// ToLayout returns formatted datetime string according to given layout.
func (dt *Datetime) ToLayout(layout string) string {
	return dt.Time().Format(layout)
}

// Format returns formatted datetime string according to given format
func (dt *Datetime) Format(format string) string {
	return dt.ToLayout(FormatToStdLayout(format))
}

// FormatToStdLayout returns buildin standard time layout
func FormatToStdLayout(format string) string {

	layout := formatToLayoutReplacer.Replace(format)

	return layout
}
