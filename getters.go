package datetime

// Year => ..2018..
func (dt *Datetime) Year() int {
	return dt.Time().Year()
}

// Month returns the month of the year 1-12
func (dt *Datetime) Month() int {
	return int(dt.Time().Month())
}

// Day returns the day of the month specified by t
func (dt *Datetime) Day() int {
	return dt.Time().Day()
}

// Hour returns the hour within the day specified by t, in the range [0, 23].
func (dt *Datetime) Hour() int {
	return dt.Time().Hour()
}

// Minute returns the minute offset within the hour specified by t, in the range [0, 59].
func (dt *Datetime) Minute() int {
	return dt.Time().Minute()
}

// Second returns the second offset within the minute specified by t, in the range [0, 59].
func (dt *Datetime) Second() int {
	return dt.Time().Second()
}

// Nanosecond returns the nanosecond offset within the second specified by t,
// in the range [0, 999999999].
func (dt *Datetime) Nanosecond() int {
	return dt.Time().Nanosecond()
}

// DayOfWeek returns a number between 0 (sunday) and 6 (saturday)
func (dt *Datetime) DayOfWeek() int {
	return int(dt.Time().Weekday())
}

// DayOfYear returns the day of the year specified by t, in the range [1,365] for non-leap years,
// and [1,366] in leap years.
func (dt *Datetime) DayOfYear() int {
	return dt.Time().YearDay()
}
