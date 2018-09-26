package email

import (
	"github.com/stretchr/testify/assert"
	"testing"
	"time"
)

func Test_formatEnglishDate(t *testing.T) {
	birthday := time.Date(2016, time.August, 22, 0, 0, 0, 0, time.UTC)
	formattedDate := enTemplate{}.formatDate(birthday)
	assert.Equal(t, "08/22", formattedDate)
}
