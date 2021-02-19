package system

import (
	"time"
)

// System holds system-dependant methods which are hard to test/mock
type System interface {
	Now() time.Time
}

// RealSystem is how the hosting system works in real life
type RealSystem struct {
}

//Now return the current date and time
func (rs RealSystem) Now() time.Time {
	return time.Now()
}
