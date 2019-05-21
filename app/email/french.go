package email

import (
	"time"
)

//French represents an email in french
type French struct {
}

func (fr French) subject() string {
	return "Anniversaires du {{formatDate .RemindParams.RemindDay}}"
}

func (fr French) body() string {
	return `{{if .RemindParams.Inclusive}}Durant les 7 prochains jours{{else}}Le {{formatDate .RemindParams.RemindDay}}{{end}}, n'oubliez pas de souhaiter l'anniversaire de :
{{range .Contacts}}
- {{.Name}}{{if yearValid .BirthDate}} ({{.Age $.RemindParams.RemindDay}} ans{{if $.RemindParams.Inclusive}} le {{formatDate .BirthDate}}{{- end}}) {{- end}}
{{end}}`
}

func (fr French) dateLayout() string {
	return "02/01"
}

func (fr French) formatDate(date time.Time) string {
	return date.Format(fr.dateLayout())
}
