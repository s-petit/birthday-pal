package email

import "time"

type frDateTemplate struct {
}

func (fr frDateTemplate) dateLayout() string {
	return "02/01"
}

func (fr frDateTemplate) formatDate(date time.Time) string {
	return date.Format(fr.dateLayout())
}
