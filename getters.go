package datetime

// Year => ..2018..
func (dt *Datetime) Year() int {
	return dt.Time().Year()
}

// Month returns the month of the year 1-12
func (dt *Datetime) Month() int {
	return int(dt.Time().Month())
}

// Day returns the day of the month specified by t.
func (dt *Datetime) Day() int {
	return dt.Time().Day()
}
