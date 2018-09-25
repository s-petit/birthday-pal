package email

import "time"

type frTemplate struct {
}

func (fr frTemplate) subject() string {
	return "Anniversaire de {{.Name}}{{if yearValid .BirthDate}} -{{.Age}} an(s)- {{- end}}"
}

func (fr frTemplate) body() string {
	return "Ce sera l'anniversaire de {{.Name}} le {{formatDate .BirthDate}}. Pensez à le lui souhaiter!"
}

func (fr frTemplate) dateLayout() string {
	return "02/01"
}

func (fr frTemplate) formatDate(date time.Time) string {
	return date.Format(fr.dateLayout())
}
