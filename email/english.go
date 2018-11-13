package email

import "time"

type enTemplate struct {
}

func (en enTemplate) subject() string {
	return "Your {{formatDate .RemindDate}} birthday reminder"
}

func (en enTemplate) body() string {
	return `The {{formatDate .RemindDate}}, don't forget to wish birthdays of :
{{range .Contacts}}
- {{.Name}}{{if yearValid .BirthDate}} ({{.Age}} yo) {{- end}}
{{end}}`
}
func (en enTemplate) dateLayout() string {
	return "01/02"
}

func (en enTemplate) formatDate(date time.Time) string {
	return date.Format(en.dateLayout())
}
