package email

import "time"

type frTemplate struct {
}

func (fr frTemplate) subject() string {
	return "Anniversaires du {{formatDate .RemindDate}}"
}

func (fr frTemplate) body() string {
	return `Le {{formatDate .RemindDate}}, n'oubliez pas de souhaiter l'anniversaire de :
{{range .Contacts}}
- {{.Name}}{{if yearValid .BirthDate}} ({{.Age $.RemindDate}} ans) {{- end}}
{{end}}`
}

func (fr frTemplate) dateLayout() string {
	return "02/01"
}

func (fr frTemplate) formatDate(date time.Time) string {
	return date.Format(fr.dateLayout())
}
