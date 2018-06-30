# DateTime
[![Build Status](https://travis-ci.org/nayzawoo/datetime.svg?branch=master)](https://travis-ci.org/nayzawoo/datetime)
[![codecov](https://codecov.io/gh/nayzawoo/datetime/branch/master/graph/badge.svg)](https://codecov.io/gh/nayzawoo/datetime)
[![Go Report Card](https://goreportcard.com/badge/github.com/nayzawoo/datetime)](https://goreportcard.com/report/github.com/nayzawoo/datetime)

Date and Time For Go(Under Development)

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

func main() {
    datetime.Now()
    datetime.NewFromFormat("{YYYY}-{M}-{D}", "2018-2-3", time.UTC)
    datetime.NewFromDate(2018, 1, 30)

    t := time.Date(2018, 2, 3, 10, 20, 30, 0, time.UTC)
    dt := datetime.New(t)
	dt.StartOfMinute() // 2018-02-03 10:20:00
	dt.StartOfHour() // 2018-02-03 10:00:00
	dt.StartOfDay() // 2018-02-03 00:00:00
	dt.StartOfMonth() // 2018-02-01 00:00:00
	dt.StartOfYear() // 2018-01-01 00:00:00

	dt.EndOfMinute() // 2018-01-01 00:00:59
	dt.EndOfHour() // 2018-01-01 00:59:59
	dt.EndOfDay() // 2018-01-01 23:59:59
	dt.EndOfMonth() // 2018-01-31 23:59:59
	dt.EndOfYear() // 2018-12-31 23:59:59

	dt.DateTimeString() // 2018-12-31 23:59:59
	dt.Year() // 2018
	dt.Month() // 12
	dt.Day() // 31
	dt.Hour() // 23
	dt.Minute() // 59
	dt.Second() // 59
	dt.Nanosecond() // 999999999
	dt.Weekday() // 1

	dt.AddTime(24, 60, 60) // 2018-01-02 01:01:00
	dt.SubTime(24, 60, 60) // 2018-01-01 00:00:00
	dt.AddDate(1, 2, 3) // 2019-03-04 00:00:00
	dt.SubDate(1, 2, 3) // 2018-01-01 00:00:00
	dt.Clock() // 0, 0, 0
	dt.Date() // 2018, January, 1

	dt.Time() // https://golang.org/pkg/time/#Time
}
```

## Formats

```go
dt.Format("{YYYY}-{MMM}-{DD}") // eg: 2018-Jan-04
```

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