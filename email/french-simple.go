package email

type frSimpleTemplate struct {
}

func (fr frSimpleTemplate) subject() string {
	return "Anniversaires du {{formatDate .RemindDate}}"
}

func (fr frSimpleTemplate) body() string {
	return `Le {{formatDate .RemindDate}}, n'oubliez pas de souhaiter l'anniversaire de :
{{range .Contacts}}
- {{.Name}}{{if yearValid .BirthDate}} ({{.Age}} ans) {{- end}}
{{end}}`
}
