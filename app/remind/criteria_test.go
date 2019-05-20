package remind

import (
	"github.com/s-petit/birthday-pal/testdata"
	"github.com/stretchr/testify/assert"
	"testing"
	"time"
)

func Test_should_calculate_remind_day(t *testing.T) {

	currentDate := testdata.LocalDate(2018, time.August, 30)
	remindParams := Criteria{Today: currentDate, InNbDays: 3}

	assert.Equal(t, testdata.BirthDate(2018, time.September, 2), remindParams.RemindDay())
	assert.Equal(t, time.UTC, remindParams.RemindDay().Location())
}
