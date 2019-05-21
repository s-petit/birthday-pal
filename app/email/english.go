package email

import (
	"time"
)

//English represents an email in english
type English struct {
}

func (en English) subject() string {
	return "Your {{formatDate .RemindParams.RemindDay}} birthday reminder"
}

func (en English) body() string {
	return `{{if .RemindParams.Inclusive}}During the next 7 days{{else}}The {{formatDate .RemindParams.RemindDay}}{{end}}, don't forget to wish birthdays of :
{{range .Contacts}}
- {{.Name}}{{if yearValid .BirthDate}} ({{.Age $.RemindParams.RemindDay}} yo{{if $.RemindParams.Inclusive}} the {{formatDate .BirthDate}}{{- end}}) {{- end}}
{{end}}`
}
func (en English) dateLayout() string {
	return "01/02"
}

func (en English) formatDate(date time.Time) string {
	return date.Format(en.dateLayout())
}
