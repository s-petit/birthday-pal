package email

import (
	"fmt"
	"github.com/s-petit/birthday-pal/remind"
	"time"
)

const frenchBody = "To: Birthday Pals \r\n" +
	"Subject: Anniversaire de %s !\r\n" +
	"\r\n" +
	"Ce sera l'anniversaire de %s le %s. Il aura %d an(s). Pensez a le lui souhaiter !\r\n"

const frenchLayout = "02/01"
const englishLayout = "01/01"

func formatDate(layout string, birthday time.Time) string {
	return birthday.Format(layout)
}

//TODO implement i18n

//French sends a reminder email in French.
func French(c remind.ContactBirthday) string {
	return fmt.Sprintf(frenchBody, c.Name, c.Name, formatDate(frenchLayout, c.BirthDate), c.Age)
}
