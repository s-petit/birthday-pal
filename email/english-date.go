package email

import "time"

type enDateTemplate struct {
}

func (en enDateTemplate) dateLayout() string {
	return "01/02"
}

func (en enDateTemplate) formatDate(date time.Time) string {
	return date.Format(en.dateLayout())
}
