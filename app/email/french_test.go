package email

import (
	"github.com/stretchr/testify/assert"
	"testing"
	"time"
)

func Test_formatFrenchDate(t *testing.T) {
	birthday := time.Date(2016, time.August, 22, 0, 0, 0, 0, time.UTC)
	formattedDate := frTemplate{}.formatDate(birthday)
	assert.Equal(t, "22/08", formattedDate)
}
