# DateTime
[![Build Status](https://travis-ci.org/nayzawoo/datetime.svg?branch=master)](https://travis-ci.org/nayzawoo/datetime)
[![codecov](https://codecov.io/gh/nayzawoo/datetime/branch/master/graph/badge.svg)](https://codecov.io/gh/nayzawoo/datetime)
[![Go Report Card](https://goreportcard.com/badge/github.com/nayzawoo/datetime)](https://goreportcard.com/report/github.com/nayzawoo/datetime)

## Installation

```
go get github.com/nayzawoo/datetime
```

```go
package main

import (
	"github.com/nayzawoo/datetime"
	"time"
)

datetime.Now(nil)
datetime.NewFromFormat("{YYYY}-{M}-{D}", "2018-2-3", time.UTC)
year2000 := datetime.NewFromDate(2000, 1, 1, time.UTC)
year3000 := datetime.NewFromDate(3000, 1, 1, time.UTC)

t := time.Date(2018, 2, 3, 10, 20, 30, 40, time.UTC)
dt := datetime.New(t)

dt.DateTimeString() // 2018-02-03 10:20:30
dt.Year() // 2018
dt.Month() // 2
dt.Day() // 3
dt.Hour() // 10
dt.Minute() // 20
dt.Second() // 30
dt.Nanosecond() // 40
dt.Unix() // 1517653230
dt.Weekday() // 6
dt.YearDay() // 34
dt.DaysInMonth() // 28
dt.DaysInYear() // 365
dt.Clock() // [10 20 30]
dt.Date() // [2018 February 3]
dt.Time() // https://golang.org/pkg/time/#Time

year2000.Eq(year2000) // true
year2000.Lt(year2000) // false
year2000.Lte(year2000) // true
year3000.Gt(year2000) // true
year3000.Gte(year2000) // true
year3000.Diff(year2000) // time.Duration

dt.EndOfMinute() // 2018-02-03 10:20:59
dt.EndOfHour() // 2018-02-03 10:59:59
dt.EndOfDay() // 2018-02-03 23:59:59
dt.EndOfMonth() // 2018-02-28 23:59:59
dt.EndOfYear() // 2018-12-31 23:59:59

dt.StartOfMinute() // 2018-02-03 10:20:00
dt.StartOfHour() // 2018-02-03 10:00:00
dt.StartOfDay() // 2018-02-03 00:00:00
dt.StartOfMonth() // 2018-02-01 00:00:00
dt.StartOfYear() // 2018-01-01 00:00:00

dt.Add(time.Second * 10) // 2018-02-03 10:20:40
dt.Sub(time.Second * 10) // 2018-02-03 10:20:20
dt.AddTime(5, 10, 15) // 2018-02-03 15:30:45
dt.SubTime(5, 10, 15) // 2018-02-03 05:10:15
dt.AddDate(1, 2, 3) // 2019-04-06 10:20:30
dt.SubDate(1, 2, 3) // 2016-11-30 10:20:30
dt.AddSeconds(10) // 2018-02-03 10:20:40
dt.SubSeconds(10) // 2018-02-03 10:20:20
dt.AddMinutes(10) // 2018-02-03 10:30:30
dt.SubMinutes(10) // 2018-02-03 10:10:30
dt.AddHours(10) // 2018-02-03 20:20:30
dt.SubHours(10) // 2018-02-03 00:20:30
dt.AddDays(1) // 2018-02-04 10:20:30
dt.SubDays(1) // 2018-02-02 10:20:30
dt.AddMonths(1) // 2018-03-03 10:20:30
dt.SubMonths(1) // 2018-01-03 10:20:30
dt.AddYears(1) // 2019-02-03 10:20:30
dt.SubYears(1) // 2017-02-03 10:20:30

dt.In(time.UTC) // 2018-02-03 10:20:30
dt.UTC() // 2018-02-03 10:20:30

## Format

```go
// Using predefined layouts
//
// ANSIC, UnixDate, RubyDate, RFC822, RFC822Z, RFC850, RFC1123, RFC1123Z, RFC3339, RFC3339Nano,
// Kitchen, Stamp, StampMilli, StampMicro, StampNano
dt.Format(time.RubyDate)

// Using custom formats
dt.Format("{YYYY}-{MMM}-{DD}") // eg: 2018-Jan-04
```

## Custom Formats

| Format | Meaning | Example  |
| --- | --- | --- |
| YYYY | Year | 2006 |
| YY | Year | 06 |
| M | Month | 1 |
| MM | Month | 01 |
| MMM | Month | Jan |
| MMMM | Month | January |
| D | Day | 2 |
| DD | Day | 02 |
| DDD | Day | Mon |
| DDDD | Day | Monday |
| h | Hour | 1 |
| hh | Hour(Zero Padded) | 01 |
| H | Hour | 15 |
| HH | Hour(Zero Padded) | 15 |
| m | Minute | 4 |
| mm | Minute | 04 |
| s | Second | 5 |
| ss | Second | 05 |
| pm | AM or PM | pm |
| PM | AM or PM | PM |
| z | Timezone | -07 |
| zz | Timezone | -0700 |
| zzz | Timezone | -070000 |
| z: | Timezone | -07:00 |
| z:: | Timezone | -07:00:00 |
| Z | Timezone | UTC |

## License

[The MIT License (MIT)](https://raw.githubusercontent.com/nayzawoo/datetime/master/LICENSE)
