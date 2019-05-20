package en

import (
	"time"
)

type Template struct {
}

func (t Template) subject() string {
	return "Your {{formatDate .RemindParams.RemindDay}} birthday reminder"
}

func (t Template) body() string {
	return `{{if .RemindParams.Inclusive}}During the next 7 days{{else}}The {{formatDate .RemindParams.RemindDay}}{{end}}, don't forget to wish birthdays of :
{{range .Contacts}}
- {{.Name}}{{if yearValid .BirthDate}} ({{.Age $.RemindParams.RemindDay}} yo{{if $.RemindParams.Inclusive}} the {{formatDate .BirthDate}}{{- end}}) {{- end}}
{{end}}`
}
func (t Template) dateLayout() string {
	return "01/02"
}

func (t Template) formatDate(date time.Time) string {
	return date.Format(t.dateLayout())
}
