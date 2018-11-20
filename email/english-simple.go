package email

type enSimpleTemplate struct {
}

func (en enSimpleTemplate) subject() string {
	return "Your {{formatDate .RemindDate}} birthday reminder"
}

func (en enSimpleTemplate) body() string {
	return `The {{formatDate .RemindDate}}, don't forget to wish birthdays of :
{{range .Contacts}}
- {{.Name}}{{if yearValid .BirthDate}} ({{.Age}} yo) {{- end}}
{{end}}`
}
