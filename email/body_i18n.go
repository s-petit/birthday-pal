package email

import (
	"time"
	"fmt"
	"github.com/s-petit/birthday-pal/contact"
)

const frenchBody  = "To: Birthday Pals \r\n" +
	"Subject: Anniversaire de %s !\r\n" +
	"\r\n" +
	"Ce sera l'anniversaire de %s le %s. Il aura %s ans. Pensez a le lui souhaiter !\r\n"

const frenchLayout  = "02/01"
const englishLayout  = "01/01"

func formatDate(layout string, birthday time.Time) string {
	return birthday.Format(layout)
}

// Age return the age of the contact at his incoming birthday
func FormatFrench(r contact.Contact) string {
	return fmt.Sprintf(frenchBody, r.Name, r.Name, formatDate(frenchLayout, r.LocalBirthDate()), r.Age(time.Now()))
}
