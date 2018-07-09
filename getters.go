package datetime

import (
	"time"
)

// Location returns the time zone information associated with t
func (dt DateTime) Location() *time.Location {
	return dt.Time().Location()
}

// Date returns the year, month, and day in which t occurs.
func (dt DateTime) Date() (year int, month time.Month, day int) {
	return dt.Time().Date()
}

// Clock returns the hour, minute, and second within the day specified by t.
func (dt DateTime) Clock() (hour, min, sec int) {
	return dt.Time().Clock()
}

// Year => ..2018..
func (dt DateTime) Year() int {
	return dt.Time().Year()
}

// Month returns the month of the year 1-12
func (dt DateTime) Month() int {
	return int(dt.Time().Month())
}

// Day returns the day of the month specified by t
func (dt DateTime) Day() int {
	return dt.Time().Day()
}

// Hour returns the hour within the day specified by t, in the range [0, 23].
func (dt DateTime) Hour() int {
	return dt.Time().Hour()
}

// Minute returns the minute offset within the hour specified by t, in the range [0, 59].
func (dt DateTime) Minute() int {
	return dt.Time().Minute()
}

// Second returns the second offset within the minute specified by t, in the range [0, 59].
func (dt DateTime) Second() int {
	return dt.Time().Second()
}

// Nanosecond returns the nanosecond offset within the second specified by t,
// in the range [0, 999999999].
func (dt DateTime) Nanosecond() int {
	return dt.Time().Nanosecond()
}

// Unix returns t as a Unix time, the number of seconds elapsed
// since January 1, 1970 UTC.
func (dt DateTime) Unix() int64 {
	return dt.Time().Unix()
}

// Weekday returns a number between 0 (sunday) and 6 (saturday)
func (dt DateTime) Weekday() int {
	return int(dt.Time().Weekday())
}

// YearDay returns the day of the year specified by t, in the range [1,365] for non-leap years,
// and [1,366] in leap years.
func (dt DateTime) YearDay() int {
	return dt.Time().YearDay()
}

// Time https://golang.org/pkg/time/#Time
func (dt DateTime) Time() time.Time {
	return dt.t
}

// DaysInMonth returns the number of days in a current month
func (dt DateTime) DaysInMonth() int {
	return dt.Copy().EndOfMonth().Day()
}

// DaysInYear returns the number of days in a current year
func (dt DateTime) DaysInYear() int {
	return dt.Copy().EndOfYear().YearDay()
}
