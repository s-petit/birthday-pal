package email

import (
	"fmt"
	"github.com/s-petit/birthday-pal/contact"
	"time"
)

const frenchBody = "To: Birthday Pals \r\n" +
	"Subject: Anniversaire de %s !\r\n" +
	"\r\n" +
	"Ce sera l'anniversaire de %s le %s. Il aura %d ans. Pensez a le lui souhaiter !\r\n"

const frenchLayout = "02/01"
const englishLayout = "01/01"

func formatDate(layout string, birthday time.Time) string {
	return birthday.Format(layout)
}

//TODO implement i18n

//French sends a reminder email in French.
func French(r contact.Contact) string {
	return fmt.Sprintf(frenchBody, r.Name, r.Name, formatDate(frenchLayout, r.BirthDate), r.Age)
}
