package email

import (
	"time"
)

type enTemplate struct {
}

func (en enTemplate) subject() string {
	return "Your {{formatDate .RemindParams.RemindDay}} birthday reminder"
}

func (en enTemplate) body() string {
	return `{{if .RemindParams.Inclusive}}During the next 7 days{{else}}The {{formatDate .RemindParams.RemindDay}}{{end}}, don't forget to wish birthdays of :
{{range .Contacts}}
- {{.Name}}{{if yearValid .BirthDate}} ({{.Age $.RemindParams.RemindDay}} yo{{if $.RemindParams.Inclusive}} the {{formatDate .BirthDate}}{{- end}}) {{- end}}
{{end}}`
}
func (en enTemplate) dateLayout() string {
	return "01/02"
}

func (en enTemplate) formatDate(date time.Time) string {
	return date.Format(en.dateLayout())
}
