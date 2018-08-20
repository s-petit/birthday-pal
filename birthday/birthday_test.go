package birthday

import (
	"testing"
	"time"
	"github.com/stretchr/testify/assert"
)

func Test_reminderToSend(t *testing.T) {

	birthday := time.Date(2016, time.August, 22, 0, 0, 0, 0, time.UTC)
	now := time.Date(2018, time.August, 21, 0, 0, 0, 0, time.Local)

	remind := Remind(now, birthday, 1)

	assert.Equal(t, true, remind)
}

func Test_reminderToSend2(t *testing.T) {

	birthday := time.Date(2016, time.August, 22, 0, 0, 0, 0, time.UTC)
	now := time.Date(2018, time.August, 20, 0, 0, 0, 0, time.Local)

	remind := Remind(now, birthday, 1)

	assert.Equal(t, false, remind)
}

func Test_reminderToSend3(t *testing.T) {

	birthday := time.Date(2016, time.August, 22, 0, 0, 0, 0, time.UTC)
	now := time.Date(2018, time.August, 22, 0, 0, 0, 0, time.Local)

	remind := Remind(now, birthday, 1)

	assert.Equal(t, false, remind)
}

func Test_reminderToSend4(t *testing.T) {

	birthday := time.Date(2016, time.August, 22, 0, 0, 0, 0, time.UTC)
	now := time.Date(2018, time.August, 23, 0, 0, 0, 0, time.Local)

	remind := Remind(now, birthday, 1)

	assert.Equal(t, false, remind)
}
