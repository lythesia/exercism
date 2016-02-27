package clock

import "fmt"

const TestVersion = 2

type Clock struct {
	hh, mm int
}

// ctor
func Time(hour, minute int) Clock {
	minutes := hour*60 + minute
	if minutes < 0 {
		minutes = minutes%1440 + 1440
	} else {
		minutes = minutes % 1440
	}
	return Clock{minutes / 60, minutes % 60}
}

// to_s
func (c Clock) String() string {
	return fmt.Sprintf("%02d:%02d", c.hh, c.mm)
}

// add
func (c Clock) Add(minutes int) Clock {
	return Time(c.hh, c.mm+minutes)
}

// equal
func (c Clock) Equal(x Clock) bool {
	return c.hh == x.hh && c.mm == x.mm
}
