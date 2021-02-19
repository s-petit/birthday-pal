package testdata

import (
	"github.com/stretchr/testify/mock"
	"time"
)

//FakeSystem represents a mockable system
type FakeSystem struct {
	mock.Mock
}

//Now return a mocked now
func (fs *FakeSystem) Now() time.Time {
	called := fs.Called()
	return called.Get(0).(time.Time)
}
