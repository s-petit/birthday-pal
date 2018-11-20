package email

type frWeeklyTemplate struct {
}

func (fr frTemplate) subject() string {
	return "Vos anniversaires de la semaine"
}

func (fr frTemplate) body() string {
	return `Cette semaine, n'oubliez pas de souhaiter l'anniversaire de :
{{range .Contacts}}
- {{.Name}}{{if yearValid .BirthDate}} ({{.Age}} ans), le {{formatDate .BirthDate}} {{- end}}
{{end}}`
}

