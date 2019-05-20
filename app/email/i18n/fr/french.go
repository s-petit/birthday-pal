package fr

import (
	"time"
)

type Template struct {
}

func (t Template) subject() string {
	return "Anniversaires du {{formatDate .RemindParams.RemindDay}}"
}

func (t Template) body() string {
	return `{{if .RemindParams.Inclusive}}Durant les 7 prochains jours{{else}}Le {{formatDate .RemindParams.RemindDay}}{{end}}, n'oubliez pas de souhaiter l'anniversaire de :
{{range .Contacts}}
- {{.Name}}{{if yearValid .BirthDate}} ({{.Age $.RemindParams.RemindDay}} ans{{if $.RemindParams.Inclusive}} le {{formatDate .BirthDate}}{{- end}}) {{- end}}
{{end}}`
}

func (t Template) dateLayout() string {
	return "02/01"
}

func (t Template) formatDate(date time.Time) string {
	return date.Format(t.dateLayout())
}
