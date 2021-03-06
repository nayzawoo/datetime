package datetime

import (
	"testing"
	"time"
)

func TestFormat(t *testing.T) {
	loc, _ := time.LoadLocation("America/New_York")
	tm := time.Date(2018, time.February, 4, 20, 7, 5, 9, loc)
	dt := New(tm)
	assertTrue(t, dt.Format("{YYYY} {YY}") == "2018 18", "Format Test: Year")
	assertTrue(t, dt.Format("{MMMM} {MMM} {MM} {M}") == "February Feb 02 2", "Format Test: Month")
	assertTrue(t, dt.Format("{DDDD} {DDD} {DD} {D}") == "Sunday Sun 04 4", "Format Test: Day")
	assertTrue(t, dt.Format("{HH} {H} {hh} {h}") == "20 20 08 8", "Format Test: Hour")
	assertTrue(t, dt.Format("{mm} {m}") == "07 7", "Format Test: Minute")
	assertTrue(t, dt.Format("{ss} {s}") == "05 5", "Format Test: Second")
	assertTrue(t, dt.Format("{z} {zz} {zzz}") == "-05 -0500 -050000", "Format Test: TimeZone")
	assertTrue(t, dt.Format("{z:} {z::}") == "-05:00 -05:00:00", "Format Test: TimeZone With Colon")
	assertTrue(t, dt.Format("{Z}") == "EST", "Format Test: TimeZone Name")
	assertTrue(t, dt.Format("{pm} {PM}") == "pm PM", "Format Test: Meridian")
	assertTrue(t, dt.Format(time.RFC822) == "04 Feb 18 20:07 EST", "")
	assertTrue(t, dt.DateTimeString() == "2018-02-04 20:07:05", "Formater Test: DateTimeString")
	assertDateTime(t, dt, dt.String())
}

func TestNonZero24HourFormat(t *testing.T) {
	tm := time.Date(2018, time.February, 4, 20, 7, 5, 9, time.UTC)
	dt := New(tm)

	assertTrue(t, dt.Format("{HH} {H}") == "20 20", "Test: HH H With 20hr")

	tm = time.Date(2018, time.February, 4, 9, 7, 5, 9, time.UTC)
	dt = New(tm)
	assertTrue(t, dt.Format("{HH} {H}") == "09 9", "Test: HH H With 09hr")
}

func TestIsStdFormat(t *testing.T) {
	assertTrue(t, isStdFormat(time.RFC1123), "")
}
