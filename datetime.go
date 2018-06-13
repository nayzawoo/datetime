package datetime

import (
	"strings"
	"time"
)

// Datetime Wrapper Like Carbon, Moment, DaysJS
type Datetime struct {
	t              *time.Time
	formatReplacer *strings.Replacer
}

// New : Build
func New(t *time.Time) *Datetime {
	dt := &Datetime{
		t: t,
	}

	return dt
}

// Time https://golang.org/pkg/time/#Time
func (dt *Datetime) Time() *time.Time {
	return dt.t
}
