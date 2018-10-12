package email

import "time"

type enTemplate struct {
}

func (en enTemplate) simpleReminderSubject() string {
	return "Your {{formatDate .RemindDate}} birthday reminder"
}

func (en enTemplate) simpleReminderBody() string {
	return `The {{formatDate .RemindDate}}, don't forget to wish birthdays of :
{{range .Contacts}}
- {{.Name}}{{if yearValid .BirthDate}} ({{.Age}} yo) {{- end}}
{{end}}`
}

func (en enTemplate) weeklyDigestSubject() string {
	return "Anniversaire de {{.Name}}{{if yearValid .BirthDate}} -{{.Age}} an(s)- {{- end}}"
}

func (en enTemplate) weeklyDigestBody() string {
	return "Ce sera l'anniversaire de {{.Name}} le {{formatDate .BirthDate}}. Pensez à le lui souhaiter!"
}

func (en enTemplate) monthlyDigestSubject() string {
	return "Anniversaire de {{.Name}}{{if yearValid .BirthDate}} -{{.Age}} an(s)- {{- end}}"
}

func (en enTemplate) monthlyDigestBody() string {
	return "Ce sera l'anniversaire de {{.Name}} le {{formatDate .BirthDate}}. Pensez à le lui souhaiter!"
}

func (en enTemplate) dateLayout() string {
	return "01/02"
}

func (en enTemplate) formatDate(date time.Time) string {
	return date.Format(en.dateLayout())
}
