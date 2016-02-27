package gigasecond

import "time"

const TestVersion = 2

var Birthday time.Time
func AddGigasecond(birthday time.Time) time.Time {
  Birthday = birthday
  return birthday.Add(time.Duration(1000000000) * time.Second)
}
